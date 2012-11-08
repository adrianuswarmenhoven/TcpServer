package main
import (
  "log"
  "net"
  "bufio"
)
func main() {
  if conn, err := net.Dial("tcp", "localhost:1337"); err != nil {
    log.Fatal(err)
  }
  fmt.Fprintf(conn, "GET THIS\n")
  if response, err := bufio.NewReader(conn).ReadString('\n'); err != nil {
    log.Fatal(err)
  }
  fmt.Printf("Response: %s \n", response)
}