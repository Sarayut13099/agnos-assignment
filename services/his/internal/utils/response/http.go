package response

import (
	"encoding/json"
	"net/http"
)

func JSON(w http.ResponseWriter, status int, resp Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(resp)
}

func Success(w http.ResponseWriter, data interface{}) {
	JSON(w, http.StatusOK, Response{
		Success: true,
		Code:    "SUCCESS",
		Message: "request successful",
		Data:    data,
	})
}

func Error(w http.ResponseWriter, err error) {
	appErr, ok := err.(*AppError)
	if !ok {
		appErr = NewAppError(
			http.StatusInternalServerError,
			"INTERNAL_SERVER_ERROR",
			"internal server error",
		)
	}

	JSON(w, appErr.HTTPStatus, Response{
		Success: false,
		Code:    appErr.Code,
		Message: appErr.Message,
		Errors:  appErr.Errors,
	})
}
