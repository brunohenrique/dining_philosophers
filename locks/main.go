package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type chopstick struct {
	sync.Mutex
}

type philosopher struct {
	name        string
	left, right int
}

func (t *philosopher) Think() {
	fmt.Printf("%s is thinking...\n", t.name)
	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	fmt.Printf("%s is hungry!\n", t.name)
}

func (t *philosopher) HoldChopsticks(cs []*chopstick) {
	cs[t.left].Lock()
	cs[t.right].Lock()
	fmt.Printf("%s hold chopsticks: %d and %d\n", t.name, t.left, t.right)
}

func (t *philosopher) Eat() {
	fmt.Printf("%s is eating using chopsticks: %d and %d\n", t.name, t.left, t.right)
	d := time.Duration(rand.Intn(10)) * time.Second
	time.Sleep(d)
}

func (t *philosopher) DropChopsticks(cs []*chopstick) {
	fmt.Printf("%s finished!\n", t.name)
	cs[t.left].Unlock()
	cs[t.right].Unlock()
	fmt.Printf("%s released chopsticks: %d and %d\n", t.name, t.left, t.right)
}

type table struct {
	philosophers []*philosopher
	chopsticks   []*chopstick
}

func NewTable(phils []*philosopher) table {
	t := table{}

	t.philosophers = phils
	for i := 0; i <= len(phils)-1; i++ {
		t.chopsticks = append(t.chopsticks, new(chopstick))
	}

	return t
}

func (t *table) Start() {
	for _, p := range t.philosophers {
		go func(p *philosopher, cs []*chopstick) {
			for {
				p.Think()
				p.HoldChopsticks(cs)
				p.Eat()
				p.DropChopsticks(cs)
			}
		}(p, t.chopsticks)
	}
}

func main() {
	philosophers := []*philosopher{
		&philosopher{name: "Plato", left: 0, right: 1},
		&philosopher{name: "Friedrich Nietzsche", left: 1, right: 2},
		&philosopher{name: "Alberto Camus", left: 2, right: 3},
		&philosopher{name: "Michel de Montaigne", left: 3, right: 4},
		&philosopher{name: "Jean-Paul Sartre", left: 4, right: 5},
		&philosopher{name: "Immanuel Kant", left: 5, right: 6},
		&philosopher{name: "René Descartes", left: 6, right: 7},
		&philosopher{name: "Arthur Schopenhauer", left: 7, right: 8},
		&philosopher{name: "Baruch Espinoza", left: 8, right: 0},
	}

	t := NewTable(philosophers)
	t.Start()

	fmt.Scanln()
}
