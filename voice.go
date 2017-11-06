package ghasedakapi

import (
	"net/url"
	"io/ioutil"
	"encoding/json"
)
func (voice *VoiceService) Send(message string,receptor []string) (apiResult *ApiResult, err error) {
	v := url.Values{}
	v.Set("receptor", arrayToString(receptor))
	v.Set("message", message)
	return voice.sendMessage(v)
}

// Core method to send voice
func (voice *VoiceService) sendMessage(formValues url.Values) (apiResult *ApiResult, err error) {
	smsUrl := voice.client.BaseUrl + "/api/v1/voice/send"
	res, err:=voice.client.Execute(smsUrl,formValues)
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
