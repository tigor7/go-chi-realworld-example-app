package httputil

import (
	"encoding/json"
	"net/http"
)

type validator interface {
	Validate() error
}
type apiErrorResponse struct {
	Message string `json:"message"`
}

type validationError struct {
	raw error
}

type validationErrorResponse struct {
	Errors any `json:"errors"`
}

func (e *validationError) Error() string {
	return e.raw.Error()
}

func BindAndValidate(r *http.Request, request validator) error {
	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		return err
	}
	if err := request.Validate(); err != nil {
		return &validationError{
			raw: err,
		}
	}
	return nil
}

func Respond(w http.ResponseWriter, status int, body any) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(body)
}

func RespondErr(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	if e, ok := err.(*validationError); ok {
		errs, _ := json.Marshal(e.raw)
		json.NewEncoder(w).Encode(validationErrorResponse{
			Errors: json.RawMessage(errs),
		})
		return
	}
	json.NewEncoder(w).Encode(apiErrorResponse{
		Message: err.Error(),
	})
}
