package main
import(
  "bufio"
  "fmt"
  "os"
  "regexp"
  "strconv"
)

func main() {
  if len(os.Args) < 2 {
    return
  }
  infile, err := os.Open(os.Args[1])
  if err != nil { return }
  scanner := bufio.NewScanner(infile)

  re, _ := regexp.Compile(`(\w+) (\w+) (-?\d+) if (\w+) ([<>=!]+) (-?\d+)`)

  registers := make(map[string]int)

  maxmax := -int(^uint(0) >> 1) - 1
  for scanner.Scan() {
    match  := re.FindAllStringSubmatch(scanner.Text(), -1)
    trgi   := match[0][1]
    trg,_  := registers[trgi]
    op     := match[0][2]
    arg1,_ := strconv.Atoi(match[0][3])
    src,_  := registers[match[0][4]]
    cond   := match[0][5]
    arg2,_ := strconv.Atoi(match[0][6])

    success := false
    switch cond {
      case "<": success = (src < arg2)
      case ">": success = (src > arg2)
      case "==": success = (src == arg2)
      case "!=": success = (src != arg2)
      case "<=": success = (src <= arg2)
      case ">=": success = (src >= arg2)
      default: fmt.Println(cond)
    }

    if success {
      switch op {
        case "inc": registers[trgi] = trg + arg1
        case "dec": registers[trgi] = trg - arg1
        default: fmt.Println(op)
      }
    }

    max := -int(^uint(0) >> 1) - 1
    for _, v := range registers {
      if v > max {
        max = v
      }
    }
    if max > maxmax {
      maxmax = max
    }
  }


  fmt.Println(maxmax)
}

