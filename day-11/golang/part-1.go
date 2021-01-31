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
  line, err := ioutil.ReadAll(infile)
  if err != nil { return }

  // Y Z
  // |/
  // o
  //  \
  //   X
  x, y, z := 0, 0, 0
  var pp, p byte = ',', ','
  for _, c := range line {
    switch (c) {
      case '\n': 
        fallthrough
      case ',':
        switch (p) {
          case 'n':
            x--
            z++
          case 's':
            x++
            z--
          case 'e':
            if pp == 'n' {
              x--
              y++
            } else {
              y++
              z--
            }
          case 'w':
            if pp == 'n' {
              y--
              z++
            } else {
              x++
              y--
            }
        }
        fmt.Println(c,x,y,z)
      default:
        pp = p
        p  = c
    }
  }

  abs := func(i int) int {
    if i < 0 {
      return -i
    }
    return i
  }

  distance := (abs(x)+abs(y)+abs(z)+1)/2

  fmt.Println(distance)
}

