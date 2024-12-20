package endpoint

import (
	"net/http"

	"{{ .SourcePath }}/{{ .Module }}/transport/rest/dto/response"
	"{{ .SourcePath }}/{{ .Module }}/transport/rest/dto/request"
	"{{ .SourcePath }}/shared/transport/rest"
)

func (e *Endpoint) make{{ .Name }}Handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req, err := rest.ParseRequest[request.PingRequest](r)
		if err != nil {
			rest.BadRequestErrorResponse[any](w, err.Error())
			return
		}

		rest.OkResponse[response.PingResponse](w, "success", response.PingResponse{
			Name: req.Name,
		})
	}
}
