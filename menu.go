package terminus

import (
	"github.com/nsf/termbox-go"
)

const (
	Continue = iota
	Exit
)

type Menu struct {
	Title    string
	app      *App
	options  []*MenuOption
	selected int
}

func NewMenu(a *App) *Menu {
	return &Menu{app: a, options: make([]*MenuOption, 0)}
}

func (m *Menu) AddOption(o *MenuOption) {
	m.options = append(m.options, o)
}

type MenuOption struct {
	Label  string
	Action func(*App) int
}

func NewExitOption(label string) *MenuOption {
	return &MenuOption{label, func(a *App) int { return Exit }}
}

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
			if eventMeansCancel(ev) {
				break loop
			}
			// TODO support vim and readline style menu navigation
			if ev.Key == termbox.KeyArrowDown {
				if m.selected < len(m.options)-1 {
					m.selected++
				}
			} else if ev.Key == termbox.KeyArrowUp {
				if m.selected > 0 {
					m.selected--
				}
			} else if ev.Key == termbox.KeyEnter {
				if result := m.options[m.selected].Action(m.app); result == Exit {
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
	m.app.Clear()
	w, _ := termbox.Size()

	// Draw our title string in the middle of the screen. Figure out how wide
	// it is at it's widest point and calculate x from that
	x := w/2 - maxLineLength(m.Title)/2
	y := 0

	y = m.app.DrawLines(m.Title, x, y)

	var pre string
	for i, opt := range m.options {
		if i == m.selected {
			pre = "→"
		} else {
			pre = " "
		}

		m.app.DrawLine(pre+" "+opt.Label, x, y)
		y++
	}

	termbox.Flush()
}
