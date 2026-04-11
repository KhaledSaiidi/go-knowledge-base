package main

import "fmt"

func main() {
	s := newServer(withTLS, withMaxConn(200), withID("pg-cluster-1"))
	fmt.Printf("Server: %+v\n", s)
}
