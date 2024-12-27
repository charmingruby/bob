package endpoint

import (
	"errors"
	"fmt"
	"net/http"

	"{{ .SourcePath }}/{{ .Module }}/core/service"
	"{{ .SourcePath }}/{{ .Module }}/transport/rest/dto/response"
	"{{ .SourcePath }}/{{ .Module }}/transport/rest/dto/request"
	"{{ .SourcePath }}/shared/custom_err/core_err"
	"{{ .SourcePath }}/shared/transport/rest"
)

func (e *Endpoint) make{{ .Name }}Handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req, err := rest.ParseRequest[request.{{ .Name }}Request](r)
		if err != nil {
			rest.BadRequestErrorResponse[any](w, err.Error())
			return
		}

		res, err := e.service.{{ .ServiceName }}(service.{{ .ServiceName }}Params{
			Name: req.Name,
		})

		if err != nil {
			var notFoundErr *core_err.ResourceNotFoundErr
			if errors.As(err, &notFoundErr) {
				rest.NotFoundErrorResponse[any](w, err.Error())
				return
			}

			rest.InternalServerErrorResponse[any](w)
			return
		}

		rest.OKResponse[response.{{ .Name }}Response](w, "", response.{{ .Name }}Response{
			Greeting: fmt.Sprintf("Long time no see! `%s` was managed.", res.ID),
		})
	}
}
