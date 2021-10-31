package response

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func SuccessResponse() *Response {
	return &Response{
		200,
		"Operation Success",
	}
}

func CreatedResponse() *Response {
	return &Response{
		201,
		"Created Successfully",
	}
}

func BadRequestResponse() *Response {
	return &Response{
		400,
		"Bad Request",
	}
}

func NotFoundResponse() *Response {
	return &Response{
		404,
		"Not Found",
	}
}

func InternalServerErrorResponse() *Response {
	return &Response{
		500,
		"Internal Server Error",
	}
}
