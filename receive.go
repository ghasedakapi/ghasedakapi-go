package ghasedakapi

import (
	"net/url"
	"io/ioutil"
	"encoding/json"
	"time"
	"strconv"
)
type ReceiveResult struct {
	Result    ResultItems
	Items	  ReceiveItem
}
type ReceiveItem struct{
	messageId		int64   `json:"messageId"`
	sender			string  `json:"sender"`
	message			string  `json:"message"`
	linenumber		string  `json:"linenumber"`
	receivedate		time.Time  `json:"receivedate"`
}
func (receive *ReceiveService) Receive(linenumber string,isread int) (receiveResult *ReceiveResult, err error) {
	v := url.Values{}
	v.Set("linenumber", linenumber)
	v.Set("isread", strconv.Itoa(isread))
	return receive.receiveMessage(v)
}

func (receive *ReceiveService) receiveMessage(formValues url.Values) (receiveResult *ReceiveResult, err error) {
	smsUrl := receive.client.BaseUrl + "/api/v1/sms/receive"
	res, err:=receive.client.Execute(smsUrl,formValues)
	if err != nil {
		return receiveResult, err
	}
	defer res.Body.Close()
	responseBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return receiveResult, err
	}
	receiveResult = new(ReceiveResult)
	err = json.Unmarshal(responseBody, receiveResult)
	return receiveResult, err
}
