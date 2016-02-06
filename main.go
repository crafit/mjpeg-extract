package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
)

func main() {
	resp, err := http.Get(os.Args[1])
	if err != nil {
		panic(err)
	}

	contentType := resp.Header.Get("Content-Type")
	boundaryRe := regexp.MustCompile(`^multipart/x-mixed-replace;boundary=(.*)$`)
	boundary := boundaryRe.FindStringSubmatch(contentType)[1]
	buffer := bufio.NewReader(resp.Body)

	delimiter := fmt.Sprintf("--%s\r\n", boundary)
	data := make([]byte, 0)
	for i := 0; i < 4; i++ {
		buffer.ReadBytes('\n')
	}

	for {
		line, _ := buffer.ReadBytes('\n')
		found := bytes.HasSuffix(line, []byte(delimiter))
		if found == true {
			data = append(data, line[:(len(line)-len(delimiter))]...)
			ioutil.WriteFile("test.jpg", data, 0644)
			break
		} else {
			data = append(data, line...)
		}
	}
	resp.Body.Close()
}
