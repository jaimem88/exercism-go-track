package main

import (
	"exercism/go/bob"
	"exercism/go/gigasecond"
	greeting "exercism/go/hello-world"
	linkedlist "exercism/go/linked-list"
	"exercism/go/space-age"
	twofer "exercism/go/two-fer"
	"fmt"
	"time"
)

func main() {
	// hello world
	fmt.Println(greeting.HelloWorld())
	// two fer
	fmt.Println(twofer.ShareWith("Bob"))
	// space age
	fmt.Printf("Age in Jupiter: %f earth years\n", space.Age(1000000000, space.Jupiter))
	// gigasecond
	fmt.Printf("now+gigasecond: %s\n ", gigasecond.AddGigasecond(time.Now()))

	// bob
	fmt.Println(bob.Hey("SCREAMING!"))

	linkedList()

}

func linkedList() {
	l := linkedlist.NewList([]interface{}{1, 2, 3, 4, 5}...)
	fmt.Println("LinkedList: before...")
	l.Traverse()
	l.Reverse()
	fmt.Println("LinkedList: after...")
	l.Traverse()
}
