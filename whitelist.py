# You can get a list of syscalls via strace:
# $ strace -qcf ./team-suse

dump = """\
read
write
close
mmap
munmap
rt_sigaction
rt_sigprocmask
clone
execve
sigaltstack
arch_prctl
gettid
futex
sched_getaffinity
epoll_ctl
openat
newfstatat
readlinkat
pselect6
epoll_pwait
epoll_create1"""

whitelist = dump.split("\n")
whitelist.append("exit_group")  # I guess we alwas need to exit the program
output = ['"{}"'.format(elem) for elem in whitelist ]
output = ", ".join(output)

print(output)