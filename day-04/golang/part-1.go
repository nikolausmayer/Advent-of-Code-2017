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
  if err != nil {
    return
  }
  defer infile.Close()

  validCount := 0
  scanner := bufio.NewScanner(infile)
  for scanner.Scan() {
    dict := make(map[string]int)
    good := true
    for _, word := range strings.Fields(scanner.Text()) {
      if _, ok := dict[word]; ok {
        good = false
        break
      }
      dict[word] = 0
    }
    if good {
      validCount++
    }
  }


  fmt.Println(validCount);
}

