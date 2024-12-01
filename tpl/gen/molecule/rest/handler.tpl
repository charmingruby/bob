package endpoint

import (
	"fmt"
	"net/http"

	"github.com/charmingruby/txgo/internal/giftshop/transport/rest/dto/request"
	"github.com/charmingruby/txgo/internal/shared/transport/rest"
)

func (e *Endpoint) createWalletHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
			req, err := rest.ParseRequest[request.CreateWalletRequest](*e.validator, r)
			if err != nil {
				rest.BadRequestErrorResponse[any](w, err.Error())
				return
			}
        
		rest.CreatedResponse[any](w, "wallet")
	}
}
