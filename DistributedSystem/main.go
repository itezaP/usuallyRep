/*
 author : Masahiro Ohtomo
 date : 2016/08/05
 main.go : This program writes test-cases to run for 2-phase commit.
*/

package main

import (
  "fmt"
)

func main() {
  database1 := newDatabase(10)
  database2 := newDatabase(5)
  clients = make([]*database,2)
  clients[0] = database1
  clients[1] = database2
  isCurrent = true

  /* transaction */
  transaction(clients[0], 100)
  transaction(clients[1], 10)

  /* 2-phase commit */
  res := prepare()
  if (res == 0) {
    commit()
  } else if (res == 1) {
    abort()
  } else if (res == -1) {
    fmt.Println("coordinator failed.")
  }
}
