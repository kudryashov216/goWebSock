package data

import (
	"strings"

	"github.com/kudryashov216/goWebSock/internal/websock"
)

func HeadersIsValid(headers string) (map[string]string, bool) {

	var countHeadings uint8 = 0

	headersMap := GetHeadersToMap(headers)

	for key, value := range *headersMap {

		key = strings.TrimSpace(key)
		value = strings.TrimSpace(value)

		switch key {
		case "GET":

			countHeadings++

			if value == "null" {
				return nil, false
			}

		case "Host":

			countHeadings++

			if value == "null" {
				return nil, false
			}

		case "Connection":

			countHeadings++

			if value != "Upgrade" {
				return nil, false
			}

		case "Upgrade":

			countHeadings++

			if value != "websocket" {
				return nil, false
			}

		case "Sec-WebSocket-Key":

			countHeadings++

			if value == "null" {
				return nil, false
			}

		case "Sec-WebSocket-Version":

			countHeadings++

			if value != websock.CurrentVersion() {
				return nil, false
			}

		}

	}

	if countHeadings < websock.NUMBERHEADINGS || countHeadings > websock.NUMBERHEADINGS {
		return nil, false
	}

	return *headersMap, true

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
