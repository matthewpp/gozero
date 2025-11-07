package errs

// APIError is a struct for handling API errors with HTTP status code
//
// swagger:model APIError
type APIError struct {
	httpStatus int
	code       string
	message    string
}

func newAPIError(httpStatus int, code, message string) APIError {
	return APIError{
		httpStatus: httpStatus,
		code:       code,
		message:    message,
	}
}

func (e APIError) Error() string {
	return e.code + " : " + e.message
}

func (e APIError) Code() string {
	return e.code
}

func (e APIError) Message() string {
	return e.message
}

func (e APIError) MarshalJSON() ([]byte, error) {
	// data := map[string]any{"error": e.code}
	// if e.message != "" {
	// 	data["message"] = e.message
	// }
	// return json.Marshal(data)

	return []byte(`{"error":"` + e.code + `","message":"` + e.message + `"}`), nil
}
