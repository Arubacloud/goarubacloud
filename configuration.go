package goarubacloud

import (
	"gopkg.in/ini.v1"
	"os"
	"strings"
	"path/filepath"
	"fmt"
)

// Use variables for easier test overload
var (
	systemConfigPath = "/etc/arubacloud.conf"
	userConfigPath   = "/.arubacloud.conf" // prefixed with homeDir
	localConfigPath  = "./arubacloud.conf"
)


// appendConfigurationFile only if it exists. We need to do this because
// ini package will fail to load configuration at all if a configuration
// file is missing. This is racy, but better than always failing.
func appendConfigurationFile(cfg *ini.File, path string) {
	if file, err := os.Open(path); err == nil {
		file.Close()
		cfg.Append(path)
	}
}

// getConfigValue returns the value of AC_<NAME> or ``name`` value from ``section``
func getConfigValue(cfg *ini.File, section, name string) string {
	// Attempt to load from environment
	fromEnv := os.Getenv("AC_" + strings.ToUpper(name))
	if len(fromEnv) > 0 {
		return fromEnv
	}

	// Attempt to load from configuration
	fromSection := cfg.Section(section)
	if fromSection == nil {
		return ""
	}

	fromSectionKey := fromSection.Key(name)
	if fromSectionKey == nil {
		return ""
	}
	return fromSectionKey.String()
}

// loadConfig loads client configuration from params, environments or configuration
// files (by order of decreasing precedence).
//
// loadConfig will check AC_USERNAME, AC_PASSWORD
// and AC_ENDPOINT environment variables. If any is present, it will take precedence
// over any configuration from file.
//
// Configuration files are ini files.
//
// - ./arubacloud.conf
// - $HOME/.arubacloud.conf
// - /etc/arubacloud.conf
//
func (c *Client) loadConfig(endpointName string) error {
	// Load configuration files by order of increasing priority. All configuration
	// files are optional. Only load file from user home if home could be resolve
	cfg := ini.Empty()
	appendConfigurationFile(cfg, systemConfigPath)
	if home, err := currentUserHome(); err == nil {
		userConfigFullPath := filepath.Join(home, userConfigPath)
		appendConfigurationFile(cfg, userConfigFullPath)
	}
	appendConfigurationFile(cfg, localConfigPath)

	// Canonicalize configuration
	if endpointName == "" {
		endpointName = getConfigValue(cfg, "default", "dc1")
	}

	if c.Username == "" {
		c.Username = getConfigValue(cfg, endpointName, "username")
	}

	if c.Password == "" {
		c.Password = getConfigValue(cfg, endpointName, "password")
	}

	// Load real endpoint URL by name. If endpoint contains a '/', consider it as a URL
	if strings.Contains(endpointName, "/") {
		c.EndPoint = endpointName
	} else {
		c.EndPoint = Endpoints[endpointName]
	}

	// If we still have no valid endpoint, AppKey or AppSecret, return an error
	if c.EndPoint == "" {
		return fmt.Errorf("Unknown endpoint '%s'. Consider checking 'Endpoints' list of using an URL.", endpointName)
	}
	if c.Username == "" {
		return fmt.Errorf("Missing username. Please check your configuration or consult the documentation to create one.")
	}
	if c.Password == "" {
		return fmt.Errorf("Missing password. Please check your configuration or consult the documentation to create one.")
	}

	return nil
}
