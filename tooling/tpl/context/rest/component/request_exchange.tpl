package request

type {{ .Name }}{{ .Exchange }} struct {
	Name string `json:"name" validate:"required"`
}
