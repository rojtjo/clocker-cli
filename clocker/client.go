/*
Copyright © 2020 Roj Vroemen <me@rojvroemen.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package clocker

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
)

// Client is handles the
type Client struct {
	rootURL *url.URL
	token   string
	http    *http.Client
}

type startPayload struct {
	Client      string `json:"client"`
	Description string `json:"description"`
}

// NewClient returns a new Client given a url
func NewClient(rootURL, token string) (*Client, error) {
	u, err := url.Parse(rootURL)
	if err != nil {
		return nil, err
	}

	return &Client{
		rootURL: u,
		http:    &http.Client{},
		token:   token,
	}, nil
}

// Start clocking for a client with the given description
func (c *Client) Start(client, description string) (*StartResponse, error) {
	var res *StartResponse
	req := StartRequest{
		Client:      client,
		Description: description,
	}

	if err := c.Do(req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

// Stop clocking if clocking
func (c *Client) Stop() (*StopResponse, error) {
	var res *StopResponse
	req := StopRequest{}

	if err := c.Do(req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

// Status returns the current status of clocker
func (c *Client) Status() (*StatusResponse, error) {
	var res *StatusResponse
	req := StatusRequest{}

	if err := c.Do(req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

// Do executes the given request and decodes the response into the given var
func (c *Client) Do(r Request, o interface{}) error {
	body := &bytes.Buffer{}
	if err := json.NewEncoder(body).Encode(r); err != nil {
		return err
	}

	path := &url.URL{Path: r.Endpoint()}
	url := c.rootURL.ResolveReference(path)

	req, err := http.NewRequest("POST", url.String(), body)
	if err != nil {
		return err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+c.token)

	resp, err := c.http.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&o); err != nil {
		return err
	}

	return nil
}
