package game

import (
	termbox "github.com/nsf/termbox-go"
)

func Play() {
	termbox.Init()
	defer termbox.Close()

LOOP:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				break LOOP
			default:
				draw()
			}
		default:
			draw()
		}
	}
}

func draw() {
	coldef := termbox.ColorDefault
	termbox.SetCell(0, 0, '┏', coldef, coldef)
	termbox.SetCell(1, 0, '┓', coldef, coldef)
	termbox.SetCell(0, 1, '┗', coldef, coldef)
	termbox.SetCell(1, 1, '┛', coldef, coldef)
	termbox.Flush()
}
