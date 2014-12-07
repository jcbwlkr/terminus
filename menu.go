package terminus

import (
	"github.com/nsf/termbox-go"
	"strings"
)

type Menu struct {
	Title    string
	Options  []MenuOption
	selected int
}

type MenuOption struct {
	Label  string
	Action func() int
}

const (
	MenuContinue = iota
	MenuExit
)

// Run is the main function that draws and handles a Menu. It is a blocking
// operation.
func (m *Menu) Run() {
	m.draw()
loop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			// Break on Escape, Ctrl-c, q, or Q
			// TODO what's a more idiomatic way of doing this?
			if ev.Key == termbox.KeyEsc || ev.Key == termbox.KeyCtrlC || ev.Ch == 'q' || ev.Ch == 'Q' {
				break loop
			}
			// TODO support vim and readline menu navigation
			if ev.Key == termbox.KeyArrowDown {
				if m.selected < len(m.Options)-1 {
					m.selected++
				}
			} else if ev.Key == termbox.KeyArrowUp {
				if m.selected > 0 {
					m.selected--
				}
			} else if ev.Key == termbox.KeyEnter {
				if result := m.Options[m.selected].Action(); result == MenuExit {
					break loop
				}
			}
			m.draw()
		case termbox.EventResize:
			m.draw()
		case termbox.EventError:
			panic(ev.Err)
		}
	}
}

func (m *Menu) draw() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	w, _ := termbox.Size()

	// Draw our title string in the middle of the screen. Figure out how wide
	// it is at it's widest point and calculate x from that
	x := w/2 - maxLineLength(m.Title)/2
	y := 0

	y = draw_lines(m.Title, x, y)

	var pre string
	for i, opt := range m.Options {
		if i == m.selected {
			pre = "->"
		} else {
			pre = " -"
		}

		draw_line(pre+" "+opt.Label, x, y)
		y++
	}

	termbox.Flush()
}

func draw_lines(lines string, x, y int) int {
	for _, line := range strings.Split(lines, "\n") {
		draw_line(line, x, y)
		y++
	}

	return y
}

func draw_line(line string, x, y int) {
	for _, r := range []rune(line) {
		termbox.SetCell(x, y, r, termbox.ColorWhite, termbox.ColorDefault)
		x++
	}
}
