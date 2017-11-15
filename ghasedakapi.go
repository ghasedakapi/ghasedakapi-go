package ghasedakapi

type GhasedakApi struct {
	SMS *SMSService
	Account *AccountService
	Voice *VoiceService
}

type SMSService struct {
	client *Client
}
type VoiceService struct {
	client *Client
}
type AccountService struct {
	client *Client
}


func New(apikey string) *GhasedakApi {
	client := NewClient(apikey)
	return NewWithClient(client)
}


func NewWithClient(client *Client) *GhasedakApi {
	obj := &GhasedakApi{}
	obj.SMS = NewSMSService(client)
	obj.Account =NewAccountService(client)
	return obj
}

func NewSMSService(client *Client) *SMSService {
	return &SMSService{client: client}
}
func NewAccountService(client *Client) *AccountService {
	return &AccountService{client: client}
}

func NewVoiceService(client *Client) *VoiceService {
	return &VoiceService{client: client}
}

