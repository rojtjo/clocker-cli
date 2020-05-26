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

// Request represents an API call
type Request interface {
	Endpoint() string
}

type messageResponse struct {
	Message string `json:"message"`
}

// StartRequest contains the required data for a start request
type StartRequest struct {
	Client      string `json:"client"`
	Description string `json:"description"`
}

// StartResponse contains the response from the successful StartRequest
type StartResponse struct {
	messageResponse
}

// Endpoint for StartRequest
func (r StartRequest) Endpoint() string {
	return "/api/clocker/start"
}

// StopRequest contains the required data for a start request
type StopRequest struct {
}

// StopResponse contains the response from the successful StopRequest
type StopResponse struct {
	messageResponse
}

// Endpoint for StopRequest
func (r StopRequest) Endpoint() string {
	return "/api/clocker/stop"
}

// StatusRequest contains the required data for a start request
type StatusRequest struct {
}

// StatusResponse contains the response from the successful StatusRequest
type StatusResponse struct {
	messageResponse
}

// Endpoint for StatusRequest
func (r StatusRequest) Endpoint() string {
	return "/api/clocker/status"
}
