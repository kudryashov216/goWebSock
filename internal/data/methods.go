package data

import (
	"fmt"
	"strings"
)

func HeadersIsValid(headers string) bool {

	_ = getHeadersToMap(headers)

	return true

}

func getHeadersToMap(headers string) map[string]string {

	headersMap := make(map[string]string)

	array := strings.Split(headers, "\n")

	for _, s := range array {

		strings.TrimSpace(s)

		if s == "" {
			continue
		}

		fmt.Println(s)

	}

	return headersMap

}
