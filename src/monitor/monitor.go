package monitor

import (
    _ "expvar"
    "net"
    "os"
    "fmt"
    "net/http"
)


func Monitor() {
    go func () {
        sock, err := net.Listen("tcp", "0.0.0.0:8864")
        if err != nil {
            os.Exit(1)
        }
        fmt.Println("Monitor now available at port 8864")
        http.Serve(sock, nil)
    }()
}