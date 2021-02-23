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

func Test_checkEvent(t *testing.T) {

	events := []*Event{
		{
			Name:  "WEBRTC",
			Color: Blue,
		},
		{
			Name:   "SIGNAL",
			Color:  Yellow,
			Hidden: true,
		},
	}

	message := checkEvent(events, "[新] [WEBRTC] dddsdfsd")
	fmt.Print(message)
	fmt.Print("\n")

}
