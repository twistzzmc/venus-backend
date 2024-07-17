package customerrors

type ICustomError interface {
	ToString() string
	ToJSON() ErrorResponse
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func ToJSON(customError ICustomError) ErrorResponse {
	return ErrorResponse{
		Error: customError.ToString(),
	}
}
