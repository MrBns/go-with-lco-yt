package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:300/learn?coursename=react&paymentid=29mEabTEx39METN88"

func main() {

	customURL := url.URL{
		Scheme:   "https",
		Host:     "lco.dev",
		Path:     "/tutcss",
		RawQuery: "user=hitesh",
	}
	myUrl := &customURL

	customURL2 := &url.URL{
		Scheme:   "https",
		Host:     "lco.dev",
		Path:     "/tutcss",
		RawQuery: "user=hitesh",
	}
	myUrl2 := customURL2

	fmt.Printf("%p \n", customURL2)
	fmt.Printf("%p \n", myUrl2)
	fmt.Printf("%p \n", myUrl)
	fmt.Printf("%p \n", *&myUrl)
	fmt.Printf("%p \n", &customURL)

}
