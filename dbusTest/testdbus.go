package main

import "os"
import "fmt"
import "time"
import "github.com/go.dbus"

func main() {
    var t [10000]time.Time
    conn, err := dbus.SessionBus()
    if err != nil {
        fmt.Println("Failed to connect to system bus", conn)
        return
    }
    reply, err := conn.RequestName("com.snaproute.asicd", dbus.NameFlagDoNotQueue)
    if reply != dbus.RequestNameReplyPrimaryOwner || err != nil {
        fmt.Println("Failed to acquire primary name - ", err)
        return
    }
    msg := "Test message - "
    for id := 0; id < 10000; id++ {
        sendMsg := msg + string(id)
        t[id] = time.Now()
        conn.Emit("/com/snaproute/asicd", "com.snaproute.asicd", []byte(sendMsg))
        time.Sleep(1 * time.Millisecond)
    }
    f, _ := os.Create("./txTime")
    for id := 0; id < 10000; id++ {
        txt, _ := t[id].MarshalText()
        _, _ = f.Write(txt)
        _, _ = f.Write([]byte("\n"))
    }
    fmt.Println("File write complete")
    for {}
}
