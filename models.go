package go01

type RequestBody struct {
	Message string `json:"message"`
	Id      string `json:"id"`
}

type FailedMessage struct {
	Message string `json:"message"`
	Id      string `json:"id"`
	Error   error  `json:"error"`
}

type Response struct {
	FailedMessages          []FailedMessage
	SuccessfulMessagesCount int
	FailedMessagesCount     int
}
