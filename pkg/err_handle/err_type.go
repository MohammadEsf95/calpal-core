package errhandle

type Error struct {
	Message string
	Detail  []map[string]interface{}
}

func (e Error) Error() string {
	return e.Message
}

func NewError(message string, details ...map[string]interface{}) Error {
	return Error{
		Message: message,
		Detail:  details,
	}
}
