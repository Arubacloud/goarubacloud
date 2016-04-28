package goarubacloud


import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os/user"
	"time"
	"net/http"
	"github.com/docker/machine/libmachine/log"
	"strings"
)

// DefaultTimeout api requests after 180s
const DefaultTimeout = 180

const (
	dc1 = "https://api.dc1.computing.cloud.it/WsEndUser/v2.9/WsEndUser.svc/json"
	dc2 = "https://api.dc2.computing.cloud.it/WsEndUser/v2.9/WsEndUser.svc/json"
	dc3 = "https://api.dc3.computing.cloud.it/WsEndUser/v2.9/WsEndUser.svc/json"
	dc4 = "https://api.dc4.computing.cloud.it/WsEndUser/v2.9/WsEndUser.svc/json"
	dc5 = "https://api.dc5.computing.cloud.it/WsEndUser/v2.9/WsEndUser.svc/json"
	dc6 = "https://api.dc6.computing.cloud.it/WsEndUser/v2.9/WsEndUser.svc/json"
)

// Custom errors
var (
	ErrNoEnpoint = errors.New("arubacloud: no endpoint provided")
)

var Endpoints = map[string]string {
	"dc1": dc1,
	"dc2": dc2,
	"dc3": dc3,
	"dc4": dc4,
	"dc5": dc5,
	"dc6": dc6,
}

type ResponseBody struct {
	ExceptionInfo	string	`json:"ExceptionInfo"`
	ResultCode	int	`json:"ResultCode"`
	ResultMessage	string  `json:"ResultMessage"`
	Success		bool	`json:"Success"`
	Value		json.RawMessage
}

type Client struct {
	EndPoint string
	Username string
	Password string
	Timeout  time.Duration
	Client   *http.Client
}

type APIResponse struct {
	StatusCode int
	Status string
	Body []byte
}

// Util: get user home
func currentUserHome() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	return usr.HomeDir, nil
}

// APIError represents an error that can occurred while calling the API.
type APIError struct {
	// Error message.
	Message string
	// HTTP code.
	Code int
}

func (err *APIError) Error() string {
	return fmt.Sprintf("Error %d: %q", err.Code, err.Message)
}

// NewDefaultClient returns an ArubaCloud API Client from external configuration
func NewDefaultClient() (*Client, error) {
	return NewClient("", "", "")
}

// NewEndpointClient returns an ArubaCloud API Client from external configuration, for a specific endpoint
func NewEndpointClient(endpoint string) (*Client, error) {
	return NewClient(endpoint, "", "")
}

// NewClient returns an ArubaCloud API Client.
func NewClient(endpoint, username, password string) (*Client, error) {

	// Timeout
	timeout := time.Duration(DefaultTimeout * time.Second)

	// Create client
	client := Client{
		Username: username,
		Password: password,
		Timeout:  timeout,
		Client:   &http.Client{},
	}

	// Get and check the configuration
	if err := client.loadConfig(endpoint); err != nil {
		return nil, err
	}
	return &client, nil
}

//
// High level API
//

func (c *Client) Get(url string, resType interface{}) error {
	return c.CallAPI("GET", url, nil, resType)
}

func (c *Client) Post(url string, data interface{}, resType interface{}) error {
	return c.CallAPI("POST", url, data, resType)
}

//
// Low Level Helpers
//

func (c *Client) getResponse(response *http.Response, resType interface{}) error {
	// Read all the response body
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	// < 200 && >= 300 : API error
	if response.StatusCode < http.StatusOK || response.StatusCode >= http.StatusMultipleChoices {
		apiError := &APIError{Code: response.StatusCode}
		if err = json.Unmarshal(body, apiError); err != nil {
			return err
		}

		return apiError
	}

	// Nothing to unmarshal
	if len(body) == 0 || resType == nil {
		return nil
	}

	//log.Debugf("[client.getResponse()] body: %#v", body)
	var rb ResponseBody
	if err := json.Unmarshal(body, &rb); err != nil {
		return fmt.Errorf("Error unmarshalling response body for: %s", body)
	}
	//log.Debugf("[client.getResponse()] rb: %#v", rb)
	if rb.Success == false {
		return fmt.Errorf("API Error: %s", strings.Split(rb.ResultMessage, "\r")[0])
	}
	if len(rb.Value) == 0 {
		return nil
	}
	return json.Unmarshal(rb.Value, &resType)
}

func (c *Client) CallAPI(method, path string, reqBody, resType interface{}) error {
	var body []byte
	var err error

	if reqBody != nil {
		body, err = json.Marshal(reqBody)
		if err != nil {
			return err
		}
	}

	log.Debugf("[goarubacloud.client.CallAPI] marshalled body: %s", body)
	target := fmt.Sprintf("%s%s", c.EndPoint, path)
	log.Debugf("[goarubacloud.client.CallAPI] target: %s", target)
	req, err := http.NewRequest(method, target, bytes.NewReader(body))
	if err != nil {
		return err
	}

	// Inject headers
	if body != nil {
		req.Header.Add("Content-Type", "application/json;charset=utf-8")
	}
	req.Header.Add("Accept", "application/json")

	// Send the request with requested timeout
	c.Client.Timeout = c.Timeout
	response, err := c.Client.Do(req)

	log.Debugf("[goarubacloud.client.CallAPI] Status: %s", response.Status)
	//log.Debugf("[goarubacloud.client.CallAPI] response: %#v", response)
	if err != nil {
		return err
	}

	// Unmarshal the result into the resType if possible
	return c.getResponse(response, resType)
}