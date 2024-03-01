package main

import (
	"fmt"
	"go-http-server/app"
)

func helloService() app.HttpResponse {
	fmt.Println("Hello khobi")
	return app.HttpResponse{Status: app.OK, Body: `Hello response`}
}

func helloService2() app.HttpResponse {
	fmt.Println("Hello2 kgobi")
	return app.HttpResponse{}
}

func main() {
	var routes = []app.Route{
		{Method: "GET", Path: "/hello", Action: helloService},
		{Method: "GET", Path: "/hello2", Action: helloService2},
	}
	server := app.Server{Routes: routes}
	server.Run()

}
