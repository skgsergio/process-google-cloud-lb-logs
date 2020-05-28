package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type cdnlog struct {
	Request struct {
		Referer string `json:"referer"`
		IP      string `json:"remoteIp"`
	} `json:"httpRequest"`
	Timestamp string `json:"timestamp"`
}

func main() {
	fRequests, err := os.Create("requests.csv")
	if err != nil {
		panic(err)
	}
	defer fRequests.Close()

	_, err = fRequests.WriteString("timestamp,clientip,referer\n")
	if err != nil {
		panic(err)
	}

	for _, name := range os.Args[1:] {
		fmt.Fprintf(os.Stderr, "Processing file `%s`...\n", name)
		file, err := os.Open(name)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			var log cdnlog

			err := json.Unmarshal(scanner.Bytes(), &log)
			if err != nil {
				fmt.Fprintf(os.Stderr, "JSON parse error error: %v", err)
				continue
			}

			t, err := time.Parse(time.RFC3339Nano, log.Timestamp)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error parsing date: %v", err)
				continue
			}

			_, err = fRequests.WriteString(fmt.Sprintf("%d,%s,\"%s\"\n", t.Unix(), log.Request.IP, log.Request.Referer))
			if err != nil {
				panic(err)
			}
		}
	}
}
