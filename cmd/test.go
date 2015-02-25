package main

import (
	"fmt"
	"net/url"
)

func main() {

	var query = "filter[where][abc]=123&filter[limit]=5&filter[limit]=7"

	results, _ := url.ParseQuery(query)

	for k, v := range results {
		fmt.Print(k + " :: ")
		fmt.Println(v)
	}

}
