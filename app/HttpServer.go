package app

import (
	"fmt"
	"net"
	"strings"
)

type HttpStatus string

const (
	OK          HttpStatus = "200 OK"
	BAD_REQUEST HttpStatus = "400 Bad Request"
)

type HttpResponse struct {
	Status HttpStatus
	Body   string
}
type Route struct {
	Method string
	Path   string
	Action func() HttpResponse
}

type Server struct {
	Routes []Route
}

func (s *Server) Run() {
	server, err := net.Listen("tcp", "localhost:9090")
	if err != nil {
		fmt.Println("Can't connect to tcp server: " + err.Error())
		return
	}
	defer server.Close()

	for {
		connection, err := server.Accept()
		if err != nil {
			fmt.Println("Can't Accept the request")
		}
		go handleRequest(connection, s.Routes)
	}
}

func handleRequest(conn net.Conn, routes []Route) {
	defer conn.Close()
	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		message := string(buffer[:n])

		if n > 0 {
			parseHTTP(message, conn, routes)
		}

		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
}

func parseHTTP(rq string, conn net.Conn, routes []Route) {
	fmt.Println(rq)
	if len(rq) <= 0 {
		return
	}
	bodyParts := strings.Split(rq, "\n")
	firstLine := bodyParts[0]

	detail := strings.Split(firstLine, " ")
	if len(detail) <= 2 {
		return
	}
	method := detail[0]
	path := detail[1]

	println(method, path)
	for _, route := range routes {
		if route.Method == method && route.Path == path {
			action := route.Action()
			response := "HTTP/1.1 " + string(action.Status)
			response += "\nContent-Type: text/html"
			response += "\n\n" + action.Body
			fmt.Println(response)
			conn.Write([]byte(response))
			conn.Close()
			return
		}
	}
}
