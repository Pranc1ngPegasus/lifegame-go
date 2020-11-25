package domain

import (
	"fmt"
	"math/rand"
	"time"
)

var _ LifegameInterface = (*lifegame)(nil)

const (
	width  = 120
	height = 40
	clear  = "\033[2J"
	head   = "\033[1;1H"
)

type (
	LifegameInterface interface {
		Initialize()
		Render()
		Update()
	}

	Cells [width + 2][height + 2]int

	lifegame struct {
		cells Cells
	}
)

func NewLifegame() LifegameInterface {
	return &lifegame{
		cells: Cells{},
	}
}

func (l *lifegame) Initialize() {
	fmt.Print(clear)
	rand.Seed(time.Now().UnixNano())

	for y := 0; y < height + 2; y++ {
		for x := 0; x < width + 2; x++ {
			if x == 0 || x == width + 2 {
				continue
			}
			if y == 0 || y == height + 2 {
				continue
			}

			l.cells[x][y] = rand.Intn(2)
		}
	}
}

func (l *lifegame) Render() {
	var screen string

	for y := 0; y < height + 2; y++ {
		for x := 0; x < width + 2; x++ {
			switch l.cells[x][y] {
			case 1:
				screen += "*"
			default:
				screen += " "
			}
		}
		screen += "\n"
	}

	fmt.Print(head)
	fmt.Print(screen)
}

func (l *lifegame) Update() {
	var nextCells Cells
	for y := 1; y < height + 1; y++ {
		for x := 1; x < width + 1; x++ {
			nextCells[x][y] = 0

			cnt := 0
			for _, yElm := range []int{-1, 0, 1} {
				for _, xElm := range []int{-1, 0, 1} {
					if xElm == 0 && yElm == 0 {
						continue
					}
					cnt += l.cells[x + xElm][y + yElm]
				}
			}

			switch l.cells[x][y] {
			case 0:
				switch cnt {
				case 3:
					nextCells[x][y] = 1
				default:
					nextCells[x][y] = 0
				}
			case 1:
				switch cnt {
				case 2, 3:
					nextCells[x][y] = 1
				default:
					nextCells[x][y] = 0
				}
			}
		}
	}
	l.cells = nextCells
}
