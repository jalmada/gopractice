package main

import (
    "fmt"
    "strings"
)

type Message struct {
    To []string
    From string
    Content string
}

type FailedMessage struct {
    ErrorMessage string
    OriginalMessage Message
}

func main(){
    SwitchChannel()
    //Channels()
}

func SwitchChannel(){
    msgCh := make(chan Message, 1)
    errCh := make(chan FailedMessage, 1)

    msg := Message {
        To: []string{"frodo@underhill.me"},
        From: "gandalf@whitecouncil.org",
        Content: "Keep it secret, keep it safe.",
    }

    failedMessage := FailedMessage {
        ErrorMessage: "Message Intercepted by black rider",
        OriginalMessage: Message{},
    }

    msgCh <- msg
    errCh <- failedMessage

    select {
        case receivedMsg := <- msgCh:
            fmt.Println(receivedMsg)
        case receivedError := <- errCh:
            fmt.Println(receivedError) 
        default:
            fmt.Println("No Message Received")
    }

    
}

func Channels(){
    phrase := "These are the times tat try men's soulds. \n"
    words := strings.Split(phrase, " ")

    ch := make(chan string, len(words))

    for _, word := range words {
        ch <- word
    }

    close(ch)

    for msg := range ch {
            fmt.Println(msg + " ")
    }
}