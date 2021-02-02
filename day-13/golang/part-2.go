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

  for i := 0;; i++ {
    stop := true
    for d, p := range scanners {
      if (d+i) % (2*(p-1)) == 0 {
        stop = false
        break
      }
    }
    if stop {
      fmt.Println(i)
      break
    }
  }
}

