package domain

import "net/http"

// ResponseError struct holds error info to send in response
type ResponseError struct {
	ErrorCode        string `json:"errorCode"`
	ErrorDescription string `json:"errorDescription"`
	Status           int    `json:"-"`
}

func (re ResponseError) Error() string {
	return re.ErrorCode
}

var (
	InvalidOrder       = ResponseError{"invalidOrderPayload", "invalid payload provided", http.StatusBadRequest}
	InvalidOrderID     = ResponseError{"invalidOrderId", "invalid order id provided", http.StatusBadRequest}
	InvalidPhoneNumber = ResponseError{"invalidPhoneNumber", "invalid phone number ", http.StatusBadRequest}
)
