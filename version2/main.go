package main

import (
	"net/http"
	"os"
	"net"
	"html/template"
	"path/filepath"
	"fmt"
)

type Content struct  {
	IP_Remote string
	Hostname string
	IP_Self string
}


// Obtained these functions from various sources

func main() {
	http.HandleFunc("/", ReqHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func ReqHandler(w http.ResponseWriter, r *http.Request) {

	content := Content{IP_Remote: getip(r),
	                   Hostname: gethn(),
	                   IP_Self: getmyip()  }

	ex, _ := os.Executable()
	exPath := filepath.Dir(ex)
	w.Header().Add("Content-Type", "text/html; charset=utf-8")
	t, err := template.ParseFiles(exPath + "/template.html")
	if err != nil {
		fmt.Fprintf(w, "Unable to server from template!")
	}
//	fmt.Printf("%+v\n", content)
	t.Execute(w, content)
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

