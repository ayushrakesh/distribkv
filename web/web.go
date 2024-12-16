package web

import (
	"fmt"
	"net/http"

	"github.com/ayushrakesh/distribkv/db"
)

type Server struct {
	db *db.Database
}

func NewServer(db *db.Database) *Server {
	return &Server{db: db}
}

func (s *Server) GetHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	key := r.Form.Get("key")
	val, err := s.db.GetKey(key)
	fmt.Fprintf(w, "Value = %q , error = %v", val, err)
}

func (s *Server) SetHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	key := r.Form.Get("key")
	val := r.Form.Get("value")
	err := s.db.SetKey(key, []byte(val))
	fmt.Fprintf(w, "Error = %v", err)
}
