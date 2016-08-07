package main

import (
  "testing"
)

func TestNewDatabase(t *testing.T) {
  database := newDatabase(10)
  if database.Data != 10 {
    t.Error("Data failed initialize")
  }
  if database.State != "init" {
    t.Error("State failed initialize")
  }
  if database.ChangeData != 0 {
    t.Error("ChangeData failed initialize")
  }
  if database.IsLocked != false {
    t.Error("IsLocked failed initialize")
  }
  if database.IsCurrent != true {
    t.Error("IsCurrent failed initialize")
  }
  if database.IsChange != false {
    t.Error("IsChange failed initialize")
  }
}

func TestTransaction(t *testing.T) {
  database := newDatabase(10)
  isCurrent = true
  transaction(database, 5)
  if database.ChangeData != 5 {
    t.Error("not change ChangeData")
  }
  if database.IsChange != true {
    t.Error("database isn't ChangeMode")
  }
}

func TestPrepare(t *testing.T) {
  /* case 1 */
  database1 := newDatabase(10)
  database2 := newDatabase(5)
  isCurrent = true
  transaction(database1, 100)
  transaction(database2, 200)
  clients = make([]*database, 2)
  clients[0] = database1
  clients[1] = database2
  result := prepare()
  if result == 0 {
    if database1.State != "ready" && database2.State != "ready" {
      t.Error("database isn't ready.")
    }
    if database1.IsLocked != true && database2.IsLocked != true {
      t.Error("one database isn't locked")
    }
    if isPrepare != true {
      t.Error("coordinator isn't ready.")
    }
  } else {
    t.Error("votingSystem is failed.")
  }

  /* case 2 */
  database3 := newDatabase(10)
  database4 := newDatabase(5)
  isCurrent = true
  transaction(database3, 100)
  transaction(database4, 200)
  clients[0] = database3
  clients[1] = database4
  database4.IsCurrent = false
  result2 := prepare()
  if result2 == 1 {
    if database4.State != "init" {
      t.Error("database state should 'init'")
    }
    if database4.IsLocked != false {
      t.Error("database should not locked")
    }
    if isPrepare != false {
      t.Error("coordinator should not commit")
    }
  } else {
    t.Error("votingSystem is failed.")
  }
}

func TestCommit(t *testing.T) {
  database1 := newDatabase(10)
  database2 := newDatabase(5)
  isCurrent = true
  transaction(database1, 100)
  transaction(database2, 200)
  clients = make([]*database, 2)
  clients[0] = database1
  clients[1] = database2
  prepare()
  commit()
  if database1.Data != 100 && database2.Data != 200 {
    t.Error("missing commit! (Data)")
  }
  if database1.State != "init" && database2.State != "init" {
    t.Error("missing commit! (State)")
  }
  if database1.IsChange != false && database2.IsChange != false {
    t.Error("missing commit! (IsChange)")
  }
  if database1.IsLocked != false && database2.IsLocked != false {
    t.Error("missing commit! (IsLocked)")
  }
  if isPrepare != false {
    t.Error("missing commit! (isPrepare)")
  }
}

func TestAbort(t *testing.T) {
  database3 := newDatabase(10)
  database4 := newDatabase(5)
  isCurrent = true
  transaction(database3, 100)
  transaction(database4, 200)
  clients = make([]*database, 2)
  clients[0] = database3
  clients[1] = database4
  database4.IsCurrent = false
  prepare()
  abort()
  if database3.Data != 10 && database4.Data != 5 {
    t.Error("missing Abort! (Data)")
  }
  if database3.State != "init" && database4.State != "init" {
    t.Error("missing Abort! (State)")
  }
  if database3.IsChange != false && database4.IsChange != false {
    t.Error("missing Abort! (IsChange)")
  }
  if database3.IsLocked != false && database4.IsLocked != false {
    t.Error("missing Abort! (IsLocked)")
  }
  if isPrepare != false {
    t.Error("missing Abort! (isPrepare)")
  }
}
