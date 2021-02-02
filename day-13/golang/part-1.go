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

  scanners := make(map[int]int)

  scanner := bufio.NewScanner(infile)
  for scanner.Scan() {
    line := strings.Fields(scanner.Text())
    depth, _  := strconv.Atoi(line[0][0:len(line[0])-1])
    length, _ := strconv.Atoi(line[1])
    scanners[depth] = length
  }

  severity := 0
  for d, p := range scanners {
    if d == 0 || d % (2*(p-1)) == 0 {
      severity += d*p
    }
  }

  fmt.Println(severity)
}

