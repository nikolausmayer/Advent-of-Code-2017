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

  iptr := 0
  for {
    if iptr < 0 || iptr >= len(instructions) { break }
    line := strings.Fields(instructions[iptr])
    switch line[0] {
    case "snd":
      i, err := strconv.Atoi(line[1])
      if err != nil {
        i = registers[line[1]]
      }
      registers["sound"] = i
      //fmt.Println("beep", i)
      iptr++
    case "set":
      i, err := strconv.Atoi(line[2])
      if err != nil {
        i = registers[line[2]]
      }
      registers[line[1]] = i
      //fmt.Println("set", line[1], "to", i)
      iptr++
    case "add":
      i, err := strconv.Atoi(line[2])
      if err != nil {
        i = registers[line[2]]
      }
      registers[line[1]] += i
      //fmt.Println("add", line[1], "by", i)
      iptr++
    case "mul":
      i, err := strconv.Atoi(line[2])
      if err != nil {
        i = registers[line[2]]
      }
      registers[line[1]] *= i
      //fmt.Println("mul", line[1], "by", i)
      iptr++
    case "mod":
      i, err := strconv.Atoi(line[2])
      if err != nil {
        i = registers[line[2]]
      }
      registers[line[1]] %= i
      //fmt.Println("mod", line[1], "by", i)
      iptr++
    case "rcv":
      i, err := strconv.Atoi(line[1])
      if err != nil {
        i = registers[line[1]]
      }
      if i != 0 {
        if registers["sound"] != 0 {
          fmt.Println(registers["sound"])
          return
        }
      }
      iptr++
    case "jgz":
      i, err := strconv.Atoi(line[1])
      if err != nil {
        i = registers[line[1]]
      }
      j, err := strconv.Atoi(line[2])
      if err != nil {
        j = registers[line[2]]
      }
      if i > 0 {
        iptr += j
      } else {
        iptr++
      }
    default:
      //fmt.Println("unknown instruction", line)
      return
    }

    //fmt.Println(registers)
  }
}

