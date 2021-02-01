package main
import (
  "bufio"
  "fmt"
  "os"
  "strings"
)

func main() {
  if len(os.Args) < 2 {
    return
  }
  infile, err := os.Open(os.Args[1])
  if err != nil { return }
  scanner := bufio.NewScanner(infile)

  connections := make(map[string][]string)

  for scanner.Scan() {
    line := strings.Fields(scanner.Text())
    a := line[0]
    for _, s := range line[2:len(line)-1] {
      connections[a] = append(connections[a], s[:len(s)-1])
    }
    connections[a] = append(connections[a], line[len(line)-1])
  }

  group := make(map[string]struct{})
  group["0"] = struct{}{}
  for {
    pre_length := len(group)
    for from, _ := range group {
      for _, to := range connections[from] {
        group[to] = struct{}{}
      }
    }
    if len(group) == pre_length {
      break
    }
  }

  fmt.Println(len(group))
}

