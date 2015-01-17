package terminus

import (
	"github.com/nsf/termbox-go"
	"strings"
)

type App struct {
	FgColor termbox.Attribute
	BgColor termbox.Attribute
}

func NewApp(fg, bg termbox.Attribute) *App {
	return &App{FgColor: fg, BgColor: bg}
}

func (a *App) Clear() {
	termbox.Clear(a.BgColor, a.BgColor)
}

func (a *App) DrawLines(lines string, x, y int) int {
	for _, line := range strings.Split(lines, "\n") {
		a.DrawLine(line, x, y)
		y++
	}

	return y
}

func (a *App) DrawLine(line string, x, y int) {
	for _, r := range []rune(line) {
		a.DrawRune(r, x, y)
		x++
	}
}

func (a *App) DrawRune(r rune, x, y int) {
	termbox.SetCell(x, y, r, a.FgColor, a.BgColor)
}
