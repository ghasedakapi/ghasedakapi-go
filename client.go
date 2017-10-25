package ghasedakapi

import(
	"net/http"
	"net/url"
	"strings"
	"net"
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
	 resp,err:=c.Do(req)
	if err != nil {
		if err, ok := err.(net.Error); ok && err.Timeout() {
			return resp, err
		}
		return resp,&HTTPError{
			Code:  resp.StatusCode,
			Message: resp.Status,
			Err:     err,
		}
	}
	defer resp.Body.Close()
	if  resp.StatusCode!=200 {
		exception = new(APIError)
		err = json.Unmarshal(resp, exception)
		if err !=nil{
			return resp, err
		}
		return resp,&APIError{
			Code:  exception.Result.Code,
			Message: exception.Result.Message,
		}
	}
	return resp,nil
}
