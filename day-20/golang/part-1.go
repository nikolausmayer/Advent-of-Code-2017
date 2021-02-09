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
    Particles = append(Particles, &Particle{Vec3{px, py, pz}, Vec3{vx, vy, vz}, Vec3{ax, ay, az}})
  }

  minidx, minspeed := 0, 1 << 20
  for i, p := range Particles {
    speed := p.a.x * p.a.x + p.a.y * p.a.y + p.a.z * p.a.z;
    if speed < minspeed {
      minidx, minspeed = i, speed
    }
  }

  fmt.Println(minidx)
}

