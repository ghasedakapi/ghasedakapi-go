package ghasedakapi

type GhasedakApi struct {
	SMS *SMSService
	Voice *VoiceService
}
//MessageService ...
type SMSService struct {
	client *Client
}
type VoiceService struct {
	client *Client
}


func New(apikey string) *GhasedakApi {
	client := NewClient(apikey)
	return NewWithClient(client)
}


func NewWithClient(client *Client) *GhasedakApi {
	obj := &GhasedakApi{}
	obj.SMS = NewSMSService(client)
	obj.Voice=NewVoiceService(client)
	return obj
}

func NewVoiceService(client *Client) *VoiceService {
	return &VoiceService{client: client}
}
func NewSMSService(client *Client) *SMSService {
	return &SMSService{client: client}
}

