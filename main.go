package main

import (
 "fmt"
 "os"
 "os/user"
 "itpr/repl"
)

func main() {
    user, err := user.Current()
    if err != nil {
        panic(err)
    }
    fmt.Printf("hello %s , this is the shit langauge \n", user.Name)
    fmt.Printf("feel free to type whatever \n")

    repl.Start(os.Stdin, os.Stdout)
}
