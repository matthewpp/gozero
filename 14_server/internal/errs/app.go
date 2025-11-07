package errs

var (
	// ErrNotFound indicates that the requested resource was not found.
	ErrNotFound = newAPIError(404, "not_found", "The requested resource was not found.")

	// ErrInvalidInput indicates that the input provided by the user is invalid.
	ErrInvalidInput = newAPIError(400, "invalid_input", "The input provided is invalid.")

	// ErrInvalidUserID indicates that the user ID provided is invalid.
	ErrInvalidUserID = newAPIError(400, "invalid_user_id", "The user ID provided is invalid.")

	// ErrInternalServer indicates that an internal server error occurred.
	ErrInternalServer = newAPIError(500, "internal_server_error", "An internal server error occurred.")
)
