package customerrors

import "fmt"

type JSONError struct {
	Field string
	Tag   string
}

func (jsonError *JSONError) Error() string {
	return jsonError.ToString()
}

func (jsonError *JSONError) ToString() string {
	if jsonError.Tag == "required" {
		return fmt.Sprintf("Brakuje wymaganego pola '%s'", jsonError.Field)
	}

	return fmt.Sprintf("%s - %s", jsonError.Field, jsonError.Tag)
}

func (jsonError *JSONError) ToJSON() ErrorResponse {
	return ToJSON(jsonError)
}
