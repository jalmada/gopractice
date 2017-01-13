package main

import (
    "time"
    "runtime"
)

func main(){

    godur, _ := time.ParseDuration("10ms")
    runtime.GOMAXPROCS(2)
    
    //go keyword enabled the concurrent behavior
    go func(){
        for x := 0; x < 100; x++ {
            println("Hello")
            time.Sleep(godur)
        }
    }()

    go func(){
        for x := 0; x < 100; x++ {        
            println("Go")
            time.Sleep(godur)
        }
    }()

    dur, _ := time.ParseDuration("1s")
    time.Sleep(dur)
}