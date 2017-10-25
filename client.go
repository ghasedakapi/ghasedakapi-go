package ghasedakapi

import(
	"net/http"
	"net/url"
	"strings"
	"encoding/json"
)

const(
	baseUrl="http://ghasedakapi.parsaspace.com"
)
type Client struct {
	ApiKey string
	BaseUrl    string
	HTTPClient *http.Client
}

// Create a new ghasedakapi struct.
func NewClient(apiKey string) *Client {
	return NewHttpClient(apiKey, nil)
}

// Create a new ghasedakapi client, using a http.Client
func NewHttpClient(apiKey string, HTTPClient *http.Client) *Client {
	if HTTPClient == nil {
		HTTPClient = http.DefaultClient
	}
	return &Client{apiKey,baseUrl, HTTPClient}
}


func (client *Client) Execute(apiUrl string, formValues  url.Values)(*http.Response,error) {
	req, err := http.NewRequest("POST", apiUrl, strings.NewReader(formValues.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Accept-Charset", "utf-8")
	c := client.HTTPClient
	if c == nil {
		c = http.DefaultClient
	}
	 resp,err:=c.HTTPClient.Do(req)
	if err != nil {
		if err, ok := err.(net.Error); ok && err.Timeout() {
			return resp, err
		}
		return &HTTPError{
			Code:  resp.StatusCode,
			Message: resp.Status,
			Error:     err,
		}
	}
	defer resp.Body.Close()
	if  resp.StatusCode!=200 {
		ApiErr = new(APIError)
		err = json.Unmarshal(resp, ApiErr)
		if err !=nil{
			return resp, err
		}
		return &APIError{
			Code:  ApiErr.Result.Code,
			Message: ApiErr.Result.Message,
		}
	}
	return resp,nil
}
