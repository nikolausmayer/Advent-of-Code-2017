package main
import (
  "fmt"
  "io/ioutil"
  "os"
  "strconv"
)

func main() {
  if len(os.Args) < 2 { return }
  infile, err := os.Open(os.Args[1])
  if err != nil { return }
  raw, err := ioutil.ReadAll(infile)
  if err != nil { return }
  num, err := strconv.Atoi(string(raw[:len(raw)-1]))
  if err != nil { return }

  buffer := []int{0}
  pos := 0

  for n := 1; n <= 2017; n++ {
    pos = (pos + num) % len(buffer)
    tmp := []int{}
    for i := 0; i <= pos; i++ {
      tmp = append(tmp, buffer[i])
    }
    tmp = append(tmp, n)
    for i := pos+1; i < len(buffer); i++ {
      tmp = append(tmp, buffer[i])
    }
    buffer = tmp
    pos++
  }

  if buffer[len(buffer)-1] == 2017 {
    fmt.Println(buffer[0])
  } else {
    for i := 0; i < len(buffer); i++ {
      if buffer[i] == 2017 {
        fmt.Println(buffer[i+1])
        break
      }
    }
  }
}

