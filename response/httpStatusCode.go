package response

const (
	ErrorCodeSuccess      = 2001 // Success
	ErrorCodeParamInvalid = 2003 // Email is invalid
)

// message
var msg = map[int]string{
	ErrorCodeSuccess:      "success",
	ErrorCodeParamInvalid: "email is invalid",
}
