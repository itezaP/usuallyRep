/*
  author : Masahiro Ohtomo
  date : 2016/08/05
  transactionManager.go : This program implements operaions in 2-phase commit.
*/

package main

import (
  "fmt"
)

var clients []*database
var isPrepare bool

/* transction oepration */
func transaction(database *database, data int) {
  if (database.IsCurrent == true && database.IsLocked == false) {
    database.ChangeData = data
    database.IsChange = true
  } else {
    fmt.Println("sorry, this database failed or locked.")
  }
}

/* confirm operaion before commit in 2-phase commit*/
func prepare() int{
  vote := 0
  transactionDatabases := 0

  /* check is operated databeses */
  for j := 0; j < len(clients); j++ {
    if (clients[j].IsChange == true) {
      transactionDatabases += 1
    }
  }

  /* responds confirm operation */
  for i := 0; i < len(clients); i++ {
    if (clients[i].IsCurrent == true && clients[i].State == "init" && clients[i].IsChange == true) {
      vote += 1
      clients[i].IsLocked = true
      clients[i].State = "ready"
      fmt.Println("database",i,"prepared!")
    } else {
      fmt.Println("database",i,"failed.")
    }
  }

  /* all member consensus*/
  if (vote == transactionDatabases) {
    fmt.Println("This transaction committing.")
    isPrepare = true
    return 0
  } else { /* if not */
    fmt.Println("This transaction aborting.")
    isPrepare = false
    return 1
  }
}

/* commit operation in 2-phase commit */
func commit() {
  if (isPrepare == true) {
    for i := 0; i < len(clients); i++ {
      if (clients[i].IsChange == true) {
        clients[i].State = "commit"
        clients[i].Data = clients[i].ChangeData
        clients[i].IsChange = false
        clients[i].State = "init"
        clients[i].IsLocked = false
        } else {
          fmt.Println("This database is latest version.")
        }
    }
  } else {
    fmt.Println("parmission is denied.")
  }
  isPrepare = false
}

/* abort oepration in 2-phase commit */
func abort() {
  for i := 0; i < len(clients); i++ {
    if (clients[i].IsChange == true) {
      clients[i].State = "abort"
      clients[i].IsChange = false
      clients[i].IsLocked = false
      clients[i].State = "init"
    } else {
      fmt.Println("This database is latest version.")
    }
  }
  isPrepare = false
}
