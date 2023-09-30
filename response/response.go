package response

type responseAPI struct {
	Meta meta `json:"meta"`
	Data any  `json:"data"`
}

type meta struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

func GenerateResponseAPI(statusCode int, Message string, data any) *responseAPI {
	meta := meta{
		StatusCode: statusCode,
		Message:    Message,
	}

	return &responseAPI{
		Meta: meta,
		Data: data,
	}
}
