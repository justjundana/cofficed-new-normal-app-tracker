package attendanceHandler

import "net/http"

type AttHandlerInterface interface {
	GetAttendances(w http.ResponseWriter, r *http.Request)
	CreateAttendance(w http.ResponseWriter, r *http.Request)
	UpdateAttendance(w http.ResponseWriter, r *http.Request)
}
