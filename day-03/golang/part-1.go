package main
import (
  "fmt"
  "io/ioutil"
  "math"
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
  for i := 1; i <= target; i++ {
    x += dx
    y += dy
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
  
  fmt.Println(math.Abs(float64(x)) + math.Abs(float64(y)));
}

