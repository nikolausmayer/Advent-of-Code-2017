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

  grid := [128][128]bool{}

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
        case '0': 
          grid[row][i*4+0+idx/2-8] = false
          grid[row][i*4+1+idx/2-8] = false
          grid[row][i*4+2+idx/2-8] = false
          grid[row][i*4+3+idx/2-8] = false
        case '1': 
          grid[row][i*4+0+idx/2-8] = false
          grid[row][i*4+1+idx/2-8] = false
          grid[row][i*4+2+idx/2-8] = false
          grid[row][i*4+3+idx/2-8] = true
        case '2': 
          grid[row][i*4+0+idx/2-8] = false
          grid[row][i*4+1+idx/2-8] = false
          grid[row][i*4+2+idx/2-8] = true
          grid[row][i*4+3+idx/2-8] = false
        case '3': 
          grid[row][i*4+0+idx/2-8] = false
          grid[row][i*4+1+idx/2-8] = false
          grid[row][i*4+2+idx/2-8] = true
          grid[row][i*4+3+idx/2-8] = true
        case '4': 
          grid[row][i*4+0+idx/2-8] = false
          grid[row][i*4+1+idx/2-8] = true
          grid[row][i*4+2+idx/2-8] = false
          grid[row][i*4+3+idx/2-8] = false
        case '5': 
          grid[row][i*4+0+idx/2-8] = false
          grid[row][i*4+1+idx/2-8] = true
          grid[row][i*4+2+idx/2-8] = false
          grid[row][i*4+3+idx/2-8] = true
        case '6': 
          grid[row][i*4+0+idx/2-8] = false
          grid[row][i*4+1+idx/2-8] = true
          grid[row][i*4+2+idx/2-8] = true
          grid[row][i*4+3+idx/2-8] = false
        case '7': 
          grid[row][i*4+0+idx/2-8] = false
          grid[row][i*4+1+idx/2-8] = true
          grid[row][i*4+2+idx/2-8] = true
          grid[row][i*4+3+idx/2-8] = true
        case '8': 
          grid[row][i*4+0+idx/2-8] = true
          grid[row][i*4+1+idx/2-8] = false
          grid[row][i*4+2+idx/2-8] = false
          grid[row][i*4+3+idx/2-8] = false
        case '9': 
          grid[row][i*4+0+idx/2-8] = true
          grid[row][i*4+1+idx/2-8] = false
          grid[row][i*4+2+idx/2-8] = false
          grid[row][i*4+3+idx/2-8] = true
        case 'a': 
          grid[row][i*4+0+idx/2-8] = true
          grid[row][i*4+1+idx/2-8] = false
          grid[row][i*4+2+idx/2-8] = true
          grid[row][i*4+3+idx/2-8] = false
        case 'b': 
          grid[row][i*4+0+idx/2-8] = true
          grid[row][i*4+1+idx/2-8] = false
          grid[row][i*4+2+idx/2-8] = true
          grid[row][i*4+3+idx/2-8] = true
        case 'c': 
          grid[row][i*4+0+idx/2-8] = true
          grid[row][i*4+1+idx/2-8] = true
          grid[row][i*4+2+idx/2-8] = false
          grid[row][i*4+3+idx/2-8] = false
        case 'd': 
          grid[row][i*4+0+idx/2-8] = true
          grid[row][i*4+1+idx/2-8] = true
          grid[row][i*4+2+idx/2-8] = false
          grid[row][i*4+3+idx/2-8] = true
        case 'e': 
          grid[row][i*4+0+idx/2-8] = true
          grid[row][i*4+1+idx/2-8] = true
          grid[row][i*4+2+idx/2-8] = true
          grid[row][i*4+3+idx/2-8] = false
        case 'f': 
          grid[row][i*4+0+idx/2-8] = true
          grid[row][i*4+1+idx/2-8] = true
          grid[row][i*4+2+idx/2-8] = true
          grid[row][i*4+3+idx/2-8] = true
        }
      }
    }

  }


  var deleter func(x int, y int)
  deleter = func(x int, y int) {
    if x < 0 || x >= 128 || 
       y < 0 || y >= 128 {
      return
    }
    grid[y][x] = false;
    if x > 0 && grid[y][x-1] {
      deleter(x-1, y)
    }
    if x < 127 && grid[y][x+1] {
      deleter(x+1, y)
    }
    if y > 0 && grid[y-1][x] {
      deleter(x, y-1)
    }
    if y < 127 && grid[y+1][x] {
      deleter(x, y+1)
    }
  }

  regions := 0
  for y := 0; y < 128; y++ {
    for x := 0; x < 128; x++ {
      if grid[y][x] {
        deleter(x, y)
        regions++
      }
    }
  }

  fmt.Println(regions)
}

