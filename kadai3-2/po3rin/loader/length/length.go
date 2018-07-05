package length

import (
	"fmt"
	"net/http"
	"strconv"
)

// Calc calcrate contents-length
func Calc() int {
	// get contents length
	res, err := http.Head("http://localhost:8080/5mqx.cif")
	if err != nil {
		fmt.Println(err)
	}
	maps := res.Header
	length, err := strconv.Atoi(maps["Content-Length"][0])
	if err != nil {
		fmt.Println(err)
	}
	return length
}
