package main

import (
	"fmt"
)

type Game struct {
	bestScore int
	scores    chan int
}

func (g *Game) run() {
	for score := range g.scores {
		if g.bestScore < score {
			g.bestScore = score
		}
	}
}

func NewGame() (g *Game) {
	g = &Game{
		bestScore: 0,
		scores:    make(chan int),
	}
	go g.run()
	return
}

type Player interface {
	NextScore() (score int, err error)
}

func (g *Game) HandlePlayer(p Player) error {
	for {
		score, err := p.NextScore()
		if err != nil {
			return err
		}
		g.scores <- score
	}
	// return err
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
	var arr = [5]interface{}{1, "2", 3.15, []int{21}}
	fmt.Println(arr)
	lol := NewGame()
	player := &User{}
	err := lol.HandlePlayer(player)
	if err != nil {
		fmt.Println(err)
	}
	for i := 0; i < 15; i++ {
		player.updateScore(i)
	}
	fmt.Println(player)        // &{14}
	fmt.Println(lol.bestScore) // 10
}
