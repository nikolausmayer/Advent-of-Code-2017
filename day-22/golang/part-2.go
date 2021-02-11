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

  grid := map[int]map[int]int{}
  originalGrid := map[int]map[int]int{}

  x, y := 0, 0
  scanner := bufio.NewScanner(infile)
  for scanner.Scan() {
    for _, c := range scanner.Text() {
      _, ok := grid[y]
      if !ok {
        grid[y] = map[int]int{} 
        originalGrid[y] = map[int]int{} 
      }
      if c == '#' {
        grid[y][x] = 2
        originalGrid[y][x] = 2
      } else {
        grid[y][x] = 0
        originalGrid[y][x] = 0
      }

      x++
    }
    y++
    x = 0
  }

  infectiousBursts := 0
  x = y/2
  y = x
  dx, dy := 0, -1
  for burst := 0; burst < 10000000; burst++ {
    _, ok := grid[y]
    if !ok { grid[y] = map[int]int{} }
    _, ok = grid[y][x]
    if ! ok { grid[y][x] = 0 }

    switch grid[y][x] {
    case 0:
      grid[y][x] = 1
      dx, dy = dy, -dx
    case 1:
      grid[y][x] = 2
      infectiousBursts++
    case 2:
      grid[y][x] = 3
      dx, dy = -dy, dx
    case 3:
      grid[y][x] = 0
      dx, dy = -dx, -dy
    }

    x, y = x+dx, y+dy
  }

  fmt.Println(infectiousBursts)
}

