package pages

import (
	"net/http"
	"strings"
	"time"
)

/*
	RateLimiter.go is a basic implementation of a HTTP request rate limiter builtin
	for this golang http service. it will allow us to modify and change what triggers
	and when it resets
*/

// limiter is the amount of requests which can be made from an ip every x amount of time
const limiter int = 200

// requests will be recached every X amount of time
var requests []string = make([]string, 0)

// Cleaner is a self contained goroutine which will cleanup all requests every X amount of time
func Cleaner(x int) {
	ticker := time.NewTicker(time.Duration(x) * time.Second)
	for recv := range ticker.C {
		requests = make([]string, 0)
		_ = recv
		continue
	}
}

// AddRequest will append another request to the array
func AddRequest(request *http.Request) bool {
	sender := strings.Join(strings.Split(request.RemoteAddr, ":")[:strings.Count(request.RemoteAddr, ":")], ":")

	appears := 0

	for _, element := range requests {
		if !strings.EqualFold(element, sender) {
			continue
		}

		appears++
	}

	requests = append(requests, sender)
	return appears < limiter
}
