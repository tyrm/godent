package http

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type errResponse struct {
	Code   ErrCode `json:"errcode"`
	String string  `json:"error"`
}

func ReturnError(w http.ResponseWriter, code ErrCode, errStr string, vars ...interface{}) {
	l := logger.WithField("func", "ReturnError")

	newErr := errResponse{
		Code:   code,
		String: fmt.Sprintf(errStr, vars...),
	}

	httpStatus, ok := errCodeHTTPStatus[code]
	if ok {
		w.WriteHeader(httpStatus)
	} else {
		// fallback
		w.WriteHeader(http.StatusInternalServerError)
	}

	err := json.NewEncoder(w).Encode(newErr)
	if err != nil {
		l.Errorf("encoding error response: %s", err.Error())
	}
}
