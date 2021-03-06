package ghasedakapi

import (
	"net/url"
	"io/ioutil"
	"encoding/json"
	"fmt"
)

type AccountResult struct {
	Result    ResultItems
	Items	  AccountInfoItem
}
type AccountInfoItem struct{
	Balance		int  `json:"balance"`
	Expire		int64  `json:"expire"`
}

//Get Account Information
func (account *AccountService) GetInfo() (accountres *AccountResult, err error) {
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
	fmt.Println(res.Body)
	err = json.Unmarshal(responseBody, accountres)
	return accountres, err
}
