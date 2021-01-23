package main
import (
  "fmt"
  "io/ioutil"
  "os"
  "strconv"
)


func main() {
  if len(os.Args) < 2 {
    return
  }
  infile, err := os.Open(os.Args[1])
  if err != nil {
    return
  }

  data, err := ioutil.ReadAll(infile)
  if err != nil {
    return
  }
  target, err := strconv.Atoi(string(data[0:len(data)-1]))
  if err != nil {
    return
  }

  x, y, dx, dy := -1, 0, 1, 0
  field := make(map[[2]int]int)
  for i := 1; i < target; i++ {
    x += dx
    y += dy

    sum := 0
    for ty := -1; ty <= 1; ty++ {
      for tx := -1; tx <= 1; tx++ {
        v, ok := field[[2]int{x+tx, y+ty}]
        if ok {
          sum += v
        }
      }
    }
    if i == 1 {
      sum = 1
    }
    field[[2]int{x, y}] = sum
    if sum > target {
      fmt.Println(sum)
      break
    }

    if dx == 1 && y == x-1 {
      // [>] --> [^]
      dx, dy = 0, -1
    } else if dy == -1 && x == -y {
      // [^] --> [<]
      dx, dy = -1, 0
    } else if dx == -1 && x == y {
      // [<] --> [v]
      dx, dy = 0, 1
    } else if dy == 1 && x == -y {
      // [v] --> [>]
      dx, dy = 1, 0
    }
  }
}

