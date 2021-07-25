package main

import (
    "fmt"
    "net/http"
	"log"
	"os"
)
// create a logs file and prepare variables 
func init() {
    logs_file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
    if err != nil {
        log.Fatal(err)
    }

    InfoLogger = log.New(logs_file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
    WarningLogger = log.New(logs_file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
    ErrorLogger = log.New(logs_file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}
// endpoints functions 
func helloWorld(w http.ResponseWriter, r *http.Request){
    InfoLogger.Println("access main route")
    fmt.Fprintf(w, "Hello World")
}
func test_endpt(w http.ResponseWriter, r *http.Request){
    InfoLogger.Println("access test_endpt endpoint")
    fmt.Fprintf(w, "This is a new test endpoint ")
}
func logs_endpt(w http.ResponseWriter, r *http.Request){
    InfoLogger.Println("access logs endpoint")
	fmt.Fprintf(w, "This is a logs endpoint, please check the log file")



}
func main() {
    http.HandleFunc("/", helloWorld)
	http.HandleFunc("/test", test_endpt)
	http.HandleFunc("/logs", logs_endpt)
    http.ListenAndServe(":6111", nil)
}
