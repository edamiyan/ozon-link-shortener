package model

import (
	"fmt"
	"regexp"
)

type Link struct {
	ID      int    `json:"id" db:"id"`
	BaseURL string `json:"base_url,omitempty" db:"base_url"`
	Token   string `json:"short_url,omitempty" db:"short_url"`
}

func ValidateBaseURL(p *Link) error {

	if p == nil {
		return fmt.Errorf("pass nil pointer")
	}

	if p.BaseURL == "" {
		return fmt.Errorf("empty query")
	}

	pattern := `^(https?://|www.)?[a-zA-Z0-9-]{1,256}([.][a-zA-Z-]{1,256})?([.][a-zA-Z]{1,30})([/]?[a-zA-Z0-9/?=%&#_.-]+)`
	if p.BaseURL != "" {
		if valid, _ := regexp.Match(pattern, []byte(p.BaseURL)); !valid {
			return fmt.Errorf("%v is a invalid base url", p.BaseURL)
		}
	}

	return nil
}
