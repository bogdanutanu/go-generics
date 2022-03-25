package main

import (
	"fmt"

	"github.com/bradenaw/juniper/container/xlist"
)

func j() {
	var l xlist.List[string]

	l.PushFront("Rolling Stones")
	l.PushFront("Beatles")
	l.PushFront("Beach Boys")
	l.PushFront("Kinks")

	fmt.Printf("Lenght: %d\n", l.Len())
	fmt.Printf("Front: %s\n", l.Front().Value)
	fmt.Printf("Back: %s\n", l.Back().Value)

	// l.PushBack(5)
}
