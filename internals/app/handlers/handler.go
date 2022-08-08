package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/wertick01/dclib/internals/pkg/logger"
)

func WrapError(w http.ResponseWriter, r *http.Request, err error) {
	WrapErrorWithStatus(w, err, http.StatusBadRequest, r)
}

func WrapErrorWithStatus(w http.ResponseWriter, err error, httpStatus int, r *http.Request) {
	var m = map[string]string{
		"result": "error",
		"data":   err.Error(),
	}

	res, _ := json.Marshal(m)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Add("Status", strconv.Itoa(httpStatus))
	logger.Errorer(r.RequestURI, err)

	fmt.Fprintln(w, string(res))
}

func WrapOK(w http.ResponseWriter, m map[string]interface{}) {
	res, _ := json.Marshal(m)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Add("Status", strconv.Itoa(http.StatusOK))
	fmt.Fprintln(w, string(res))
}
