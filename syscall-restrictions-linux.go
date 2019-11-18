// +build linux

package main

import (
	"fmt"
	"syscall"

	libseccomp "github.com/seccomp/libseccomp-golang"
)

func applySyscallRestrictions() {
	var syscalls = []string{"read", "write", "close", "mmap", "munmap",
		"rt_sigaction", "rt_sigprocmask", "clone", "execve", "sigaltstack",
		"arch_prctl", "gettid", "futex", "sched_getaffinity", "epoll_ctl",
		"openat", "newfstatat", "readlinkat", "pselect6", "epoll_pwait",
		"epoll_create1", "exit_group"}
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
