package data

import (
	"net/http"
	"strings"

	"github.com/kudryashov216/goWebSock/internal/websock"
)

func HeadersIsValid(headers string) bool {

	var countHeadings uint8 = 0

	headersMap := GetHeadersToMap(headers)

	for key, value := range *headersMap {

		keyLow := strings.ToLower(key)
		keyLow = strings.TrimSpace(key)
		valueLow := strings.ToLower(value)
		valueLow = strings.TrimSpace(value)

		switch keyLow {
		case "get":

			countHeadings++

			if value == "null" {
				return false
			}

		case "host":

			countHeadings++

			if valueLow == "null" {
				return false
			}

		case "connection":

			countHeadings++

			if valueLow != "upgrade" {
				return false
			}

		case "upgrade":

			countHeadings++

			if valueLow != "websocket" {
				return false
			}

		case "sec-websocket-key":

			countHeadings++

			if valueLow == "null" {
				return false
			}

		case "sec-websocket-version":

			countHeadings++

			if valueLow != websock.CurrentVersion() {
				return false
			}

		}

	}

	return true

}

func GetHeadersToMap(headers string) *map[string]string {

	headersMap := make(map[string]string)

	array := strings.Split(headers, "\n")

	for _, s := range array {

		strings.TrimSpace(s)

		if s == "" {
			continue
		}

		keyValue := strings.SplitN(s, " ", 2)

		if len(keyValue) > 1 {

			if keyValue[0] == "GET" {
				keyValue[1] = strings.Split(strings.TrimSpace(keyValue[1]), " ")[1]
			}

			headersMap[strings.TrimRight(keyValue[0], ":")] = strings.Replace(keyValue[1], "\r", "", 1)
		}

	}

	return &headersMap

}

func HeaderToString(header *http.Header) string {

	var (
		readyString string
		countItems  uint16
	)

	for key, values := range *header {

		countItems = uint16(len(values) - 1)

		readyString += key + ": "

		for i, value := range values {

			readyString += value

			if uint16(i) < countItems {
				readyString += ","
			}

			if uint16(i) >= countItems {
				readyString += "\n"
			}

		}

	}

	return readyString

}

func KeyMapToLower(map_ *map[string]string) *map[string]string {

	newMap := make(map[string]string)

	for key, value := range *map_ {
		newMap[strings.ToLower(key)] = value
	}

	return &newMap

}
