package main
import (
  "fmt"
  "os"
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
    line := []int{}
    for _, v := range strings.Fields(scanner.Text()) {
      a, _ := strconv.Atoi(v)
      for _, b := range line {
        if a < b && b%a == 0 {
          sum += b/a
          break
        }
        if b < a && a%b == 0 {
          sum += a/b
          break
        }
      }
      line = append(line, a)
    }
  }

  fmt.Println(sum)
}

