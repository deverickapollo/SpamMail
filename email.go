package main

import (
    "bufio"
    "fmt"
    "log"
    "os"

)




func openFile(filename string) (*os.File ) {
	file, err := os.Open(filename)
    if err != nil {
        log.Fatal(err)
    }
    return file
}

func batchReadLines(file *os.File, currline int, chunk int){

}


func main() {

	file := openFile("1mill.txt")
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        fmt.Println(scanner.Text())
        //Send to a worker

    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    defer file.Close()

}
//Scanner does not deal well with lines longer than 65536 characters