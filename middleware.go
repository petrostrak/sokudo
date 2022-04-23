package sokudo

import "net/http"

func (s *Sokudo) SessionLoad(next http.Handler) http.Handler {
	return s.Session.LoadAndSave(next)
}
