package main
import(
  "bufio"
  "fmt"
  "os"
)

func main() {
  if len(os.Args) < 2 {
    return
  }
  infile, err := os.Open(os.Args[1])
  if err != nil { return }
  scanner := bufio.NewScanner(infile)

  for scanner.Scan() {
    line := scanner.Text();
    garbage := false
    escape := false
    score := 0
    for _, letter := range line {
      if escape {
        escape = false
        continue
      }
      switch letter {
        case '!': {
          escape = true
          continue
        }
        case '>': {
          garbage = false
          continue
        }
        case '<': {
          if !garbage {
            garbage = true
            continue
          }
        }
      }
      if garbage { 
        score++ 
      }
    }
    fmt.Println(score)
  }
}

