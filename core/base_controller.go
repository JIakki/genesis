package core

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

type BaseController struct {
	StartedAt time.Time
}

func (ctrl *BaseController) Response(w http.ResponseWriter, r *http.Request, message []byte, statusCode int) {
	host, _ := os.Hostname()
	w.Header().Set("X-Server-Name", host)
	w.Header().Set("X-Response-Time", time.Now().Sub(ctrl.StartedAt).String())
	w.WriteHeader(statusCode)
	w.Write([]byte(message))
}

func (ctrl *BaseController) JSON(w http.ResponseWriter, r *http.Request, message interface{}) {
	w.Header().Set("Content-Type", "application/json")
	response, err := json.Marshal(message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ctrl.Response(w, r, response, 200)
}

func (ctrl *BaseController) InternalError(w http.ResponseWriter, r *http.Request, err error) {
	log.Println(err)
	ctrl.Response(w, r, []byte("Internal Error"), 500)
}

func (ctrl *BaseController) NotFound(w http.ResponseWriter, r *http.Request, message string) {
	ctrl.Response(w, r, []byte(message), 404)
}

func NewBaseController() *BaseController {
	return &BaseController{
		StartedAt: time.Now(),
	}
}
