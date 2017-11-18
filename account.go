package ghasedakapi

import (
	"net/url"
	"io/ioutil"
	"encoding/json"
)

type AccountResult struct {
	Result    ResultItems
	Items	  AccountInfoItem
}
type AccountInfoItem struct{
	Balance		int
	Expire		int
}

//Get Account Information
func (account *AccountService) getinfo() (accountres *AccountResult, err error) {
	v := url.Values{}
	return account.makeRequest(v)
}



func (account *AccountService) makeRequest(formValues url.Values) (accountres *AccountResult, err error) {
	smsUrl := account.client.BaseUrl + "/api/v1/account/info"
	res, err:=account.client.Execute(smsUrl,formValues)
	if err != nil {
		return accountres, err
	}
	defer res.Body.Close()
	responseBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return accountres, err
	}
	accountres = new(AccountResult)
	err = json.Unmarshal(responseBody, accountres)
	return accountres, err
}
