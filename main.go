package main

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"strings"
)

type Request struct {
	Method          string
	URL             string
	ProtocolVersion string
	Headers         map[string]string
}

var ValidMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE"}

func parseRequestLine(line string, req *Request) error {
	splitted := strings.Split(strings.Trim(line, "\n\r"), " ")

	if len(splitted) != 3 {
		return errors.New("parseRequestLine error")
	}

	req.Method = splitted[0]
	req.URL = splitted[1]
	req.ProtocolVersion = splitted[2]

	return nil
}

func parseHeaderLine(line string, req *Request) error {
	splitted := strings.Split(strings.Trim(line, "\n\r"), ": ")

	if req.Headers == nil {
		req.Headers = make(map[string]string)
	}

	req.Headers[splitted[0]] = splitted[1]

	return nil
}

func handle(conn net.Conn) error {
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

	response := "HTTP/1.1 200 OK\r\n" +
		"Content-Type: text/plain\r\n" +
		"Content-Length: 13\r\n" +
		"\r\n" +
		"Hello, World!\n"

	_, err := conn.Write([]byte(response))

	if err != nil {
		return fmt.Errorf("error writing response: %v", err)
	}

	return nil
}

func main() {
	server, err := net.Listen("tcp", ":8080")

	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer server.Close()

	fmt.Println("Server started on port 8080")

	for {
		conn, err := server.Accept()

		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		go handle(conn)
	}
}
