package game

import (
	termbox "github.com/nsf/termbox-go"
)

var x int = 0
var y int = 0

func Play() {
	termbox.Init()
	defer termbox.Close()

	draw()
LOOP:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				break LOOP
			default:
				switch ev.Ch {
				case 'q':
					break LOOP
				case 'h':
					x -= 2
				case 'j':
					y += 1
				case 'k':
					y -= 1
				case 'l':
					x += 2
				default:
					draw()
				}
			}
		}
		draw()
	}
}

func draw() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	coldef := termbox.ColorBlue
	termbox.SetCell(x,     y, ' ', coldef, coldef)
	termbox.SetCell(x + 1, y, ' ', coldef, coldef)
	termbox.Flush()
}
