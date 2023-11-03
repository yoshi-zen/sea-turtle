package myerrors

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/yoshi-zen/sea-turtle/backend/api/middlewares"
)

func ErrorHandler(w http.ResponseWriter, req *http.Request, err error) {
	var myErr *MyError
	if !errors.As(err, &myErr) {
		myErr = &MyError{
			ErrCode: Unknown,
			Message: "internal process failed",
			Err:     err,
		}
	}

	traceID := middlewares.GetTraceID(req.Context())
	log.Printf("[%d]error: %s\n", traceID, myErr)

	var statusCode int

	switch myErr.ErrCode {
	case NoData:
		statusCode = http.StatusNotFound
	case ReqDecodeFailed, BadParameter:
		statusCode = http.StatusBadRequest
	default:
		statusCode = http.StatusInternalServerError
	}

	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(myErr)
}
