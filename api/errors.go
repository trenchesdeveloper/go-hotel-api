package api

type Error struct {
	Message string `json:"message"`
	Code   int    `json:"code"`

}

func (e Error) Error() string {
	return e.Message
}

func NewError(code int, message string) Error {
	return Error{
		Message: message,
		Code:   code,
	}
}

func ErrInvalidID() Error {
	return NewError(400, "invalid id")
}

