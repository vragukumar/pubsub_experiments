package main

import "os"
import "fmt"
import "time"
import "github.com/op/go-nanomsg"

func main() {
    var t [10000]time.Time
    pub, err := nanomsg.NewPubSocket()
    if err != nil {
        fmt.Println("Failed to open pub socket")
        return
    }
    ep, err := pub.Bind("ipc:///tmp/test.ipc")
    if err != nil {
        fmt.Println("Failed to bind pub socket - ", ep)
        return
    }
    err = pub.SetSendBuffer(1024*1024)
    if err != nil {
        fmt.Println("Failed to set send buffer size")
        return
    }
    time.Sleep(100 * time.Millisecond)
    msg := "Test message - "
    for id := 0; id < 10000; id++ {
        sendMsg := msg + string(id)
        buf := []byte(sendMsg)
        t[id] = time.Now()
        pub.Send(buf, 0)
        time.Sleep(1 * time.Millisecond)
    }
    f, _ := os.Create("./txTime")
    for id := 0; id < 10000; id++ {
        txt, _ := t[id].MarshalText()
        _, _ = f.Write(txt)
        _, _ = f.Write([]byte("\n"))
    }
}
