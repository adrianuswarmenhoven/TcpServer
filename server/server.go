package server

import (
    "fmt"
    "log"
    "net"
    "io"
    "bufio"
)
type ServerBranch struct {
    port int
    protocol string
    listenerChan chan string
    handlerChan chan string
    connChan chan net.Conn
}


type Server struct {
    branches map[int]*ServerBranch

}

func NewServer() *Server {
    return new(Server)
}

func (self *Server) Start(protocol string, port int) (ok bool, err error){
}

func (self *Server) Stop(port int) (ok bool, err error) {
}

func (self *Server) StopAll() (ok bool, err error) {
}
