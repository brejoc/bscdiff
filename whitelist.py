# You can get a list of syscalls via strace:
# $ strace -qcf ./bscdiff

dump = """\
read
write
close
fstat
mmap
mprotect
munmap
brk
rt_sigaction
rt_sigprocmask
access
nanosleep
clone
execve
uname
fcntl
sigaltstack
arch_prctl
gettid
futex
sched_getaffinity
set_tid_address
epoll_ctl
openat
newfstatat
readlinkat
set_robust_list
epoll_create1
pipe2
prlimit64"""

whitelist = dump.split("\n")
whitelist.append("exit_group")  # I guess we alwas need to exit the program
output = ['"{}"'.format(elem) for elem in whitelist ]
output = ", ".join(output)

print(output)