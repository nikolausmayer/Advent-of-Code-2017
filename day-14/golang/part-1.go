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
  if err != nil { return }

  originallengths := []int{}
  for _, i := range raw {
    if i == '\n' {
      continue
    }
    originallengths = append(originallengths, int(i))
  }

  squares := 0
  for row := 0; row < 128; row++ {

    lengths := []int{}
    for _, i := range originallengths {
      lengths = append(lengths, i)
    }
    lengths = append(lengths, int('-'))
    numstr := fmt.Sprintf("%d", row)
    for _, c := range numstr {
      lengths = append(lengths, int(c))
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
      hexdigits := fmt.Sprintf("%02x", xor)
      for i := 0; i < 2; i++ {
        switch hexdigits[i] {
        case '0': { squares += 0; } //fmt.Print("....") }
        case '1': { squares += 1; } //fmt.Print("...#") }
        case '2': { squares += 1; } //fmt.Print("..#.") }
        case '3': { squares += 2; } //fmt.Print("..##") }
        case '4': { squares += 1; } //fmt.Print(".#..") }
        case '5': { squares += 2; } //fmt.Print(".#.#") }
        case '6': { squares += 2; } //fmt.Print(".##.") }
        case '7': { squares += 3; } //fmt.Print(".###") }
        case '8': { squares += 1; } //fmt.Print("#...") }
        case '9': { squares += 2; } //fmt.Print("#..#") }
        case 'a': { squares += 2; } //fmt.Print("#.#.") }
        case 'b': { squares += 3; } //fmt.Print("#.##") }
        case 'c': { squares += 2; } //fmt.Print("##..") }
        case 'd': { squares += 3; } //fmt.Print("##.#") }
        case 'e': { squares += 3; } //fmt.Print("###.") }
        case 'f': { squares += 4; } //fmt.Print("####") }
        }
      }
    }

  }
  
  fmt.Println(squares)
}

