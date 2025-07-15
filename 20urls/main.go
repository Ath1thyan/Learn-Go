package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://example.com:3000/account?user=demo&password=test1234"

func main() {
	fmt.Println("Handling URLs")
	fmt.Println(myurl)

	// parsing
	result, _ := url.Parse(myurl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)
	fmt.Println(result.Port())

	qparams := result.Query()
	fmt.Printf("Type of query params: %T\n", qparams)

	fmt.Println(qparams["user"])

	for _, val := range qparams {
		fmt.Println("Param is: ", val)
	}

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "example.com",
		Path:    "/account",
		RawPath: "user=demo",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)
}
