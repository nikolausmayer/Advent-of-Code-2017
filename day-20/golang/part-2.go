package main
import (
  "bufio"
  "fmt"
  "os"
  "regexp"
  "strconv"
)

type Vec3 struct {
  x int
  y int
  z int
}

type Particle struct {
  alive bool
  p Vec3
  v Vec3
  a Vec3
}

func main() {
  if len(os.Args) < 2 { return }
  infile, err := os.Open(os.Args[1])
  if err != nil { return }
  scanner := bufio.NewScanner(infile)

  Particles := []*Particle{}
  
  re, _ := regexp.Compile(`p=<(-?\d+),(-?\d+),(-?\d+)>, v=<(-?\d+),(-?\d+),(-?\d+)>, a=<(-?\d+),(-?\d+),(-?\d+)>`)
  for scanner.Scan() {
    match := re.FindAllStringSubmatch(scanner.Text(), -1)
    px, _ := strconv.Atoi(match[0][1])
    py, _ := strconv.Atoi(match[0][2])
    pz, _ := strconv.Atoi(match[0][3])
    vx, _ := strconv.Atoi(match[0][4])
    vy, _ := strconv.Atoi(match[0][5])
    vz, _ := strconv.Atoi(match[0][6])
    ax, _ := strconv.Atoi(match[0][7])
    ay, _ := strconv.Atoi(match[0][8])
    az, _ := strconv.Atoi(match[0][9])
    Particles = append(Particles, &Particle{true, Vec3{px, py, pz}, Vec3{vx, vy, vz}, Vec3{ax, ay, az}})
  }

  lastCollisionAt := 0
  for iter := 0;; iter++ {
    takenPositions := make(map[Vec3]int)
    for _, p := range Particles {
      if p.alive {
        takenPositions[p.p]++
      }
    }
    for _, p := range Particles {
      if p.alive {
        count, _ := takenPositions[p.p]
        if count > 1 {
          p.alive = false
          lastCollisionAt = iter
          //fmt.Println(p, "died at", iter)
        }
      }
    }

    /// Extreme eyeballing going on right here
    if iter - lastCollisionAt >= 100 {
      left := 0
      for _, p := range Particles {
        if p.alive {
          left++
        }
      }
      fmt.Println(left)
      break
    }

    for _, p := range Particles {
      p.v.x += p.a.x
      p.v.y += p.a.y
      p.v.z += p.a.z

      p.p.x += p.v.x
      p.p.y += p.v.y
      p.p.z += p.v.z
    }
  }
}

