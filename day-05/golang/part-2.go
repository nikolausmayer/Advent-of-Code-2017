package main
import (
  "bufio"
  "fmt"
  "os"
  "strconv"
)

func main() {
  if len(os.Args) < 2 {
    return
  }
  infile, err := os.Open(os.Args[1]);
  if err != nil {
    return
  }

  instructions := []int{}

  scanner := bufio.NewScanner(infile)
  for scanner.Scan() {
    raw := scanner.Text()
    num, err := strconv.Atoi(string(raw[0:len(raw)]))
    if err != nil {
      continue
    }
    instructions = append(instructions, num)
  }

  steps := 0
  iptr := 0
  for {
    if iptr < 0 || iptr >= len(instructions) {
      break
    }
    jmp := instructions[iptr]
    if jmp >= 3 {
      instructions[iptr]--
    } else {
      instructions[iptr]++
    }
    iptr += jmp
    steps++
  }
  fmt.Println(steps)
}

