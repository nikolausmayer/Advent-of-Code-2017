package main
import (
  "bufio"
  "fmt"
  "os"
  "strconv"
  "strings"
)

func isPrime(n int) bool {
  for i := 2; i*i <= n; i++ {
    if n % i == 0 {
      return false
    }
  }
  return true
}

func nextPrime(n int) int {
  n++
  for ; !isPrime(n); n = nextPrime(n) { }
  return n
}

func primeFactors(n int) map[int]int {
  results := make(map[int]int)
  for i := 2; n > 1; {
    if n % i == 0 {
      results[i]++
      n /= i
    } else {
      i++
    }
  }
  return results
}

func coPeriod(n int, m int) int {
  pfn := primeFactors(n)
  pfm := primeFactors(m)

  for pf, pow := range pfm {
    if pow > pfn[pf] {
      pfn[pf] = pow
    }
  }

  result := 1
  for pf, pow := range pfn {
    for i := 0; i < pow; i++ {
      result *= pf
    }
  }
  return result
}


func main() {
  if len(os.Args) < 2 {
    return
  }
  infile, err := os.Open(os.Args[1])
  if err != nil { return }

  scanners := make(map[int]int)

  scanner := bufio.NewScanner(infile)
  for scanner.Scan() {
    line := strings.Fields(scanner.Text())
    depth, _  := strconv.Atoi(line[0][0:len(line[0])-1])
    length, _ := strconv.Atoi(line[1])
    scanners[depth] = length
  }

  //start := 1
  //period := 1

  //for d, p := range scanners {
  //  if d % 2*(p-1) == 0 {
  //    start++
  //  }
  //  period = coPeriod(period, 2*(p-1))
  //}

  for i := 0;; i++ {
    stop := true
    for d, p := range scanners {
      if (d+i) % (2*(p-1)) == 0 {
        stop = false
        break
      }
    }
    if stop {
      fmt.Println(i)
      break
    }
  }
}

