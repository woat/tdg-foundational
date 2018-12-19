package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	exec()
}

func exec() {
	g := newGame()
	g.board.populate()

	for {
		for _, e := range g.vboard {
			fmt.Printf("%v\n", e)
		}

		r := bufio.NewReader(os.Stdin)
		fmt.Printf("Coordinates (r,c): ")
		t, _ := r.ReadString('\n')

		if cheatCode(t[:len(t)-1]) {
			fmt.Printf("shrekt\n")
			for _, e := range g.board {
				fmt.Printf("%v\n", e)
			}
			fmt.Printf("shrekt\n")
			continue
		}

		tc := strings.Split(t[:len(t)-1], ",")

		coords := [2]int{}
		for i, c := range tc {
			cx, _ := strconv.Atoi(c)
			coords[i] = cx
		}

		g.hit(coords)

		if g.finished() {
			break
		}
	}

	for _, e := range g.vboard {
		fmt.Printf("%v\n", e)
	}

	if g.win {
		fmt.Printf("You win!\n")
	} else {
		fmt.Printf("GG\n")
	}
}

type board [10][10]string

func (b *board) populate() {
	rc := func(x int) (int, int) { return rand.Intn(x), rand.Intn(x) }
	for i := 0; i < 15; i++ {
		r, c := rc(len(*b))
		b[r][c] = strconv.Itoa(9)
		b.populateAdjacent(r, c)
	}
}

func (b *board) populateAdjacent(r, c int) {
	for ri := 0; ri < 3; ri++ {
		if r-1+ri < 0 || r-1+ri >= len(b) {
			continue
		}

		for ci := 0; ci < 3; ci++ {
			if c-1+ci < 0 || c-1+ci >= len(b) {
				continue
			}

			brc := &b[r-1+ri][c-1+ci]
			if *brc != "9" {
				n, _ := strconv.Atoi(*brc)
				*brc = strconv.Itoa(n + 1)
			}
		}
	}
}

type game struct {
	board  board
	vboard board
	fin    bool
	win    bool
}

func (g *game) hit(cs [2]int) {
	g.updatev(cs[0], cs[1])
	if g.board[cs[0]][cs[1]] == "9" {
		g.fin = true
		g.win = false
	}
}

func (g *game) updatev(r, c int) {
	v := &g.vboard
	uv(&g.board, v, r, c)
	v[r][c] = g.board[r][c]
}

func uv(b, v *board, r, c int) {
	if r < 0 || r >= len(v) {
		return
	}
	if c < 0 || c >= len(v) {
		return
	}
	vrc := &v[r][c]
	if *vrc != "-" {
		return
	}

	if (*b)[r][c] == "0" {
		*vrc = b[r][c]
		for ri := 0; ri < 3; ri++ {
			for ci := 0; ci < 3; ci++ {
				uv(b, v, r-1+ri, c-1+ci)
			}
		}
	}
}

func (g *game) finished() bool {
	if g.fin {
		return g.fin
	}

	for i := 0; i < len(g.vboard); i++ {
		for j := 0; j < len(g.vboard); j++ {
			if g.vboard[i][j] == "-" {
				return false
			}
		}
	}

	g.fin = true
	return g.fin
}

func newGame() *game {
	b := board{}
	v := board{}
	for i := 0; i < len(v); i++ {
		for j := 0; j < len(v); j++ {
			v[i][j] = "-"
			b[i][j] = "0"
		}
	}
	return &game{b, v, false, true}
}

func cheatCode(t string) bool {
	return t == "blacksheepwall"
}
