package artifactorymc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type clientConfig struct {
	user string
	pass string
	url  string
}

// Client is used to call Artifactory REST APIs
type Client interface {
	Ping() error
}

var _ Client = clientConfig{}

// NewClient constructs a new artifactory client
func NewClient(username string, pass string, url string) Client {
	return clientConfig{
		username,
		pass,
		strings.TrimRight(url, "/"),
	}
}

// Ping calls the system to verify connectivity
func (c clientConfig) Ping() error {
	resp, err := c.execute("GET", "system/ping", nil)

	if err != nil {
		return err
	}

	if err := c.validateResponse(200, "OK", resp); err != nil {
		return err
	}

	return resp.Body.Close()
}

func (c clientConfig) execute(method string, endpoint string, payload interface{}) (*http.Response, error) {
	client := &http.Client{}
	url := fmt.Sprintf("%s/api/%s", c.url, endpoint)
	log.Printf("[DEBUG] Sending Request to method/url: %s %s", method, url)

	var jsonpayload *bytes.Buffer
	if payload == nil {
		jsonpayload = &bytes.Buffer{}
	} else {
		var jsonbuffer []byte
		jsonpayload = bytes.NewBuffer(jsonbuffer)
		enc := json.NewEncoder(jsonpayload)
		err := enc.Encode(payload)
		if err != nil {
			log.Printf("[ERROR] Error Encoding Payload: %s", err)
			return nil, err
		}
	}

	req, err := http.NewRequest(method, url, jsonpayload)
	if err != nil {
		log.Printf("[ERROR] Error creating new request: %s", err)
		return nil, err
	}
	req.SetBasicAuth(c.user, c.pass)
	req.Header.Add("content-type", "application/json")

	return client.Do(req)
}

func (c clientConfig) validateResponse(expectedCode int, action string, resp *http.Response) (err error) {
	if resp.StatusCode != expectedCode {
		response := ""
		if resp, err := ioutil.ReadAll(resp.Body); err == nil {
			response = fmt.Sprintf(" Response: %s", string(resp))
		}
		request := ""
		if req, err := ioutil.ReadAll(resp.Request.Body); err == nil {
			headers := map[string][]string{}
			for name, header := range resp.Request.Header {
				headers[name] = header
			}
			request = fmt.Sprintf(" Request:%s\n%s", headers, string(req))
		}
		return fmt.Errorf("Failed to %s. Status: %s.%s%s", action, resp.Status, request, response)
	}
	return nil
}
