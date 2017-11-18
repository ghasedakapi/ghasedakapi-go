package ghasedakapi

import (
	"net/url"
	"io/ioutil"
	"encoding/json"
)

type AccountResult struct {
	Result    ResultItems
	Items	  []int64  `json:"items"`
}
type AccountInfoItem struct{
	Balance		int
	Expire		int
}

//Get Account Information
func (account *AccountService) getinfo() (apiResult *AccountResult, err error) {
	v := url.Values{}
	return account.makeRequest(v)
}



func (account *AccountService) makeRequest(formValues url.Values) (apiResult *AccountResult, err error) {
	smsUrl := account.client.BaseUrl + "/api/v1/account/info"
	res, err:=account.client.Execute(smsUrl,formValues)
	if err != nil {
		return apiResult, err
	}
	defer res.Body.Close()
	responseBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return apiResult, err
	}
	apiResult = new(AccountResult)
	err = json.Unmarshal(responseBody, apiResult)
	return apiResult, err
}
