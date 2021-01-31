package main

// This is the rpc server that fronts slowdb

import (
	"fmt"
	"groupcache-db-experiment/api"
	"groupcache-db-experiment/slowdb"
	"net"
	"net/http"
	"net/rpc"
)

type Server struct {
	db *slowdb.SlowDB
}

func (s *Server) Get(args *api.Load, reply *api.ValueResult) error {
	data := s.db.Get(args.Key)
	reply.Value = string(data)
	return nil
}

func (s *Server) Set(args *api.Store, reply *api.NullResult) error {
	s.db.Set(args.Key, args.Value)
	*reply = 0
	return nil
}

func NewServer(db *slowdb.SlowDB) *Server {
	server := new(Server)
	server.db = db
	return server
}

func (s *Server) Start(port string) {
	err := rpc.Register(s)
	if err != nil {
		fmt.Println("dbserver rpc Register() failed!")
		return
	}

	rpc.HandleHTTP()
	l, e := net.Listen("tcp", port)
	if e != nil {
		fmt.Println("fatal")
	}

	err = http.Serve(l, nil)
	if err != nil {
		fmt.Println("dbserver http Serve() failed!")
		return
	}
}

func main() {
	db := slowdb.NewSlowDB()
	server := NewServer(db)
	fmt.Println("dbserver starting on localhost:9090")
	server.Start(":9090")
}
