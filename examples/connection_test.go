package examples

import (
	"fmt"

	"github.com/alienchow/slackdk"
)

func ExamplePrintln() {
	c := slackdk.NewClient("mock api token")
	c.Connect()
	err := c.Send(&slackdk.Message{
		Type:    "message",
		Channel: "C0BQ28UAG",
		Text:    "testing 123",
	})
	fmt.Println(err)
	// Output: <nil>
}
