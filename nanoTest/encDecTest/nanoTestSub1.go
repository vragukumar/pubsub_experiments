package main

import "fmt"
import "bytes"
import "encoding/binary"
import "github.com/op/go-nanomsg"

type linkStatus struct {
    Port uint8
    LinkStatus uint8
}

func main() {
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
    var MsgType uint16
    for id := 0; id < 1; id++ {
        v, err := sub.Recv(0)
        if err != nil {
            fmt.Println("Sub1 : failed to receive")
            return
        }
        fmt.Println(v)
        buf := bytes.NewReader(v)
        err = binary.Read(buf, binary.LittleEndian, &MsgType)
        if err != nil {
            fmt.Println("Failed to decode msg")
            return
        }
        fmt.Println("Decoded msg type - ", MsgType)
        var msg linkStatus
        err = binary.Read(buf, binary.LittleEndian, &msg)
        if err != nil {
            fmt.Println("Failed to decode msg")
            return
        }
        fmt.Println("Decoded info for msg", msg)
    }
}
