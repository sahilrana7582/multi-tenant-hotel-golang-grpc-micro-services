package routes

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/sahilrana7582/multi-tenant-hotel/pkg/errs"
)

func CustomHandler(f func(http.ResponseWriter, *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {

			status := http.StatusInternalServerError
			var sc interface{ HTTPStatus() int }
			if errors.As(err, &sc) {
				status = sc.HTTPStatus()
			}

			var resp *errs.AppError
			if appErr, ok := err.(*errs.AppError); ok {
				resp = appErr
			} else {
				resp = errs.Wrap(err, errs.ErrInternalServer.Code, errs.ErrInternalServer.Status)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(status)
			_ = json.NewEncoder(w).Encode(map[string]interface{}{
				"code":    resp.Code,
				"message": resp.Message,
			})
		}
	}
}
