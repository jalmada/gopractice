package main
 
 import(
     "fmt"
     "runtime"
     "time"
     "os"
     // "sync"
 )

 func main(){
     // mutex := new(sync.Mutex)
     runtime.GOMAXPROCS(4)


     start := time.Now()

     f, _ := os.Create("./log.txt")
     f.Close()

     logCh := make(chan string, 50)

     go func(){
         for{
             msg, ok := <- logCh
             if ok {
                 f, _ := os.OpenFile("./log.txt", os.O_APPEND, os.ModeAppend)

                 logTime := time.Now().Format(time.RFC3339)
                 f.WriteString(logTime + " - " + msg)
                 f.Close()
             } else {
                 break
             }
         }
     }()


     mutex := make(chan bool, 1)

     for i := 1; i < 10; i++{
         
         for j := 1; j < 10;  j++ {
             // mutex.Lock()
             mutex <- true
             go func( x int, y int){
                
                 msg := fmt.Sprintf("%d + %d = %d\n", x, y, x + y)
                 logCh <- msg 
                 fmt.Print(msg)
                 <- mutex
                 // mutex.Unlock()
             }(i, j)
         }
     }

     mutex <- true
     //mutex.Lock()
     elapsed := time.Since(start)    
     fmt.Printf("Excecution time %s", elapsed)
     fmt.Scanln()
 }