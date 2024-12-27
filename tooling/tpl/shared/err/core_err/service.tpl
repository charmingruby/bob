package core_err

import "fmt"

type ResourceNotFoundErr struct {
	Message string `json:"message"`
}

func NewResourceNotFoundErr(resource string) *ResourceNotFoundErr {
	return &ResourceNotFoundErr{
		Message: fmt.Sprintf("%s not found", resource),
	}
}

func (e *ResourceNotFoundErr) Error() string {
	return e.Message
}
