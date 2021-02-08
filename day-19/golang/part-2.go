package main
import (
  "bufio"
  "fmt"
  "os"
)

func main() {
  if len(os.Args) < 2 { return }
  infile, err := os.Open(os.Args[1])
  if err != nil { return }

  network := []string{}

  scanner := bufio.NewScanner(infile)
  for scanner.Scan() {
    network = append(network, scanner.Text())
  }

  x, y := 0, 0
  for i, c := range network[0] {
    if c == '|' {
      x = i
      break
    }
  }

  steps := 0
  dx, dy := 0, 1
  for {
    c := network[y][x]
    switch c{
    case ' ':
      fmt.Println(steps)
      return
    case '+':
      if dy == 1 {
        if x < len(network[0])-1 && network[y][x+1] != ' ' {
          dx, dy = 1, 0
        } else {
          dx, dy = -1, 0
        }
      } else if dy == -1 {
        if x > 0 && network[y][x-1] != ' ' {
          dx, dy = -1, 0
        } else {
          dx, dy = 1, 0
        }
      } else if dx == 1 {
        if y > 0 && network[y-1][x] != ' ' {
          dx, dy = 0, -1
        } else {
          dx, dy = 0, 1
        }
      } else {
        if y < len(network)-1 && network[y+1][x] != ' ' {
          dx, dy = 0, 1
        } else {
          dx, dy = 0, -1
        }
      }
    case '-':
      fallthrough
    case '|':
      break
    }

    x += dx
    y += dy
    steps++
  }
}

