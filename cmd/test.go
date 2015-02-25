package main

import (
	"fmt"
	"github.com/thetruetrade/dbmetamodel"
	"net/url"
	//	"regexp"
)

func main() {

	var query = "filter[where][abc][gt]=123&filter[limit]=5&filter[limit]=7&filter[order]=abc ASC,somefiled DESC&filter[skip]=3"

	// re := regexp.MustCompile("(?i)filter\\s*\\[\\s*(?P<filterType>limit)\\s*\\]\\s*=\\s*(?P<limitTo>\\d*)")

	// if re.MatchString(query) {
	// 	n1 := re.SubexpNames()
	// 	r2All := re.FindAllStringSubmatch(query, -1)

	// 	for _, r2 := range r2All {
	// 		md := map[string]string{}
	// 		for i, n := range r2 {
	// 			fmt.Printf("%d. match='%s'\tname='%s'\n", i, n, n1[i])
	// 			md[n1[i]] = n
	// 		}
	// 		fmt.Printf("The names are  : %v\n", n1)
	// 		fmt.Printf("The matches are: %v\n", r2)
	// 		fmt.Printf("The filterType is %s\n", md["filterType"])
	// 		fmt.Printf("The limit is %s\n", md["limitTo"])
	// 	}
	// }
	results, _ := url.ParseQuery(query)

	parser := dbmetamodel.NewQueryStringUrlParser()

	//lq := &dbmetamodel.LimitQueryPart{}
	findQ := &dbmetamodel.FindQuery{}
	err := parser.ParseFindQueryString(findQ, results) //  ParseUrlLimitFilterPart(lq, results)

	if err != nil {
		fmt.Println(err)
	}

	if findQ.Limit != nil {
		fmt.Println(*findQ.Limit)
	}
	if findQ.Skip != nil {
		fmt.Println(*findQ.Skip)
	}

	if findQ.OrderBy != nil {
		for i := range findQ.OrderBy.OrderBy {
			fmt.Println(*findQ.OrderBy.OrderBy[i])
		}
	}
}
