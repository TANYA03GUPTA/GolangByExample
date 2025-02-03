package main
import (
    "fmt"
    "net/http"
    "os"
    "time"
)
func main() {
    client := &http.Client{
        Timeout: time.Nanosecond * 1,
    }
    _, err := client.Get("https://google.com")
    if os.IsTimeout(err) {
        fmt.Println("Timeout Happened")
    } else {
        fmt.Println("Timeout Did not Happened")
    }
}