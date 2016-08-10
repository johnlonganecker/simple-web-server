package parser

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func errz(err error) {
	if err != nil {
		panic(err)
	}
}

type RequestHeader struct {
	Method      string
	Path        string
	HttpVersion string
	Headers     map[string]string
	Body        string
}

func (r RequestHeader) String() string {
	result := r.Method + " " + r.Path + " " + r.HttpVersion + "\n"

	for key, value := range r.Headers {
		result = result + key + ": " + value + "\n"
	}

	result = result + "\n" + r.Body

	return result
}

func Parse(s *bufio.Scanner) RequestHeader {

	var (
		req      RequestHeader
		bodyRead int
	)

	parseBody := false
	bodyRead = 0

	req = RequestHeader{}
	req.Headers = make(map[string]string)

	for i := 0; ; i++ {

		if done := s.Scan(); done == false {
			break
		}
		line := s.Text()

		fmt.Println(line)

		// start line
		if i == 0 {
			words := strings.Split(line, " ")
			req.Method = words[0]
			req.Path = words[1]
			req.HttpVersion = words[2]
		} else {
			if string(line) == "" {
				parseBody = true
			}
			if parseBody {
				req.Body += line
				bodyRead += len(line) + 1
				contentLength, _ := strconv.Atoi(req.Headers["Content-Length"])
				if bodyRead >= contentLength {
					break
				}
			} else {
				words := strings.Split(line, ":")
				key := words[0]
				value := strings.TrimPrefix(words[1], " ")
				req.Headers[key] = value
			}
		}
	}

	return req
}
