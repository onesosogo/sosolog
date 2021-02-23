package sosolog

import (
	"fmt"
	"testing"
)

func Test_eventColorMessage(t *testing.T) {

	events := []*Event{
		{
			Name:  "新",
			Color: Blue,
		},
		{
			Name:  "SIGNAL",
			Color: Yellow,
		},
	}

	message := eventColorMessage(events, "[新] [SIGNALd] dddsdfsd")
	fmt.Print(message)
	fmt.Print("\n")
}
