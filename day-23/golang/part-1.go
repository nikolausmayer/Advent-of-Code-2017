package main
import (
  "bufio"
  "fmt"
  "os"
  "strconv"
  "strings"
)

func main() {
  if len(os.Args) < 2 { return }
  infile, err := os.Open(os.Args[1])
  if err != nil { return }
  scanner := bufio.NewScanner(infile)

  instructions := []string{}
  for scanner.Scan() {
    instructions = append(instructions, scanner.Text())
  }

  registers := make(map[string]int)

  debug := 0

  iptr := 0
  for {
    if iptr < 0 || iptr >= len(instructions) { break }
    line := strings.Fields(instructions[iptr])
    switch line[0] {
    case "set":
      i, err := strconv.Atoi(line[2])
      if err != nil {
        i = registers[line[2]]
      }
      registers[line[1]] = i
      iptr++
    case "sub":
      i, err := strconv.Atoi(line[2])
      if err != nil {
        i = registers[line[2]]
      }
      registers[line[1]] -= i
      iptr++
    case "mul":
      i, err := strconv.Atoi(line[2])
      if err != nil {
        i = registers[line[2]]
      }
      registers[line[1]] *= i
      debug++
      iptr++
    case "jnz":
      i, err := strconv.Atoi(line[1])
      if err != nil {
        i = registers[line[1]]
      }
      j, err := strconv.Atoi(line[2])
      if err != nil {
        j = registers[line[2]]
      }
      if i != 0 {
        iptr += j
      } else {
        iptr++
      }
    default:
      return
    }

  }
  fmt.Println(debug)
}

