package ghasedakapi

import (
	"net/url"
	"io/ioutil"
	"encoding/json"
)

type ApiResult struct {
	Result    ResultItems
	Items	  int64
}

type ResultItems struct{
	Code 		string		`json:"code"`
	Message		string		`json:"message"`
}


//Send ...
func (sms *SMSService) Send(message string,sender string,receptor string) (apiResult *ApiResult, err error) {
	v := url.Values{}
	v.Set("sender", sender)
	v.Set("receptor", receptor)
	v.Set("message", message)
	return sms.sendMessage(v)
}

// Core method to send message
func (sms *SMSService) sendMessage(formValues url.Values) (apiResult *ApiResult, err error) {
	smsUrl := sms.client.BaseUrl + "/api/v1/sms/send/simple"
	res, err :=sms.client.Execute(smsUrl,formValues)
	if err != nil {
		return apiResult, err
	}
	defer res.Body.Close()

	responseBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return apiResult, err
	}
	apiResult = new(ApiResult)
	err = json.Unmarshal(responseBody, apiResult)
	return apiResult, err
}
