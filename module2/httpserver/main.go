package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	http.HandleFunc("/header", headers)
	http.HandleFunc("/healthz", healthz)

	http.ListenAndServe(":8001", nil)
}

func healthz(w http.ResponseWriter, r *http.Request) {
	httpState := http.StatusOK

	defer func() {
		requestInfo(httpState, r)
	}()
}

func headers(w http.ResponseWriter, r *http.Request) {
	httpState := http.StatusOK
	defer func() {
		requestInfo(httpState, r)
	}()
	v := os.Getenv("VERSION")
	w.Header().Set("VERSION", v)
	for name, headers := range r.Header {
		for _, header := range headers {
			w.Header().Add(name, header)
		}
	}

}

func requestInfo(httpState int, r *http.Request) {
	cIP := clientIP(r)
	path := r.URL
	fmt.Printf("[%s] [%d] [%s] client ip: %s\n", time.Now(), httpState, path, cIP)
}

//  参照 gin.clientIP()
func clientIP(r *http.Request) string {
	cIP := r.Header.Get("X-Forwarded-For")
	cIP = strings.TrimSpace(strings.Split(cIP, ",")[0])
	if cIP == "" {
		cIP = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	}
	if cIP != "" {
		return cIP
	}

	if addr := r.Header.Get("X-Appengine-Remote-Addr"); addr != "" {
		return addr
	}

	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}

	return ""
}
