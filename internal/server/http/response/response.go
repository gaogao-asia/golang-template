package response

type ResponseBody struct {
	Data  interface{}        `json:"data,omitempty"`
	Error *ErrorResponseBody `json:"error,omitempty"`
}

type ErrorResponseBody struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Detail  string `json:"detail,omitempty"`
}
