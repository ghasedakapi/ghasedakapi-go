package ghasedakapi

import (
	"net/url"
	"io/ioutil"
	"encoding/json"
	"strconv"
)
func (t *TemplateService) Verify(template string,messagetype int,receptor []string,param1 string,param2 string,param3 string) (apiResult *ApiResult, err error) {
	v := url.Values{}
	v.Set("template", template)
	v.Set("receptor", arrayToString(receptor))
	v.Set("type", strconv.Itoa(messagetype))
	v.Set("param1", param1)
	v.Set("param2", param2)
	v.Set("param3", param3)
	return t.makeRequest(v)
}

func (t *TemplateService) makeRequest(formValues url.Values) (apiResult *ApiResult, err error) {
	smsUrl := t.client.BaseUrl + "/api/v1/sms/verify"
	res, err:=t.client.Execute(smsUrl,formValues)
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
