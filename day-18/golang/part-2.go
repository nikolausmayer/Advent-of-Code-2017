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

  registersA := make(map[string]int)
  registersA["p"] = 0
  iptrA := 0
  sendA := []int{}

  registersB := make(map[string]int)
  registersB["p"] = 1
  iptrB := 0
  sendB := []int{}

  sendBcount := 0

  for {
    if (iptrA < 0 || iptrA >= len(instructions)) &&
       (iptrB < 0 || iptrB >= len(instructions)) { 
      break 
    }

    pauseA, pauseB := false, false

    if iptrA >= 0 && iptrA < len(instructions) {
      line := strings.Fields(instructions[iptrA])
      switch line[0] {
      case "snd":
        i, err := strconv.Atoi(line[1])
        if err != nil {
          i = registersA[line[1]]
        }
        sendA = append(sendA, i)
        iptrA++
      case "set":
        i, err := strconv.Atoi(line[2])
        if err != nil {
          i = registersA[line[2]]
        }
        registersA[line[1]] = i
        iptrA++
      case "add":
        i, err := strconv.Atoi(line[2])
        if err != nil {
          i = registersA[line[2]]
        }
        registersA[line[1]] += i
        iptrA++
      case "mul":
        i, err := strconv.Atoi(line[2])
        if err != nil {
          i = registersA[line[2]]
        }
        registersA[line[1]] *= i
        iptrA++
      case "mod":
        i, err := strconv.Atoi(line[2])
        if err != nil {
          i = registersA[line[2]]
        }
        registersA[line[1]] %= i
        iptrA++
      case "rcv":
        if len(sendB) == 0 {
          pauseA = true
          break
        } else {
          registersA[line[1]] = sendB[0]
          sendB = append(sendB[:0], sendB[1:]...)
          }
          iptrA++
      case "jgz":
        i, err := strconv.Atoi(line[1])
        if err != nil {
          i = registersA[line[1]]
        }
        j, err := strconv.Atoi(line[2])
        if err != nil {
          j = registersA[line[2]]
        }
        if i > 0 {
          iptrA += j
        } else {
          iptrA++
        }
      }
    }

    if iptrB >= 0 && iptrB < len(instructions) {
      line := strings.Fields(instructions[iptrB])
      switch line[0] {
      case "snd":
        i, err := strconv.Atoi(line[1])
        if err != nil {
          i = registersB[line[1]]
        }
        sendB = append(sendB, i)
        iptrB++
        sendBcount++
      case "set":
        i, err := strconv.Atoi(line[2])
        if err != nil {
          i = registersB[line[2]]
        }
        registersB[line[1]] = i
        iptrB++
      case "add":
        i, err := strconv.Atoi(line[2])
        if err != nil {
          i = registersB[line[2]]
        }
        registersB[line[1]] += i
        iptrB++
      case "mul":
        i, err := strconv.Atoi(line[2])
        if err != nil {
          i = registersB[line[2]]
        }
        registersB[line[1]] *= i
        iptrB++
      case "mod":
        i, err := strconv.Atoi(line[2])
        if err != nil {
          i = registersB[line[2]]
        }
        registersB[line[1]] %= i
        iptrB++
      case "rcv":
        if len(sendA) == 0 {
          pauseB = true
          break
        } else {
          registersB[line[1]] = sendA[0]
          sendA = append(sendA[:0], sendA[1:]...)
        }
        iptrB++
      case "jgz":
        i, err := strconv.Atoi(line[1])
        if err != nil {
          i = registersB[line[1]]
        }
        j, err := strconv.Atoi(line[2])
        if err != nil {
          j = registersB[line[2]]
        }
        if i > 0 {
          iptrB += j
        } else {
          iptrB++
        }
      }
    }

    if pauseA && pauseB {
      break 
    }
  }

  fmt.Println(sendBcount)
}

