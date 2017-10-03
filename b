package main

import (
    "fmt"
    "os"
)

func main() {
    if len(os.Args) < 3 {
        fmt.Printf("Wrong number of args")
        os.Exit(2)
    }
    
    os.Exit(0)
}