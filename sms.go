package ghasedakapi

import (
	"net/url"
	"io/ioutil"
	"encoding/json"
	"fmt"
)
func (sms *SMSService) Send(message string,sender string,receptor []string) (apiResult *ApiResult, err error) {
	v := url.Values{}
	v.Set("linenumber", sender)
	v.Set("receptor", arrayToString(receptor))
	v.Set("message", message)
	return sms.sendMessage(v)
}

// Core method to send message
func (sms *SMSService) sendMessage(formValues url.Values) (apiResult *ApiResult, err error) {
	smsUrl := sms.client.BaseUrl + "/api/v1/sms/send/bulk2"
	res, err:=sms.client.Execute(smsUrl,formValues)
	if err != nil {
		return apiResult, err
	}
	defer res.Body.Close()
	fmt.Println(res.Body)
	responseBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return apiResult, err
	}
	apiResult = new(ApiResult)
	err = json.Unmarshal(responseBody, apiResult)
	return apiResult, err
}
