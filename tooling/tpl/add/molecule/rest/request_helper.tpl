package rest

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"{{ .SourcePath }}/shared/validation"
)

func ParseRequest[T any](request *http.Request) (*T, error) {
	body, err := io.ReadAll(request.Body)
	defer request.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("unable to read request body: %v", err)
	}

	var req T
	if err := json.Unmarshal(body, &req); err != nil {
		return nil, fmt.Errorf("unable to unmarshal request body: %v", err)
	}

	if err := validation.ValidateStructByTags(req); err != nil {
		return nil, fmt.Errorf("request validation failed: %v", err)
	}

	return &req, nil
}
