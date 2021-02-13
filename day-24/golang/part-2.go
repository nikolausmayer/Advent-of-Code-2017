package main
import (
  "bufio"
  "fmt"
  "os"
  "regexp"
  "strconv"
)

type component struct {
  ports [2]int
}

func inList(i int, list []int) bool {
  for _, j := range list {
    if i == j {
      return true
    }
  }
  return false
}

func copyList(list []int) []int {
  result := []int{}
  for _, i := range list {
    result = append(result, i)
  }
  return result
}

func copyListApp(list []int, extra int) []int {
  l := copyList(list)
  return append(l, extra)
}


func main() {
  if len(os.Args) < 2 { return }
  infile, err := os.Open(os.Args[1])
  if err != nil { return }

  components := []component{}

  re, _ := regexp.Compile(`(\d+)/(\d+)`)
  scanner := bufio.NewScanner(infile)
  for scanner.Scan() {
    match := re.FindAllStringSubmatch(scanner.Text(), -1)
    a, _ := strconv.Atoi(match[0][1])
    b, _ := strconv.Atoi(match[0][2])
    components = append(components, component{[2]int{a,b}})
  }

  longest := 0
  strongestLongest := 0

  var recurser func(used []int, lastPort int)
  recurser = func(used []int, lastPort int) {
    if len(used) >= longest {
      strength := 0
      for _, i := range used {
        strength += components[i].ports[0]
        strength += components[i].ports[1]
      }
      if len(used) > longest {
        longest = len(used)
        strongestLongest = 0
      }
      if strength >= strongestLongest {
        strongestLongest = strength
      }
    }

    for i, c := range components {
      if inList(i, used) { continue }
      if c.ports[0] == lastPort {
        recurser(copyListApp(used, i), c.ports[1])
      } else if c.ports[1] == lastPort {
        recurser(copyListApp(used, i), c.ports[0])
      }
    }
  }
  recurser([]int{}, 0)

  fmt.Println(strongestLongest)
}

