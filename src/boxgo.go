package main

import (
  "os"
  "os/exec"
  "fmt"
  "syscall"
)

func Run() {
  // make the command

  cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
  cmd.Stdin = os.Stdin
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr

  // isolation
  // SysProcAttr is struct

  cmd.SysProcAttr = &syscall.SysProcAttr {
    Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWUSER | syscall.CLONE_NEWPID,                  // creates unprivileged user with new uts, new pid, new randon user
  }

  // run the command

  ClearOut(cmd.Run())
}

func Child() {
  fmt.Printf("executing: %v as %d\n", os.Args[2:], os.Getpid())

  syscall.Sethostname([]byte("box"))

  // make the command
  cmd := exec.Command(os.Args[2], os.Args[3:]...)
  cmd.Stdin = os.Stdin
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr

  // run the commmand

  ClearOut(cmd.Run())
}

func ClearOut (err error) {
  if err != nil {
    panic(err)
  }
}
