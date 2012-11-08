package server

import (
    "fmt"
    "log"
    "net"
    "io"
    "bufio"
)
type ServerBranch struct {    
    protocol string
    port int
    quit chan bool
    status chan string
}


type Server struct {
    branches map[int]*ServerBranch

}

func NewServer() *Server {
    server := new(Server)
    server.branches = make(map[int]*ServerBranch)
    return server
}

func (self *Server) Start(protocol string, port int) (ok bool, err error){
    self.branches[port] = &ServerBranch{protocol, port, make(chan bool), make(chan string, self.chanBuffer)}
    listener, err = net.Listen(protocol, ":"+strconv.Itoa(port))
    go listen(listener, self.branches[port].quit, self.branches[port].status)
}

func (self *Server) Stop(port int) (ok bool, err error) {
}

func (self *Server) StopAll() (ok bool, err error) {
}

func(self *Server) listen(listener net.Listener, quit chan bool, status chan string) (err error) {
    status <- fmt.Sprintf("Starting listener")
    for {
        select {
            case <-quit:
                break
            default:
                continue
        }
        status <- fmt.Sprintf("Listening")
        if conn, err := listener.Accept(); err != nil {
            status <- err.Error()
            break
        }
        status <- fmt.Sprintf("Connection accepted: %s %s", conn.RemoteAddr().Network(), conn.RemoteAddr().String())
        go handle(conn, status)
    }
    
}

func (self *Server) handle(conn net.Conn, status chan string) (err error) {
    //read query
    //process query
    //return answer to query
    //close connection
    conn.Close()
    return
    
}