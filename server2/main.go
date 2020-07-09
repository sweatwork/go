// Server2 is a minimal "echo" and counter server
package main 

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	// HandleFunc registers the handler function for the given pattern
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)

	/* ListenAndServe listens on the TCP network address addr and 
	then calls Serve with handler to handle requests on incoming connections.
	Accepted connections are configured to enable TCP keep-alives.
	The handler is typically nil, in which case the DefaultServeMux is used.
	ListenAndServe always returns a non-nil error.
	func ListenAndServe(addr string, handler Handler) error */ 
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}


// type Request - A Request represents an HTTP request received by a server or 
// to be sent by a client.
//  type ResponseWriter - A ResponseWriter interface is used by an HTTP handler to construct
//  an HTTP response.

// handler echoes the Path component of the requested URL
func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// counter echoes the number of calls so far
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}