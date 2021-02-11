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

  grid := map[int]map[int]bool{}
  originalGrid := map[int]map[int]bool{}

  x, y := 0, 0
  scanner := bufio.NewScanner(infile)
  for scanner.Scan() {
    for _, c := range scanner.Text() {
      _, ok := grid[y]
      if !ok {
        grid[y] = map[int]bool{} 
        originalGrid[y] = map[int]bool{} 
      }
      grid[y][x] = (c == '#')
      originalGrid[y][x] = (c == '#')

      x++
    }
    y++
    x = 0
  }

  infectiousBursts := 0
  x = y/2
  y = x
  dx, dy := 0, -1
  for burst := 0; burst < 10000; burst++ {
    _, ok := grid[y]
    if !ok { grid[y] = map[int]bool{} }
    _, ok = grid[y][x]
    if ! ok { grid[y][x] = false }

    inOriginal := true
    _, ok = originalGrid[y]
    if ok { 
      v, ok := grid[y][x]
      if ok {
        inOriginal = v
      } else {
        inOriginal = false
      }
    } else {
      inOriginal = false
    }

    if grid[y][x] {
      dx, dy = -dy, dx
      grid[y][x] = false
    } else {
      dx, dy = dy, -dx
      grid[y][x] = true
      if !inOriginal {
        infectiousBursts++
      }
    }

    x, y = x+dx, y+dy
  }

  fmt.Println(infectiousBursts)
}

