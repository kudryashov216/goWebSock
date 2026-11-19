package data

import (
	"strings"

	"github.com/kudryashov216/goWebSock/internal/websock"
)

func HeadersIsValid(headers string) bool {

	var countHeadings uint8 = 0

	headersMap := GetHeadersToMap(headers)

	for key, value := range headersMap {

		switch key {
		case "GET":

			countHeadings++

			if value == "null" {
				return false
			}

		case "Host":

			countHeadings++

			if value == "null" {
				return false
			}

		case "Connection":

			countHeadings++

			if value != "Upgrade" {
				return false
			}

		case "Upgrade":

			countHeadings++

			if value != "websocket" {
				return false
			}

		case "Sec-WebSocket-Key":

			countHeadings++

			if value == "null" {
				return false
			}

		case "Sec-WebSocket-Version":

			countHeadings++

			if value != websock.CurrentVersion() {
				return false
			}

		}

	}

	if countHeadings < websock.NUMBERHEADINGS || countHeadings > websock.NUMBERHEADINGS {
		return false
	}

	return true

}

func GetHeadersToMap(headers string) map[string]string {

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

			headersMap[strings.TrimRight(keyValue[0], ":")] = keyValue[1]
		}

	}

	return headersMap

}
