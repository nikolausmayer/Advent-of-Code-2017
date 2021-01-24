package main
import (
  "bufio"
  "fmt"
  "os"
  "sort"
  "strings"
)


type word []byte
func (w word) Less(i, j int) bool {
  return w[i] < w[j]
}
func (w word) Swap(i, j int) {
  w[i], w[j] = w[j], w[i]
}
func (w word) Len() int {
  return len(w)
}


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
    for _, w := range strings.Fields(scanner.Text()) {
      r := word(w)
      sort.Sort(r)
      s := string(r)
      if _, ok := dict[s]; ok {
        good = false
        break
      }
      dict[s] = 0
    }
    if good {
      validCount++
    }
  }

  fmt.Println(validCount);
}

