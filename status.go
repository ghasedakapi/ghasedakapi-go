package ghasedakapi

import (
	"net/url"
	"io/ioutil"
	"encoding/json"
	"time"
)
type StatusResult struct {
	Result    ResultItems
	Items	  StatusItem
}
type StatusItem struct{
	messageId		int64  `json:"messageId"`
	status			int  `json:"status"`
}
type SelectResult struct {
	Result    ResultItems
	Items	  SelectItem
}
type SelectItem struct{
	messageId		int64  `json:"messageId"`
	status			int  `json:"status"`
	message			string  `json:"message"`
	price			int  `json:"price"`
	sender			string  `json:"sender"`
	receptor		string  `json:"receptor"`
	senddate		time.Time `json:"senddate"`

}



func (status *StatusService) Status( messageid []int64) (statusResult *StatusResult, err error) {
	v := url.Values{}
	v.Set("messageid", arrayToString(messageid))
	return status.GetStatus(v)
}

func (status *StatusService) Select( messageid []int64) (selectResult *SelectResult, err error) {
	v := url.Values{}
	v.Set("messageid", arrayToString(messageid))
	return status.SelectMessage(v)
}

func (status *StatusService) SelectMessage(formValues url.Values) (selectResult *SelectResult, err error) {
	smsUrl := status.client.BaseUrl + "/api/v1/sms/select"
	res, err:=status.client.Execute(smsUrl,formValues)
	if err != nil {
		return selectResult, err
	}
	defer res.Body.Close()
	responseBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return selectResult, err
	}
	selectResult = new(SelectResult)
	err = json.Unmarshal(responseBody, selectResult)
	return selectResult, err
}

func (status *StatusService) GetStatus(formValues url.Values) (statusResult *StatusResult, err error) {
	smsUrl := status.client.BaseUrl + "/api/v1/sms/status"
	res, err:=status.client.Execute(smsUrl,formValues)
	if err != nil {
		return statusResult, err
	}
	defer res.Body.Close()
	responseBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return statusResult, err
	}
	statusResult = new(StatusResult)
	err = json.Unmarshal(responseBody, statusResult)
	return statusResult, err
}
