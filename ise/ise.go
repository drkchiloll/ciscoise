package ise

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Client properties of an ISE Instance
type Client struct {
	BaseURL   string
	username  string
	password  string
	IP        string
	http      *http.Client
	csrfToken string
}

// New creates an Instance of an ISE client
func New(host, user, pass string, ignoreSSL bool) *Client {
	return &Client{
		BaseURL:  fmt.Sprintf("https://%s:9060/ers/config", host),
		username: user,
		password: pass,
		IP:       host,
		http: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
			Timeout: 8 * time.Second,
		},
	}
}

// Filter ISE Request Filter Options
type Filter struct {
	Equals        string
	NotEquals     string
	Greater       string
	Less          string
	StartsWith    string
	NotStartsWith string
	EndsWith      string
	NotEndsWith   string
	Contains      string
	NotContains   string
}

// FILTER enumerator
var FILTER = Filter{
	Equals:        "EQ",
	NotEquals:     "NEQ",
	Greater:       "GT",
	Less:          "LT",
	StartsWith:    "STARTSW",
	NotStartsWith: "NSTARTSW",
	EndsWith:      "ENDSW",
	NotEndsWith:   "NENDSW",
	Contains:      "CONTAINS",
	NotContains:   "NCONTAINS",
}

// MakeReq abstracts all requests for the ISE API Wrapper
func (c *Client) MakeReq(path, method string, o ReqParams, r io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, c.BaseURL+path, r)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}
	c.addQS(req, o)
	res, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %v", err)
	}
	return res, nil
}

func (c *Client) addQS(r *http.Request, p ReqParams) {
	q := r.URL.Query()
	if p.Filter != nil {
		for _, f := range p.Filter {
			q.Add("filter", f)
		}
	}
	r.URL.RawQuery = q.Encode()
	// Prints the URL ... for testing
	// fmt.Println(r.URL)
	r.SetBasicAuth(c.username, c.password)
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Accept", "application/json")
}

// Object common ISE Object
type Object struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name"`
	Desc string `json:"description,omitempty"`
}

// ObjRef reference to an ISE Object
type ObjRef struct {
	Link     `json:"link,omitempty"`
	NextPage Link `json:"nextPage,omitempty"`
	PrevPage Link `json:"previousPage,omitempty"`
}

// ReqParams ...
type ReqParams struct {
	Size   int
	Page   int
	Filter []string
	Sort   string
}

// Link is a reference to an ISE Object
type Link struct {
	Rel  string `json:"rel"`
	Href string `json:"href"`
	Type string `json:"type"`
}
