package main

/*
 * Author: Deverick Simpson
 * go get gopkg.in/gomail.v2
 * Tested against https://github.com/rnwood/smtp4dev 
 * macOS Catalina 3.1 Quad-Core Intel Core I7
 * ulimit -n 10000 Increase open files.
 * 50000 messages 4minutesmidi
 * Scanner is not adequate for lines > 65536 chars
 * Goal: Read file and spam emails concurrently 
 */
 
import (
    "bufio"
    "fmt"
    "sync"
	"gopkg.in/gomail.v2"
    "log"
    "strings"
    "os"
    "time"
    "encoding/json"
)


var d = gomail.NewPlainDialer("localhost", 25, "login", "password")

/*
 *  Email data structure
 */
type Email struct{
	From string
	To string
	Subject string
	Text string

}

var msgstruct Email

/*
 *  Open a given file
 */
func openFile(filename string) (*os.File ) {
	file, err := os.Open(filename)
    if err != nil {
        log.Fatal(err)
    }
    return file
}

/*
 * Send an email  
 */
func sendEmail(msg string){
	err := json.Unmarshal([]byte(msg), &msgstruct)
    if err != nil {
        fmt.Println(err)
    }
    parseFromEmail := strings.Split(msgstruct.From,"@")
    parseToEmail := strings.Split(msgstruct.To,"@")
    m := gomail.NewMessage()
    m.SetAddressHeader("From", msgstruct.From, parseFromEmail[0])
    m.SetHeader("To", m.FormatAddress(msgstruct.To, parseToEmail[0]))
    m.SetHeader("Subject", msgstruct.Subject)
    m.SetBody("text/plain", msgstruct.Text)
    if err := d.DialAndSend(m); err != nil {
        panic(err)
    }
}

/*
 * Function for workers to call concurrently. Fan Out Philosophy. 
 */
func worker(msgChannel <-chan string, wg *sync.WaitGroup) {
    defer wg.Done()
    for {
        task, ok := <-msgChannel
        if !ok {
            return
        }
        sendEmail(task)
    }
}

func pool(wg *sync.WaitGroup, workers int) {

	file := openFile("1mill.txt")
    scanner := bufio.NewScanner(file)
    emailTask := make(chan string)

    for i := 0; i < workers; i++ {
        go worker(emailTask, wg)
    }

    for scanner.Scan() {
		emailTask <- scanner.Text() 
    }

    close(emailTask)
    if err := scanner.Err(); err != nil {
    	log.Fatal(err)
	}

    defer file.Close()
}

func main() {
	start := time.Now()
    var wg sync.WaitGroup
    wg.Add(50)
    go pool(&wg, 50)
    wg.Wait()
    elapsed := time.Since(start)
    log.Printf("Process took %s to complete", elapsed)
}
