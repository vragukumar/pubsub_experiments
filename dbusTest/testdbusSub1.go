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
	conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0,
        "type='signal',path='/com/snaproute/asicd',interface='com.snaproute',sender='com.snaproute.asicd'")
    c := make(chan *dbus.Signal, 10000)
	conn.Signal(c)
    id := 0
    for _ = range c {
        t[id] = time.Now()
        id += 1
        if id == 10000 {
            break
        }
	}
    f, _ := os.Create("./rxTime1")
    for id := 0; id < 10000; id++ {
        txt, _ := t[id].MarshalText()
        _, _ = f.Write(txt)
        _, _ = f.Write([]byte("\n"))
    }
    fmt.Println("File write completed")
    for {}
}
