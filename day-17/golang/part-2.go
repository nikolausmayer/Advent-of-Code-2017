package main
import (
  "container/list"
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

  buffer := list.New()
  buffer.PushFront(0)
  pos := 0
  e := buffer.Front()

  for n := 1; n <= 50000000; n++ {
    if pos + num >= buffer.Len() {
      pos = (pos + num) % buffer.Len()
      e = buffer.Front()
      for i := 0; i < pos; i++ {
        e = e.Next()
      }
    } else {
      pos += num
      for i := 0; i < num; i++ {
        e = e.Next()
      }
    }
    e = buffer.InsertAfter(n, e)
    pos++
  }

  if buffer.Front().Value == 0 {
    e := buffer.Front().Next()
    fmt.Println(e.Value)
  } else {
    e := buffer.Front()
    for ; e.Value != 0; e = e.Next() { }
    fmt.Println(e.Next().Value)
  }
}

