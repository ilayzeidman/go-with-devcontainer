package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
)

const CTL_SOCKET_PATH = "./ctl.sock"

func main() {
    fmt.Println("hello world")

    os.Remove(CTL_SOCKET_PATH)
    http.HandleFunc("/ctl/terminate", func(w http.ResponseWriter, r *http.Request) {

    fmt.Fprintf(w, "Done\n")

	})

    server := http.Server{}

    unixListener, err := net.Listen("unix", CTL_SOCKET_PATH)
    if err != nil {
        panic(err)
    }
    go server.Serve(unixListener)
    fmt.Println("Server started and listening on", CTL_SOCKET_PATH)
    // Block forever
	select {}
    
}