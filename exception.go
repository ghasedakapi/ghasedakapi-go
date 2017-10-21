package ghasedakapi

type ConnectionError struct {
	Status  int
	Message string
	Err     error
}

type ApiError struct {
	Status  int
	Message string
	Err     error
}
