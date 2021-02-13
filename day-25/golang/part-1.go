package main
import (
  "bufio"
  "fmt"
  "os"
  "strconv"
  "strings"
)

type state struct {
  writeRules [2]int
  moveRules [2]int
  nextStateRules [2]string
}

func main() {
  if len(os.Args) < 2 { return }
  infile, err := os.Open(os.Args[1])
  if err != nil { return }

  scanner := bufio.NewScanner(infile)

  scanner.Scan()
  line := strings.Fields(scanner.Text())
  tmp := line[len(line)-1]
  startState := tmp[:len(tmp)-1]

  scanner.Scan()
  line = strings.Fields(scanner.Text())
  checksumAfter, _ := strconv.Atoi(line[5])

  states := map[string]state{}

  for scanner.Scan() {
    scanner.Scan()

    line = strings.Fields(scanner.Text())
    tmp = line[len(line)-1]
    name := tmp[:len(tmp)-1]

    scanner.Scan()

    scanner.Scan()
    line = strings.Fields(scanner.Text())
    tmp = line[len(line)-1]
    writeRuleZero, _ := strconv.Atoi(tmp[:len(tmp)-1])

    scanner.Scan()
    line = strings.Fields(scanner.Text())
    tmp = line[len(line)-1]
    moveRuleZero := -1
    if tmp == "right." { moveRuleZero = 1 }

    scanner.Scan()
    line = strings.Fields(scanner.Text())
    tmp = line[len(line)-1]
    nextStateRuleZero := tmp[:len(tmp)-1]
    
    scanner.Scan()

    scanner.Scan()
    line = strings.Fields(scanner.Text())
    tmp = line[len(line)-1]
    writeRuleOne, _ := strconv.Atoi(tmp[:len(tmp)-1])

    scanner.Scan()
    line = strings.Fields(scanner.Text())
    tmp = line[len(line)-1]
    moveRuleOne := -1
    if tmp == "right." { moveRuleOne = 1 }

    scanner.Scan()
    line = strings.Fields(scanner.Text())
    tmp = line[len(line)-1]
    nextStateRuleOne := tmp[:len(tmp)-1]


    states[name] = state{
      [2]int{writeRuleZero, writeRuleOne},
      [2]int{moveRuleZero, moveRuleOne},
      [2]string{nextStateRuleZero, nextStateRuleOne},
    }
  }

  tape := map[int]int{}
  cursor := 0
  state := states[startState]

  for ; checksumAfter > 0; checksumAfter-- {
    read, ok := tape[cursor]
    if !ok { read = 0 }
    tape[cursor] = state.writeRules[read]
    cursor += state.moveRules[read]
    state = states[state.nextStateRules[read]]
  }

  checksum := 0
  for _, v := range tape {
    if v == 1 { checksum++ }
  }

  fmt.Println(checksum)
}

