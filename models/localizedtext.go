package models

type LocalizedText struct {
	LanguageID                          int        `json:"LanguageID,omitempty"`
	Text 	string `json:"Text,omitempty"`
}

