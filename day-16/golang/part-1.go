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
  data, err := ioutil.ReadAll(infile)
  if err != nil { return }

  programs := []byte{}
  for i := 0; i < 16; i++ {
    programs = append(programs, byte(int('a') + i))
  }

  Spin := func(n int) {
    n = n % len(programs)
    x := []byte{}
    for i := len(programs)-n; i < len(programs); i++ {
      x = append(x, programs[i])
    }
    for i := 0; i < len(programs)-n; i++ {
      x = append(x, programs[i])
    }
    programs = x
  }

  Exchange := func(a, b int) {
    programs[a], programs[b] = programs[b], programs[a]
  }

  Partner := func(a, b byte) {
    for i := 0; i < len(programs); i++ {
      if programs[i] == a {
        programs[i] = b
      } else if programs[i] == b {
        programs[i] = a
      }
    }
  }

  mode := "x"
  num, swapNum := 0, 0
  numS, swapS := byte('x'), byte('x')
  for i := 0; i < len(data); i++ {
    c := data[i]
    switch c {
    case '\n':
      fallthrough
    case ',':
      switch mode {
      case "spin":
        Spin(num)
      case "exchange":
        Exchange(num, swapNum)
      case "partner":
        Partner(numS, swapS)
      }
      
      num, swapNum = 0, 0
      mode = "x"
    case '/':
      if mode == "partner" {
        swapS = numS
      } else {
        swapNum = num
        num = 0
      }
    default:
      if mode == "x" {
        if c == 's' {
          mode = "spin"
        } else if c == 'x' {
          mode = "exchange"
        } else if c == 'p' {
          mode = "partner"
        }
      } else {
        if mode == "partner" {
          numS = c
        } else {
          num = 10*num + (int(c) - int('0'))
        }
      }
    }
  }

  for _, c := range programs {
    fmt.Printf("%c", c)
  }
  fmt.Println()
}

