package request

type {{ .ActionName }}Request struct {
	Name string `json:"name" validate:"required"`
}
