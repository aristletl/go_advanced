package server

import (
	"context"
	"net/http"
	"week04/internal/biz"
	"week04/internal/data"
)

type Server struct {
	mux *http.ServeMux
}

func NewServer(m *data.DBModel) *Server {
	s := &Server{mux: http.NewServeMux()}

	s.mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		account := r.FormValue("account")
		res := biz.SayHello(context.Background(), m, account)
		w.Write([]byte(res))
	})
	return s
}

func (s *Server) GetMux() *http.ServeMux {
	return s.mux
}
