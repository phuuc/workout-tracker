package errors

import "net/http"

const (
	ErrInternalServerCode  = http.StatusInternalServerError
	ErrInternalServer      = "internal_server_error"
	ErrParseRequestFail    = "parse request fail: %v"
	ErrRegisterRequestFail = "fail register request : %v"
)
