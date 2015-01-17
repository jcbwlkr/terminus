package main

import (
	"github.com/jcbwlkr/terminus"
	"github.com/nsf/termbox-go"
	"time"
)

func main() {
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	defer termbox.Close()

	app := terminus.NewApp(termbox.ColorWhite, termbox.ColorBlack)

	menu := terminus.NewMenu(app)

	menu.Title = `
___
 ||_  _
 || |(/_
    _                                      _
   / \   _ __ _ __ _____      __      __ _| |_ ___  _ __
  / _ \ | '__| '__/ _ \ \ /\ / /____ / _' | __/ _ \| '__|
 / ___ \| |  | | | (_) \ V  V /_____| (_| | || (_) | |
/_/   \_\_|  |_|  \___/ \_/\_/       \__,_|\__\___/|_|
`
	menu.AddOption(&terminus.MenuOption{"Shoot Right", shoot_right})
	menu.AddOption(&terminus.MenuOption{"Shoot Left", shoot_left})
	menu.AddOption(terminus.NewExitOption("Quit"))

	menu.Run()
}

func shoot_right(app *terminus.App) int {
	app.Clear()
	w, h := termbox.Size()

	y := h / 2
	for x := 0; x <= w; x++ {
		app.DrawRune('→', x, y)
		if x > 0 {
			app.DrawRune(' ', x-1, y)
		}
		termbox.Flush()
		time.Sleep(10 * time.Millisecond)
	}

	return terminus.Continue
}

func shoot_left(app *terminus.App) int {
	app.Clear()
	w, h := termbox.Size()

	y := h / 2
	for x := w; x >= 0; x-- {
		app.DrawRune('←', x, y)
		if x < w {
			app.DrawRune(' ', x+1, y)
		}

		termbox.Flush()
		time.Sleep(10 * time.Millisecond)
	}

	return terminus.Continue
}
