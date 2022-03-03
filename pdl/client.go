package pdl

import (
	"fmt"
	"net/http"
)

const (
	defaultBaseURL = "https://api.peopledatalabs.com"
	defaultVersion = "v5"
)

// PDL represents configuration related to PDL API.
type PDL struct {
	apiKey  string
	baseURL string

	client *http.Client // consumer of API might want to supply their own http client here
}

// New creates new instance of PDL.
func New(apiKey, version string) *PDL {
	return &PDL{
		apiKey:  apiKey,
		baseURL: fmt.Sprintf("%s/%s", defaultBaseURL, version),
		client:  http.DefaultClient,
	}
}

// NewV5 creates new instance of PDL for connecting to V5 API.
// It is a convenience wrapper for New.
func NewV5(apiKey string) *PDL {
	return &PDL{
		apiKey:  apiKey,
		baseURL: fmt.Sprintf("%s/v5", defaultBaseURL),
		client:  http.DefaultClient,
	}
}

// Person and Company are placeholder, might want to rethink this.
// They probably belong to different files, but for the sake of simplicity,
// putting them here for now.
// Person represents endpoints for person-related API.
type Person struct {
	*PDL
}

// NewPerson creates new instance of PDL person services.
func NewPerson(pdl *PDL) *Person {
	pdl.baseURL = fmt.Sprintf("%s/person", pdl.baseURL)
	return &Person{pdl}
}

// Company represents endpoints for company-related API.
type Company struct {
	*PDL
}

// NewCompany creates new instance of PDL company services.
func NewCompany(pdl *PDL) *Company {
	pdl.baseURL = fmt.Sprintf("%s/company", pdl.baseURL)
	return &Company{pdl}
}
