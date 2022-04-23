package sokudo

import "net/http"

func (s *Sokudo) SessionLoad(next http.Handler) http.Handler {
	s.InfoLog.Println("SessionLoad called")
	return s.Session.LoadAndSave(next)
}
