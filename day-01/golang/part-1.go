package main
import("fmt"; "os"; "io/ioutil")

func main() {
  if len(os.Args) < 2 {
    return
  }
  file, err := os.Open(os.Args[1])
  if err != nil {
    return
  }
  data, err := ioutil.ReadAll(file)
  data[len(data)-1] = data[0]

  count := 0
  for i := 0; i < len(data)-1; i++ {
    if data[i] == data[i+1] {
      count += int(data[i]) - int('0')
    }
  }
  fmt.Println(count)
}

