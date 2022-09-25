package server

import "fmt"

type Server struct {
	IP   string
	Port int
}

func (s Server) StartServer() {
	fmt.Printf("Started server version 2 on ip %v and port %v", s.IP, s.Port)
}
