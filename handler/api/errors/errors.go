package errors

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Success bool   `json:"success"`
}

func (e *Error) Error() string {
	return e.Message
}

func New(text string) error {
	return &Error{
		Message: text,
	}
}
