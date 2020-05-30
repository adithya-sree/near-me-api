package handler

import (
	"net/http"
	"time"
)

var startTime time.Time

//UptimeResponse response object for uptime
type UptimeResponse struct {
	StartedTime string        `json:"start-time"`
	Uptime      time.Duration `json:"uptime"`
}

func init() {
	startTime = time.Now()
}

//Base handler
func (h *Handler) Base(w http.ResponseWriter, r *http.Request) {
	respondJSON(w, http.StatusOK, Response{Message: "nearme-api is up"})
}

//Running handler
func (h *Handler) Running(w http.ResponseWriter, r *http.Request) {
	respondJSON(w, http.StatusOK, Response{Message: "running"})
}

//Uptime handler
func (h *Handler) Uptime(w http.ResponseWriter, r *http.Request) {
	respondJSON(w, http.StatusOK, UptimeResponse{
		StartedTime: startTime.Format("2006.01.02 15:04:05"),
		Uptime:      time.Since(startTime),
	})
}
