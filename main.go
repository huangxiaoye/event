package main
import (
	"fmt"
	. "event/event"
)

func main() {
	event := NewEvent()
	event.AddEvent("huanganxin", func(params ...interface{}) (next bool) {
			next = false
			fmt.Printf("%v", params)
			fmt.Println("I am")
			return
		})
	event.FireEvent("huanganxin", "NIHAO", "WOLAILE")
}