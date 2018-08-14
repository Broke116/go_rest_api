package utils

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/fatih/color"
)

// Log is a common function which used return the request parameters
func Log(method string, url *url.URL) {
	green := color.New(color.FgGreen).SprintfFunc()
	fmt.Printf("%s ", green(method))
	fmt.Print(url)
	fmt.Println()
}

// Error is used to show a common error message to in the response
func Error(w http.ResponseWriter, err error, status int) {
	if err != nil {
		http.Error(w, err.Error(), status)
		return
	}
}
