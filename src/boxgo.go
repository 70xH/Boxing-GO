package main

import (
  "os"
  "os/exec"
  "fmt"
  "syscall"
)

func Run () {
  fmt.Printf("executing: %v\n", os.Args[2:])

  // make the command

  cmd := exec.Command(os.Args[2], os.Args[3:]...)
  cmd.Stdin = os.Stdin
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr

  // isolation
  // SysProcAttr is struct

  cmd.SysProcAttr = &syscall.SysProcAttr {
    Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWUSER | syscall.CLONE_NEWPID,                                                               // creates unprivileged user with new uts
  }

  // run the command

  ClearOut(cmd.Run())
}

func ClearOut (err error) {
  if err != nil {
    panic(err)
  }
}
