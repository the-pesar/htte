package htte

import (
	"errors"
	"strings"
)

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
