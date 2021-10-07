package types

type ErrorResponse struct {
	// Error returned
	Error string `json:"error"`
	// Message returned with the error
	Message string `json:"message"`
}
