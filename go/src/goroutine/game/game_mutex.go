package main

import (
	"fmt"
	"sync"
)

type Game struct {
	mtx       sync.Mutex
	bestScore int
}

func NewGame() *Game {
	return &Game{}
}

type Player interface {
	NextScore() (score int, err error)
}

func (g *Game) HandlePlayer(p Player) error {
	// for {
	score, err := p.NextScore()
	if err != nil {
		return err
	}
	g.mtx.Lock()
	if g.bestScore < score {
		g.bestScore = score
	}
	g.mtx.Unlock()
	return err
	// }
}

type User struct {
	score int
}

func (u *User) updateScore(score int) {
	u.score = score
}
func (u *User) NextScore() (score int, err error) {
	if u.score > 10 {
		return u.score, fmt.Errorf("已到最高分")
	}
	return u.score, nil
}
func main() {
	lol := NewGame()
	player := &User{}

	for i := 0; i < 15; i++ {
		player.updateScore(i)
		err := lol.HandlePlayer(player)
		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println(player)        // &{14}
	fmt.Println(lol.bestScore) // 10
}
