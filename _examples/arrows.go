package main

import (
	"github.com/jacobwalker0814/terminus-go"
	"github.com/nsf/termbox-go"
	"time"
)

func main() {
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	defer termbox.Close()

	menu := new(terminus.Menu)

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
	menu.Options = make([]terminus.MenuOption, 3)
	menu.Options[0] = terminus.MenuOption{"Shoot Right", shoot_right}
	menu.Options[1] = terminus.MenuOption{"Shoot Left", shoot_left}
	menu.Options[2] = terminus.NewExitOption("Quit")

	menu.Run()
}

func shoot_right() int {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	w, h := termbox.Size()

	y := h / 2
	for x := 0; x <= w; x++ {
		termbox.SetCell(x, y, '>', termbox.ColorWhite, termbox.ColorDefault)
		termbox.Flush()
		time.Sleep(6 * time.Millisecond)
	}

	return terminus.Continue
}

func shoot_left() int {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	w, h := termbox.Size()

	y := h / 2
	for x := w; x >= 0; x-- {
		termbox.SetCell(x, y, '<', termbox.ColorWhite, termbox.ColorDefault)
		termbox.Flush()
		time.Sleep(6 * time.Millisecond)
	}

	return terminus.Continue
}
