package main

import (
  "bytes"
  "flag"
  "fmt"
  "io/ioutil"

  "github.com/nsf/termbox-go"
)

func main() {
  flag.Parse()
  fn := flag.Arg(0)
  fb, err := ioutil.ReadFile(fn)
  if err != nil {
    panic(err)
  }
  buf := bytes.NewBuffer(fb)
  err = termbox.Init()
  if err != nil {
    panic(err)
  }
  defer termbox.Close()
  var done bool
  for {
    r, _, err := buf.ReadRune()
    if err != nil {
      done = true
    }
    e := termbox.PollEvent()
    if e.Key == termbox.KeyEsc {
      break
    }
    if done {
      continue
    }
    if int32(r) == 10 {
      fmt.Print("\r")
    }
    fmt.Printf("%c", r)
  }
}
