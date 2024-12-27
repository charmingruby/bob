package {{ .ExchangePackage }}

type {{ .Name }}{{ .Exchange }} struct {
	ID string `json:"name" validate:"required"`
}
