package errhandle

type Error struct {
	Message string
	Detail  map[string]interface{}
}

func (e Error) Error() string {
	return e.Message
}

func NewError(message string, details ...interface{}) Error {
	m := make(map[string]interface{})
	m["details"] = details

	return Error{
		Message: message,
		Detail:  m,
	}
}
