/*
  author : Masahiro Ohtomo
  date : 2016/08/05
  database.go : This program expression virtual database.
                This program Also say database class.
*/

package main

type database struct {
  Data int           // Value of database
  ChangeData int     // will commit value
  State string       // operation state
  IsLocked bool      // change denied runtime transaction
  IsCurrent bool     // databese works or not works(failed)?
  IsChange bool      // change in value?
}

/* constructor */
func newDatabase(data int) *database{
  database := &database{ Data: data,
                         ChangeData: 0,
                         State: "init",
                         IsLocked: false,
                         IsCurrent: true,
                         IsChange: false }
  return database
}
