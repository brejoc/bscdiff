// +build linux

package main

import (
	"fmt"
	"syscall"

	libseccomp "github.com/seccomp/libseccomp-golang"
)

func applySyscallRestrictions() {

	var syscalls = []string{"read", "write", "close", "fstat", "mmap",
		"mprotect", "munmap", "brk", "rt_sigaction", "rt_sigprocmask",
		"access", "nanosleep", "clone", "execve", "uname", "fcntl",
		"sigaltstack", "arch_prctl", "gettid", "futex", "sched_getaffinity",
		"set_tid_address", "epoll_ctl", "openat", "newfstatat",
		"readlinkat", "set_robust_list", "epoll_create1", "pipe2",
		"prlimit64", "exit_group"}
	whiteList(syscalls)
}

// Load the seccomp whitelist.
func whiteList(syscalls []string) {

	filter, err := libseccomp.NewFilter(
		libseccomp.ActErrno.SetReturnCode(int16(syscall.EPERM)))
	if err != nil {
		fmt.Printf("Error creating filter: %s\n", err)
	}
	for _, element := range syscalls {
		// fmt.Printf("[+] Whitelisting: %s\n", element)
		syscallID, err := libseccomp.GetSyscallFromName(element)
		if err != nil {
			panic(err)
		}
		filter.AddRule(syscallID, libseccomp.ActAllow)
	}
	filter.Load()
}
