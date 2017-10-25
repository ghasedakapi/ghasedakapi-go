package ghasedakapi

//Kavenegar ...
type GhasedakApi struct {
	SMS *SMSService
}
//MessageService ...
type SMSService struct {
	client *Client
}


func New(apikey string) *GhasedakApi {
	client := NewClient(apikey)
	return NewWithClient(client)
}


func NewWithClient(client *Client) *GhasedakApi {
	obj := &GhasedakApi{}
	obj.SMS = NewSMSService(client)
	return obj
}


func NewSMSService(client *Client) *SMSService {
	return &SMSService{client: client}
}

