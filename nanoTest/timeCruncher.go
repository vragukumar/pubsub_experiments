package main

import "os"
import "fmt"
import "time"
import "bufio"

func main() {
    var txTime [10000]time.Time
    var rxTime1 [10000]time.Time
    var rxTime2 [10000]time.Time
    var rx1delta, rx2delta int64
    var minrx1, minrx2 int64 = 1000000, 1000000
    var maxrx1, maxrx2 int64 = 0, 0
    var avgrx1, avgrx2 int64 = 0, 0

    tf, _ := os.Open("./txTime")
    rf1, _ := os.Open("./rxTime1")
    rf2, _ := os.Open("./rxTime2")
    txf := bufio.NewScanner(tf)
    rxf1 := bufio.NewScanner(rf1)
    rxf2 := bufio.NewScanner(rf2)

    id := 0
    for txf.Scan() {
        str := txf.Text()
        _ = (&txTime[id]).UnmarshalText([]byte(str))
        id += 1
    }
    id = 0
    for rxf1.Scan() {
        str := rxf1.Text()
        _ = (&rxTime1[id]).UnmarshalText([]byte(str))
        id += 1
    }
    id = 0
    for rxf2.Scan() {
        str := rxf2.Text()
        _ = (&rxTime2[id]).UnmarshalText([]byte(str))
        id += 1
    }
    for id = 0; id < 10000; id++ {
        rx1delta = int64(rxTime1[id].Sub(txTime[id]))
        rx2delta = int64(rxTime2[id].Sub(txTime[id]))
        if rx1delta < minrx1 {
            minrx1 = rx1delta
        }
        if rx2delta < minrx2 {
            minrx2 = rx2delta
        }
        if rx1delta > maxrx1 {
            maxrx1 = rx1delta
        }
        if rx2delta > maxrx2 {
            maxrx2 = rx2delta
        }
        avgrx1 += rx1delta
        avgrx2 += rx2delta
    }
    fmt.Println("Min, Max, Avg rx1 times = ", minrx1, maxrx1, avgrx1/10000)
    fmt.Println("Min, Max, Avg rx2 times = ", minrx2, maxrx2, avgrx2/10000)
}
