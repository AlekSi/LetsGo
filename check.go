package main

import (
	"bufio"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	files, _ := ioutil.ReadDir("lets-go")
	used := make(map[string]bool)
	for _, f := range files {
		used["lets-go/"+f.Name()] = false
	}

	f, _ := os.Open("lets-go.slide")
	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}

		parts := strings.Split(strings.TrimSpace(line), " ")
		if len(parts) > 1 && ((parts[0] == ".image") || (parts[0] == ".code") || (parts[0] == ".play")) {
			_, ok := used[parts[1]]
			if !ok {
				log.Printf("%s not found", parts[1])
			}
			used[parts[1]] = true
		}
	}

	for f, u := range used {
		if !u {
			log.Printf("%s not used", f)
		}
	}
}
