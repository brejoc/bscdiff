// +build !linux

package main

// We only have seccomp for linux right now.
func appylSyscallRestrictions() {
}
