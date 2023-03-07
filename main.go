package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	width  = 80
	height = 15
)

type Universe [][]bool

func NewUniverse() Universe {
	universe := make([][]bool, height)
	for i := range universe {
		universe[i] = make([]bool, width)
	}
	return universe
}
func (u Universe) Show() {

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if u[i][j] == false {
				fmt.Print(" ")
			} else {
				fmt.Print("*")
			}
		}
		fmt.Println(" ")
	}
}
func (u Universe) Seed() {
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			r := rand.Intn(3)
			if r == 1 {
				u[i][j] = true
			}
		}
	}

}
func (u Universe) Alive(x, y int) bool {

	x = (x + width) % width
	y = (y + height) % height
	return u[y][x]
}

func (u Universe) Neighbors(x, y int) int {

	cnt := 0
	for i := y - 1; i <= y+1; i++ {
		for j := x - 1; j <= x+1; j++ {

			if u.Alive(j, i) && !(i == y && j == x) {
				cnt++
			}
		}
	}
	return cnt
}
func (u Universe) Next(x, y int) bool {
	n := u.Neighbors(x, y)
	return n == 3 || n == 2 && u.Alive(x, y)
}
func Step(a, b Universe) {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			b[y][x] = a.Next(x, y)
		}
	}
}
func main() {
	a, b := NewUniverse(), NewUniverse()
	a.Seed()

	for i := 0; i < 300; i++ {
		Step(a, b)
		a.Show()
		//fmt.Print("\x0c")
		time.Sleep(time.Second / 30)
		a, b = b, a
	}
}
