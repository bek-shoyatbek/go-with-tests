package racer

import (
	"net/http"
	"time"
)

func Racer(aURL, bURL string) (winner string) {
	aDuration := measureResponseTime(aURL)
	bDuration := measureResponseTime(bURL)

	if aDuration < bDuration {
		return aURL
	}
	return bURL
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}
