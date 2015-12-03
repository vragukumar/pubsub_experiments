package main

import "os"
import "fmt"
import "time"
import "github.com/go-nanomsg"

func main() {
    var t [10000]time.Time
    sub, err := nanomsg.NewSubSocket()
    if err != nil {
        fmt.Println("Failed to open sub socket")
        return
    }
    ep, err := sub.Connect("ipc:///tmp/test.ipc")
    if err != nil {
        fmt.Println("Failed to connect to pub socket - ", ep)
        return
    }
    err = sub.Subscribe("")
    if err != nil {
        fmt.Println("Failed to subscribe to all topics")
        return
    }
    err = sub.SetRecvBuffer(1024 * 1204)
    if err != nil {
        fmt.Println("Failed to set recv buffer size")
        return
    }
    for id := 0; id < 10000; id++ {
        _, err := sub.Recv(0)
        t[id] = time.Now()
        if err != nil {
            fmt.Println("Sub1 : failed to receive")
            return
        }
    }
    f, _ := os.Create("./rxTime1")
    for id := 0; id < 10000; id++ {
        txt, _ := t[id].MarshalText()
        _, _ = f.Write(txt)
        _, _ = f.Write([]byte("\n"))
    }

}
