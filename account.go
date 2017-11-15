package ghasedakapi

import (
	"net/url"
	"io/ioutil"
	"encoding/json"
)

//Get Account Information
func (account *AccountService) getinfo() (apiResult *ApiResult, err error) {
	v := url.Values{}
	return account.makeRequest(v)
}


// Core method to send message
func (sms *SMSService) makeRequest(formValues url.Values) (apiResult *ApiResult, err error) {
	smsUrl := sms.client.BaseUrl + "/api/v1/account/info"
	res, err:=sms.client.Execute(smsUrl,formValues)
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
