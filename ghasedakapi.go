package ghasedakapi

type GhasedakApi struct {
	SMS *SMSService
	Account *AccountService
	Voice *VoiceService
	Status *StatusService
	Receive	*ReceiveService
	Template *TemplateService
}

type SMSService struct {
	client *Client
}

type StatusService struct {
	client *Client
}

type VoiceService struct {
	client *Client
}
type AccountService struct {
	client *Client
}
type ReceiveService struct {
	client *Client
}
type TemplateService struct {
	client *Client
}


func New(apikey string) *GhasedakApi {
	client := NewClient(apikey)
	return NewWithClient(client)
}

func NewWithClient(client *Client) *GhasedakApi {
	obj := &GhasedakApi{}
	obj.Account =NewAccountService(client)
	obj.SMS = NewSMSService(client)
	obj.Voice = NewVoiceService(client)
	obj.Receive=NewReceiveService(client)
	obj.Status=NewStatusService(client)
	obj.Template=NewTemplateService(client)
	return obj
}

func NewTemplateService(client *Client) *TemplateService {
	return &TemplateService{client: client}
}
func NewReceiveService(client *Client) *ReceiveService {
	return &ReceiveService{client: client}
}
func NewStatusService(client *Client) *StatusService {
	return &StatusService{client: client}
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

