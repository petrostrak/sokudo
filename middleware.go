package sokudo

import (
	"net/http"
	"strconv"

	"github.com/justinas/nosurf"
)

func (s *Sokudo) SessionLoad(next http.Handler) http.Handler {
	s.InfoLog.Println("SessionLoad called")
	return s.Session.LoadAndSave(next)
}

func (s *Sokudo) NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	secure, _ := strconv.ParseBool(s.config.cookie.secure)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   secure,
		SameSite: http.SameSiteStrictMode,
		Domain:   s.config.cookie.domain,
	})

	return csrfHandler
}
