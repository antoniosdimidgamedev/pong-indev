package main

import (
    "fmt"
    "time"
    "os"
    "math/rand"
    "github.com/nsf/termbox-go"
)

const (
    WIDTH  = 50
    HEIGHT = 20
)

type Ball struct {
    x, y    int
    dx, dy  int
}

var ( 
    ball Ball
    score int
)

func draw() {
    termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
    for i := 0; i < HEIGHT; i++ {
        termbox.SetCell(0, i, ' ', termbox.ColorDefault, termbox.ColorDefault)
        termbox.SetCell(WIDTH-1, i, ' ', termbox.ColorDefault, termbox.ColorDefault)
    }
    termbox.SetCell(ball.x, ball.y, 'O', termbox.ColorDefault, termbox.ColorDefault)
    termbox.Flush()
}

func update() {
    ball.x += ball.dx
    ball.y += ball.dy
    if ball.y <= 0 || ball.y >= HEIGHT-1 {
        ball.dy = -ball.dy
    }

    if ball.x <= 0 || ball.x >= WIDTH-1 {
        ball.dx = -ball.dx
        score++
        fmt.Printf("Score: %d\n", score)
    }
}

func input() {
    for {
        switch ev := termbox.PollEvent(); ev.Type {
        case termbox.EventKey:
            if ev.Key == termbox.KeyCtrlC {
                os.Exit(0)
            }
        }
    }
}
}

func main() {
    rand.Seed(time.Now().UnixNano())
    ball.x = WIDTH / 2
    ball.y = HEIGHT / 2
    ball.dx = 1
    ball.dy = 1

    if err := termbox.Init(); err != nil {
        panic(err)
    }
    defer termbox.Close()

    go input()
    for {
        draw()
        update()
        time.Sleep(100 * time.Millisecond)
    }
}