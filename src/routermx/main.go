package main
import "fmt"
import "time"
import "runtime"

func main() {
    var result int
    processors := runtime.GOMAXPROCS(0)-1
    for i := 0; i < processors; i++ {
        go func() {
            for { result++ }
        }()
    }
    time.Sleep(time.Second)       //wait for go function to increment the value.
    fmt.Println("result =", result)
}