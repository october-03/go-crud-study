package dto

type DefaultErrorDto struct {
	Message   string
	ErrorCode string
}

func NewError(message, errorcode string) DefaultErrorDto {
	return DefaultErrorDto{
		Message:   message,
		ErrorCode: errorcode,
	}
}
