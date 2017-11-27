package ghasedakapi

import (
	"net/url"
	"io/ioutil"
	"encoding/json"
	"strconv"
)

type GroupNumbersResult struct {
	Result    ResultItems
	Items	  GroupNumbersItem
}
type GroupNumbersItem struct{
	number		string `json:"number"`
	firstname	string `json:"firstname"`
	lastname	string `json:"lastname"`
	email		string `json:"email"`
}

type AddGroupResult struct {
	Result    ResultItems
	Items	  AddGroupItem
}
type AddGroupItem struct{
	Groupid		int  `json:"Groupid"`
}


func (contact *ContactService) GetGroupNumbers(groupid int,offset int ,page int) (groupnumbersresult *GroupNumbersResult, err error) {
	v := url.Values{}
	v.Set("groupid", strconv.Itoa(groupid))
	v.Set("offset", strconv.Itoa(offset))
	v.Set("page", strconv.Itoa(page))
	return contact.GroupNumbers(v)
}

func (account *ContactService) GroupNumbers(formValues url.Values) (groupnumbersresult *GroupNumbersResult, err error) {
	smsUrl := account.client.BaseUrl + "/api/v1/contact/group/number/add"
	res, err:=account.client.Execute(smsUrl,formValues)
	if err != nil {
		return groupnumbersresult, err
	}
	defer res.Body.Close()
	responseBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return groupnumbersresult, err
	}
	groupnumbersresult = new(GroupNumbersResult)
	err = json.Unmarshal(responseBody, groupnumbersresult)
	return groupnumbersresult, err
}


func (contact *ContactService) AddNumberToGroup(groupid int,number string ,firstname string,lastname string,email string) (apiResult *ApiResult, err error) {
	v := url.Values{}
	v.Set("groupid", strconv.Itoa(groupid))
	v.Set("number", number)
	v.Set("firstname", firstname)
	v.Set("lastname", lastname)
	v.Set("email", email)
	return contact.addGroupNumbers(v)
}

func (account *ContactService) addGroupNumbers(formValues url.Values) (apiResult *ApiResult, err error) {
	smsUrl := account.client.BaseUrl + "/api/v1/contact/group/number/add"
	res, err:=account.client.Execute(smsUrl,formValues)
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

func (contact *ContactService) AddGroup(name string ,parent int) (addgroupresult *AddGroupResult, err error) {
	v := url.Values{}
	v.Set("name", name)
	v.Set("parent", strconv.Itoa(parent))
	return contact.addGroup(v)
}

func (account *ContactService) addGroup(formValues url.Values) (addgroupresult *AddGroupResult, err error) {
	smsUrl := account.client.BaseUrl + "/api/v1/contact/group/add"
	res, err:=account.client.Execute(smsUrl,formValues)
	if err != nil {
		return addgroupresult, err
	}
	defer res.Body.Close()
	responseBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return addgroupresult, err
	}
	addgroupresult = new(AddGroupResult)
	err = json.Unmarshal(responseBody, addgroupresult)
	return addgroupresult, err
}
