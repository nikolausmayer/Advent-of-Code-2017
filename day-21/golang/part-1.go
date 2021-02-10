package main
import (
  "bufio"
  "fmt"
  "os"
  "strings"
)

func flip(s string) string {
  div := 2
  if len(s) != 5 {
    div = 3
  }
  result := []byte{}
  for y := 0; y < div; y++ {
    for x := 0; x < div; x++ {
      result = append(result, s[y*(div+1)+(div-1-x)])
    }
    if y < div-1 {
      result = append(result, '/')
    }
  }
  return string(result)
}

func rot(s string) string {
  div := 2
  if len(s) != 5 {
    div = 3
  }
  result := []byte{}
  for y := 0; y < div; y++ {
    for x := 0; x < div; x++ {
      result = append(result, s[(div-1-x)*(div+1)+y])
    }
    if y < div-1 {
      result = append(result, '/')
    }
  }
  return string(result)
}

func main() {
  if len(os.Args) < 2 { return }
  infile, err := os.Open(os.Args[1])
  if err != nil { return }

  rules := [][]string{}

  scanner := bufio.NewScanner(infile)
  for scanner.Scan() {
    line := strings.Fields(scanner.Text())
    from, to := line[0], line[2]
    rules = append(rules, []string{from, to})
  }

  state := []string{".#.", "..#", "###"}

  for round := 0; round < 5; round++ {
    newstate := []string{}

    div := 2
    if len(state) % 2 != 0 {
      div = 3 
    }

    for ytile := 0; ytile < len(state)/div; ytile++ {
      indices := []int{}

      for xtile := 0; xtile < len(state)/div; xtile++ {
        key := []byte{}
        for y := 0; y < div; y++ {
          for x := 0; x < div; x++ {
            key = append(key, state[ytile*div+y][xtile*div+x])
          }
          if y < div-1 {
            key = append(key, '/')
          }
        }
        skey := string(key)

        found := false
        for i, r := range rules {
          if r[0] == skey                      ||
             r[0] == rot(skey)                 ||
             r[0] == rot(rot(skey))            ||
             r[0] == rot(rot(rot(skey)))       ||
             r[0] == flip(skey)                ||
             r[0] == rot(flip(skey))           ||
             r[0] == rot(rot(flip(skey)))      ||
             r[0] == rot(rot(rot(flip(skey)))) {
            indices = append(indices, i)
            found = true
            break
          }
        }
        if !found {
          fmt.Println("ERR")
          return
        }
      }

      for y := 0; y <= div; y++ {
        row := []byte{}
        for _, i := range indices {
          for x := 0; x <= div; x++ {
            row = append(row, rules[i][1][y*(div+2)+x])
          }
        }
        newstate = append(newstate, string(row))
      }
    }

    state = newstate

    count := 0
    for _, l := range state {
      for _, c := range l {
        if c == '#' {
          count++
        }
      }
    }
    fmt.Println(count)
  }
}

