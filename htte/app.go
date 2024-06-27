package htte

import (
	"bufio"
	"fmt"
	"net"
)

type Request struct {
	Method          string
	URL             string
	ProtocolVersion string
	Headers         map[string]string
}

type Configs struct {
	Address string
	Port    int
}

type App struct {
	Config Configs
	Router Router
}

var ValidMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE"}

func (app *App) handleConnection(conn net.Conn) error {
	defer conn.Close()
	reader := bufio.NewReader(conn)

	var req Request
	var step = "request-line"
	for {
		line, err := reader.ReadString('\n')

		if err != nil {
			return fmt.Errorf("error reading line: %v", err)
		}

		if step == "request-line" {
			err := parseRequestLine(line, &req)

			if err != nil {
				return err
			}

			step = "header-line"
			continue
		}

		if step == "header-line" && line != "\r\n" {
			err := parseHeaderLine(line, &req)

			if err != nil {
				return err
			}
			continue
		}

		if line == "\r\n" {
			break
		}
		// does not parse body
	}

	var response string

	handler := app.Router.match(req.URL, req.Method)

	if handler == nil {
		response = "HTTP/1.1 404 Not Found\r\n" +
			"Content-Type: text/plain\r\n" +
			"Content-Length: 10\r\n" +
			"\r\n" +
			"Not Found!\n"
	} else {
		response = "HTTP/1.1 200 OK\r\n" +
			"Content-Type: text/plain\r\n" +
			"Content-Length: 13\r\n" +
			"\r\n" +
			"Hello, World!\n"

		handler(req)
	}

	_, err := conn.Write([]byte(response))

	if err != nil {
		return fmt.Errorf("error writing response: %v", err)
	}

	return nil
}

func (app *App) Serve() error {
	server, err := net.Listen("tcp", fmt.Sprint(app.Config.Address, app.Config.Port))

	if err != nil {
		return err
	}

	defer server.Close()

	fmt.Println("Server started on port", app.Config.Port)

	for {
		conn, err := server.Accept()

		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		go app.handleConnection(conn)
	}
}

func New(config Configs) App {
	var router = Router{routes: []Route{}, config: RouterConfig{}}

	var app = App{Config: config, Router: router}

	return app
}
