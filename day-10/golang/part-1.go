package main
import (
  "bufio"
  "fmt"
  "os"
  "strconv"
  "strings"
)

func main() {
  if len(os.Args) < 2 {
    return
  }
  infile, err := os.Open(os.Args[1])
  if err != nil { return }
  scanner := bufio.NewScanner(infile)
  scanner.Scan()
  line := strings.FieldsFunc(scanner.Text(), 
                             func (c rune) bool { 
                               return c == ',' 
                             })
  lengths := []int{}
  for _, i := range line {
    i, _ := strconv.Atoi(i)
    lengths = append(lengths, i)
  }

  position := 0
  skipsize := 0

  list := []int{}
  for i := 0; i < 256; i++ {
    list = append(list, i)
  }

  for _, length := range lengths {
    newlist := []int{}
    for _ = range list {
      newlist = append(newlist, 0)
    }
    ptr := position % len(list)
    for i := 0; i < length; i++ {
      newlist[(ptr + i) % len(list)] = list[(ptr + length-1-i + len(list)) % len(list)]
    }

    for i := 0; i < len(list)-length; i++ {
      newlist[(ptr+i+length) % len(list)] = list[(ptr+i+length) % len(list)]
    }

    list = newlist
    position = (position + skipsize + length) % len(list)
    skipsize++
  }

  fmt.Println(list[0]*list[1])
}

