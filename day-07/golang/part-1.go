package main
import (
  "bufio"
  "fmt"
  "regexp"
  "os"
)

func main() {
  if len(os.Args) < 2 {
    return
  }
  infile, err := os.Open(os.Args[1])
  if err != nil { return }
  defer infile.Close()
  scanner := bufio.NewScanner(infile)

  reOuter, _ := regexp.Compile(`^(\w+) \((\d+)\)(?: -> )?(.*)`)
  reInner, _ := regexp.Compile(`\w+`)

  allPrograms := make(map[string]struct{})
  topPrograms := make(map[string]struct{})

  for scanner.Scan() {
    line := scanner.Text()
    rOuter := reOuter.FindAllStringSubmatch(line, -1)
    allPrograms[rOuter[0][1]] = struct{}{}
    if len(rOuter[0][3]) > 0 {
      rInner := reInner.FindAllString(rOuter[0][3], -1)
      for _, b := range rInner {
        topPrograms[b] = struct{}{}
      }
    }
  }

  for p := range allPrograms {
    _, ok := topPrograms[p]
    if !ok {
      fmt.Println(p)
      break
    }
  }
}

