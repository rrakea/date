package main

import (
	"date/src"
	"fmt"
	"os"
)

func main(){
    if len(os.Args) == 0 {
        src.Remind()
    }    

    switch os.Args[0] {
    case "init":
        src.Init() 
    case "add":
        if len(os.Args) < 3 {
            fmt.Println("Not enough info for add command")
        }
        src.Add(os.Args[1], os.Args[2])
    case "remove":    
        if len(os.Args) < 3 {
            fmt.Println("Not enough info for remove command")
        }
        src.Add(os.Args[1], os.Args[2])
    case "list":
        if len(os.Args) < 2 {
            src.ListAll()
        } else {
            src.List(os.Args[1])
        }
    default:
        fmt.Println("Command not recognized")
        return
    }
}