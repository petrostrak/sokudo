package sokudo

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"path"
	"path/filepath"
)

func (s *Sokudo) WriteJson(w http.ResponseWriter, status int, data interface{}, headers ...http.Header) error {
	out, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_, err = w.Write(out)
	if err != nil {
		return err
	}

	return nil
}

func (s *Sokudo) WriteXML(w http.ResponseWriter, status int, data interface{}, headers ...http.Header) error {
	out, err := xml.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}

	w.Header().Set("Content-Type", "application/xml")
	w.WriteHeader(status)

	_, err = w.Write(out)
	if err != nil {
		return err
	}

	return nil
}

func (s *Sokudo) DownloadFile(w http.ResponseWriter, r *http.Request, pathToFile, fileName string) error {
	fullPath := path.Join(pathToFile, fileName)
	fileToServe := filepath.Clean(fullPath)

	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; file=\"%s\"", fileName))
	http.ServeFile(w, r, fileToServe)

	return nil
}

func (s *Sokudo) Error404(w http.ResponseWriter, r *http.Request) {
	s.ErrorStatus(w, http.StatusNotFound)
}

func (s *Sokudo) Error500(w http.ResponseWriter, r *http.Request) {
	s.ErrorStatus(w, http.StatusInternalServerError)
}

func (s *Sokudo) Error401(w http.ResponseWriter, r *http.Request) {
	s.ErrorStatus(w, http.StatusUnauthorized)
}

func (s *Sokudo) Error403(w http.ResponseWriter, r *http.Request) {
	s.ErrorStatus(w, http.StatusForbidden)
}

func (s *Sokudo) ErrorStatus(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}
