package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

var (
	defBasepath       = "man_pages"
	defSections       = []int{1, 4, 8}
	availableManPages = map[int][]string{}
)

func manHandler(w http.ResponseWriter, r *http.Request) {
	sections := defSections
	if secs, ok := r.URL.Query()["sections"]; ok {
		sections = []int{}
		for _, sec := range secs {
			num, _ := strconv.Atoi(sec)
			sections = append(sections, num)
		}

	}
	randomPage, err := GetContentOfRandomManPage(sections, availableManPages)
	if err != nil {
		fmt.Fprint(w, "Error... Sorry...")
		return
	}
	fmt.Fprintf(w, randomPage)
}

func initServer() {
	basePath := defBasepath
	if len(os.Args) > 1 {
		basePath = os.Args[1]
	}
	availableManPages = LoadAvailableManpages(basePath)
}

func main() {
	initServer()
	http.HandleFunc("/random", manHandler)
	http.ListenAndServe(":8080", nil)
}
