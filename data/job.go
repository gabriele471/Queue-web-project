package data

import "fmt"

type Job struct {
	Title       string `json:"title" form:"title"`
	Description string `json:"description" form:"description"`
}

var Queue = []Job{}

func Enqueue(element Job) {
	Queue = append(Queue, element)
}

func PopQueue() {
	if len(Queue) > 0 {
		Queue = Queue[1:]
	} else {
		fmt.Println("No element to pop") //better to make a web message
	}

}
