package server_test

import (
	server "github.com/bimgithub/my_mod"
	"testing"
)

func TestServer_StartServer(t *testing.T) {
	s := server.Server{
		IP:   "10.0.0.41",
		Port: 7070,
	}
	s.StartServer()
}
