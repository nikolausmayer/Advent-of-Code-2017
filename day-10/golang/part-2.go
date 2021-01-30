package main
import (
  "fmt"
  "io/ioutil"
  "os"
)

func main() {
  if len(os.Args) < 2 {
    return
  }
  infile, err := os.Open(os.Args[1])
  if err != nil { return }
  raw, err := ioutil.ReadAll(infile)
  lengths := []int{}
  for _, i := range raw {
    if i == '\n' {
      continue
    }
    lengths = append(lengths, int(i))
  }
  lengths = append(lengths, 17)
  lengths = append(lengths, 31)
  lengths = append(lengths, 73)
  lengths = append(lengths, 47)
  lengths = append(lengths, 23)

  position := 0
  skipsize := 0

  list := []int{}
  for i := 0; i < 256; i++ {
    list = append(list, i)
  }

  for round := 0; round < 64; round++ {
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
  }

  idx := 0
  for idx < len(list) {
    xor := uint16(0)
    for i := 0; i < 16; i++ {
      xor = xor ^ uint16(list[idx])
      idx++
    }
    fmt.Printf("%02x", xor)
  }

  fmt.Println();
}

