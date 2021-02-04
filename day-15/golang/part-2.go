package main
import (
  "bufio"
  "fmt"
  "os"
  "strconv"
  "strings"
)

func main() {
  if len(os.Args) < 2 {
    return
  }
  infile, err := os.Open(os.Args[1])
  if err != nil { return }
  scanner := bufio.NewScanner(infile)
  generators := []int{}
  for scanner.Scan() {
    line := strings.Fields(scanner.Text())
    num, _ := strconv.Atoi(line[4])
    generators = append(generators, num)
  }

  factors := []int{16807, 48271}

  agreements := 0
  for i := 0; i < 5000000; i++ {
    for {
      generators[0] = (generators[0] * factors[0]) % 2147483647
      if generators[0] % 4 == 0 {
        break
      }
    }
    for {
      generators[1] = (generators[1] * factors[1]) % 2147483647
      if generators[1] % 8 == 0 {
        break
      }
    }
    if generators[0] % (1 << 16) == generators[1] % (1 << 16) {
      agreements++
    }
  }
  fmt.Println(agreements)
}

