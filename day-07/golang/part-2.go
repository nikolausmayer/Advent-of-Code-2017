package main
import (
  "bufio"
  "fmt"
  "regexp"
  "os"
  "strconv"
)

func main() {
  if len(os.Args) < 2 {
    return
  }
  infile, err := os.Open(os.Args[1])
  if err != nil { return }
  defer infile.Close()
  scanner := bufio.NewScanner(infile)

  reOuter, _ := regexp.Compile(`^(\w+) \((\d+)\)(?: -> )?(.*)`)
  reInner, _ := regexp.Compile(`\w+`)

  allPrograms := make(map[string]int)
  topPrograms := make(map[string][]string)
  notBottomPrograms := make(map[string]struct{})

  for scanner.Scan() {
    line := scanner.Text()
    rOuter := reOuter.FindAllStringSubmatch(line, -1)
    weight, _ := strconv.Atoi(rOuter[0][2])
    a := rOuter[0][1]
    allPrograms[a] = weight
    if len(rOuter[0][3]) > 0 {
      rInner := reInner.FindAllString(rOuter[0][3], -1)
      for _, b := range rInner {
        topPrograms[a] = append(topPrograms[a], b)
        notBottomPrograms[b] = struct{}{}
      }
    }
  }

  sumWeights := make(map[string]int)
  
  var computeWeightSum func(s string) int
  computeWeightSum = func(s string) int {
    v, ok := sumWeights[s]
    if !ok {
      v = allPrograms[s]
      for _, onTop := range topPrograms[s] {
        v += computeWeightSum(onTop)
      }
      sumWeights[s] = v
    }
    return v
  }
  root := "x"
  for p := range allPrograms {
    _, ok := notBottomPrograms[p]
    if !ok {
      root = p
      break
    }
  }
  computeWeightSum(root)

  wrongNode := ""
  for {
    //fmt.Println(root)
    // Build histogram of next-level weights
    topWeights := make(map[int]int)
    for _, onTop := range topPrograms[root] {
      topWeights[sumWeights[onTop]]++
    }
    // Find out-of-balance child
    next := ""
    for k, v := range topWeights {
      if v == 1 {
        for _, onTop := range topPrograms[root] {
          if k == sumWeights[onTop] {
            next = onTop
            break
          }
        }
        break
      }
    }
    // If no child is out of balance, we are at the goal...
    if next == "" {
      wrongNode = root
      break
    } else {
      // ...else we recurse
      root = next
    }
  }  

  for p := range allPrograms {
    for _, c := range topPrograms[p] {
      if c == wrongNode {
        topWeights := make(map[int]int)
        for _, onTop := range topPrograms[p] {
          topWeights[sumWeights[onTop]]++
        }
        for k, v := range topWeights {
          if v != 1 {
            targetWeight := k
            for _, onTop := range topPrograms[wrongNode] {
              targetWeight -= sumWeights[onTop]
            }
            fmt.Println(targetWeight)
            return
          }
        }
      }
    }
  }

  fmt.Println()
}

