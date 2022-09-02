package gofiery

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

type FieryClient struct {
	ServerAddr     *net.TCPAddr
	APIVersion     int
	EndpointPrefix string
	Username       string
	Password       string
	Key            string
	Cookie         string
	HTTPClient     http.Client
}

type authPayload struct {
	username string
	password string
	apikey   string
}

type response struct {
	time time.Time
	data responseData[any]
}

type responseData[T any] struct {
	kind string
	item T
}

// NewFieryClient creates a new client to communicate with a Fiery API server.
func NewFieryClient(addr string, version int, prefix string, user string, pass string, key string) *FieryClient {
	sa, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		panic(err)
	}
	return &FieryClient{
		ServerAddr:     sa,
		APIVersion:     version,
		EndpointPrefix: prefix,
		Username:       user,
		Password:       pass,
		Key:            key,
		HTTPClient:     http.Client{Timeout: 30 * time.Second},
	}
}

func (fc *FieryClient) postflight(res *http.Response, do func()) {
	if !fc.ResponseIsJSON(res) {
		fmt.Fprintf(os.Stderr, "gofiery: HTTP error: endpoint returned %v but expected %v\n", res.Header.Get("Content-Type"), "application/json")
		os.Exit(0)
	}
	if fc.ResponseIsOK(res) {
		do()
	} else {
		fmt.Fprintf(os.Stderr, "gofiery: HTTP error: endpoint returned status code %v but expected %v\n", res.StatusCode, http.StatusOK)
		os.Exit(0)
	}
}

// Endpoint returns from uri the full URL path to the API endpoint.
// For use in requesting the API.
func (fc *FieryClient) Endpoint(uri string) string {
	return fc.ServerAddr.String() + fc.EndpointPrefix + uri
}

// ResponseIsOK checks that an HTTP response res returned a 200 status code.
// Returns false otherwise.
func (fc *FieryClient) ResponseIsOK(res *http.Response) bool {
	return res.StatusCode == http.StatusOK
}

// ResponseIsJSON checks that the HTTP response res has a content type
// of "application/json". Returns false otherwise.
func (fc *FieryClient) ResponseIsJSON(res *http.Response) bool {
	return res.Header.Get("Content-Type") == "application/json"
}

// Login initiates the login process against the Fiery API by providing
// the username, password and API key from fc to the destination server.
// If the login is successful, the session cookie is then stored in fc
// for use in authentication-required endpoints.
func (fc *FieryClient) Login() {
	payload, err := json.Marshal(authPayload{username: fc.Username, password: fc.Password, apikey: fc.Key})
	if err != nil {
		fmt.Fprint(os.Stderr, "gofiery: unable to marshal json authentication payload")
		panic(err)
	}
	r := bytes.NewReader(payload)
	req, _ := http.NewRequest(http.MethodPost, fc.Endpoint("login"), r)
	req.Header.Set("Content-Type", "application/json")
	res, err := fc.HTTPClient.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "gofiery: error making http request: %s\n", err)
	}
	fc.postflight(res, func() { fc.Cookie = res.Header.Get("Set-Cookie") })
}

// Logout initiates the logout process against the Fiery API.
// If the request is successful, the session cookie given by
// the login process is removed from the client.
func (fc *FieryClient) Logout() {
	if fc.Cookie == "" {
		fmt.Fprint(os.Stderr, "gofiery: missing cookie in client. did you login?")
	}
	r := bytes.NewReader([]byte{})
	req, _ := http.NewRequest(http.MethodPost, fc.Endpoint("logout"), r)
	req.Header.Set("Cookie", fc.Cookie)
	res, err := fc.HTTPClient.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "gofiery: error making http request: %s\n", err)
	}
	fc.postflight(res, func() { fc.Cookie = "" })
}

// Run do the request to a Fiery API endpoint and place the result in a reference of resultContainer
func (fc *FieryClient) Run(endpoint string, method string) *response {
	var data response
	r := bytes.NewReader([]byte{})
	if fc.Cookie == "" {
		fmt.Fprint(os.Stderr, "gofiery: missing cookie in client. did you login?")
	}
	req, _ := http.NewRequest(method, fc.Endpoint(endpoint), r)
	req.Header.Set("Cookie", fc.Cookie)
	res, err := fc.HTTPClient.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "gofiery: error making http request: %s\n", err)
	}
	fc.postflight(res, func() {
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				log.Fatal("gofiery: unable to close response body")
			}
		}(res.Body)
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "gofiery: error reading HTTP response: %s\n", err)
		}
		err = json.Unmarshal(body, &data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "gofiery: could not parse JSON response: %s\n", err)
		}
	})
	return &data
}
