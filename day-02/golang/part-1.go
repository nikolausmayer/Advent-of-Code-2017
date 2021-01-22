package main
import (
  "fmt"
  "os"
  "math"
  "bufio"
  "strconv"
  "strings"
)

func main() {
  if len(os.Args) < 2 {
    return
  }
  infile, err := os.Open(os.Args[1])
  if err != nil {
    return
  }

  sum := 0
  scanner := bufio.NewScanner(infile)
  for scanner.Scan() {
    line := strings.Fields(scanner.Text())
    min, max := math.MaxInt32, 0
    for _, v := range line {
      v, _ := strconv.Atoi(v)
      if v < min { min = v }
      if v > max { max = v }
    }
    sum += max - min
  }

  fmt.Println(sum)
}

