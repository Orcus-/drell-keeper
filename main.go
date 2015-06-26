/*
   simple application to manage and create passwords with support
   for public and private keys
*/
package main

import (
	"log"
	"os"
)


// leveled logging
var (
    LogErr  *log.Logger
    LogMsg  *log.Logger
    )
    
// start logging (its public so testing func can use it)
func InitLogger() {
    // open logging file
    file, err := os.OpenFile("log.txt", os.O_RDWR | os.O_APPEND | os.O_CREATE, 0666)
    // clean up
    defer file.Close()
    if err != nil {
        panic(err)
    }
    LogErr = log.New(file, "ERROR: ", log.LstdFlags | Lshortfile)
    LogMsg = log.New(file, "MSG: ", log.LstdFlags)
    