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
  /* test case 1*/
  /* initialized */
  database1 := newDatabase(10,"init")
  database2 := newDatabase(5,"init")
  clients = make([]*database,2)
  clients[0] = database1
  clients[1] = database2

  fmt.Println("database1 : ", clients[0].Data)
  fmt.Println("database2 : ", clients[1].Data)

  /* transaction */
  transaction(clients[0], 100)
  transaction(clients[1], 10)

  /* 2-phase commit */
  res := prepare()
  if (res == 0) {
    commit()
  } else {
    abort()
  }

  fmt.Println("database1 : ", clients[0].Data)
  fmt.Println("database2 : ", clients[1].Data)
}
