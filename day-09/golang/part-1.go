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
    depth := 0
    score := 0
    for _, letter := range line {
      if escape {
        escape = false
        continue
      }
      switch letter {
        case '!': escape = true
        case '>': garbage = false
        case '<': garbage = true
        case '{': {
          if !garbage { 
            depth++ 
          }
        }
        case '}': {
          if !garbage {
            score += depth
            depth--
          }
        }
      }
    }
    fmt.Println(score)
  }
}

