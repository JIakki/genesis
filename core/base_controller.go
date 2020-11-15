package core

import (
	"encoding/json"
	"net/http"
)

type BaseController struct {
}

func (ctrl *BaseController) Response(w http.ResponseWriter, req *http.Request, message interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	response, err := json.Marshal(message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(response)
}

func (ctrl *BaseController) JSON(w http.ResponseWriter, req *http.Request, message interface{}) {
	ctrl.Response(w, req, message, 200)
}

func (ctrl *BaseController) InternalError(w http.ResponseWriter, req *http.Request, message interface{}) {
	ctrl.Response(w, req, message, 500)
}

func (ctrl *BaseController) NotFound(w http.ResponseWriter, req *http.Request, message interface{}) {
	ctrl.Response(w, req, message, 404)
}
