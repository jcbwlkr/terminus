package terminus

import (
	"github.com/nsf/termbox-go"
)

func eventMeansCancel(ev termbox.Event) bool {
	return ev.Key == termbox.KeyEsc || ev.Key == termbox.KeyCtrlC || ev.Ch == 'q' || ev.Ch == 'Q'
}
