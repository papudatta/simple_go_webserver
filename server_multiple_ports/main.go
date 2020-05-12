package main

import (
	"encoding/json"
	"net/http"
	"os"
    "net"
    "fmt"
)

// Obtained these functions from various sources

func main() {
    serverMuxA := http.NewServeMux()
    serverMuxA.HandleFunc("/", helloworld)

    serverMuxB := http.NewServeMux()
    serverMuxB.HandleFunc("/", ReqHandler)
    go func() {
        http.ListenAndServe(":8080", serverMuxB)
    }()
    http.ListenAndServe(":8081", serverMuxA)
    
}
func helloworld(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "helloworld!")
}

func ReqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	resp, _ := json.Marshal(map[string]string{
		"your_ip": getip(r), "my_hostname": gethn(), "my_ip": getmyip(),
	})
	w.Write(resp)
}

// GetIP gets a requests IP address by reading off the forwarded-for
// header (for proxies) and falls back to use the remote address.
func getip(r *http.Request) string {
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		return forwarded
	}
	return r.RemoteAddr
}

func gethn() string {
	hn, err := os.Hostname()
	if err != nil {
		return "There was an error retrieving the hostname"
		panic(err)
    }
	return hn
}

func getmyip() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// Filter loopback
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}
