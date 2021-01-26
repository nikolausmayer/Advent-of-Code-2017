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
  scanner.Scan()
  line := strings.Fields(scanner.Text())
  var banks []int
  for _, v := range line {
    i, err := strconv.Atoi(v)
    if err != nil { return }
    banks = append(banks, i)
  }
  visited := make(map[int]int)
  for {
    idx, maxe := 0, 0
    for i, e := range banks {
      if e > maxe {
        idx, maxe = i, e
      }
    }

    key := 0
    for _, e := range banks {
      key = key * (maxe + 1) + e
    }
    before, ok := visited[key]
    if ok {
      fmt.Println(len(visited) - before)
      break
    }
    visited[key] = len(visited)

    toDistribute := banks[idx]
    banks[idx] = 0
    idx++
    for ;toDistribute > 0; {
      idx = idx % len(banks)
      banks[idx]++
      idx++
      toDistribute--
    }
  }
}

