package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Ashwani00002/http_client/gohttp"
)

var (
	httpClient = gohttp.New()
)

func main() {

	headers := make(http.Header)
	headers.Set("Authorization", "Bearer ABC-123")

	response, err := httpClient.Get("https://api.github.com", headers)
	if err != nil {
		panic(err)
	}
	fmt.Println(response.StatusCode)

	bytes, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(bytes))

}
