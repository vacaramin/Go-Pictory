package main

import (
	"log"
	"net/http"
)

const (
	WebsiteURL      = "https://app.pictory.ai"
	WebsiteURLHelp  = "https://help.pictory.ai"
	WebsiteURLH     = "https://h.pictory.ai"
	WebsiteURLAPI   = "https://api.pictory.ai"
	StorageFilePath = "./storage.txt"
)

func initRequest(w http.ResponseWriter, r *http.Request) {
	response := makeRequest(WebsiteURL + r.URL.Path)

	responseBody := response.Body
	responseInfo := response.Header.Get("responseInfo")

	response.Header.Set("content-type", "text/html")
	response.StatusCode = 200

	if response.Header.Get("content-encoding") != "" {
		if response.Header.Get("content-encoding") == "gzip" {
			// Implement gzdecode if necessary
		} else {
			w.Header().Set("Content-Encoding", "")
		}
	}

}

func makeRequest(url string) *http.Response {
	request, err := http.NewRequest("GET", WebsiteURL, nil)
	if err != nil {
		log.Println("failed to make request")
	}
	var headersToRemove = []string{
		"Accept-Ranges",
		"Content-Length",
		"Keep-Alive",
		"Connection",
		"content-encoding",
		"content-type",
		"strict-transport-security",
		"x-frame-options",
		"server",
		"date",
		"via",
		"x-xss-protection",
		"cache-control",
		"alt-svc",
		"expires",
		"transfer-encoding",
		"last-modified",
		"etag",
	}
	for _, header := range headersToRemove {
		request.Header.Del(header)
	}
	return &http.Response{}
}

func md5Hash(s string) string {
	// Implement md5 hashing logic
	return ""
}

func proxifyHtml(result string) string {
	// Implement proxifyHtml logic
	return ""
}

func proxify(result string) string {
	// Implement proxify logic
	return ""
}

func getCurrentURL() string {
	// Implement getCurrentURL logic
	return ""
}
