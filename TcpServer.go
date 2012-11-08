package TcpServer

import (
    "log"
    "./server"
)
func main() {
    tcpServer := server.NewServer()
    tcpServer.Start("tcp", 1337)
}
