// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by seccomp-profiler - DO NOT EDIT.

// +build linux,386

package include

import (
	"github.com/elastic/go-seccomp-bpf"

	beat "github.com/elastic/beats/libbeat/common/seccomp"
)

func init() {
	beat.MustRegisterPolicy(&seccomp.Policy{
		DefaultAction: seccomp.ActionErrno,
		Syscalls: []seccomp.SyscallGroup{
			{
				Action: seccomp.ActionAllow,
				Names: []string{
					"_llseek",
					"access",
					"brk",
					"clock_gettime",
					"clone",
					"close",
					"dup",
					"dup2",
					"epoll_create",
					"epoll_create1",
					"epoll_ctl",
					"epoll_wait",
					"exit",
					"exit_group",
					"fchdir",
					"fchmod",
					"fchown32",
					"fcntl",
					"fcntl64",
					"fdatasync",
					"flock",
					"fstat64",
					"fsync",
					"ftruncate64",
					"futex",
					"getcwd",
					"getdents",
					"getdents64",
					"getpid",
					"getppid",
					"getrandom",
					"getrlimit",
					"getrusage",
					"gettid",
					"gettimeofday",
					"ioctl",
					"kill",
					"lstat64",
					"madvise",
					"mincore",
					"mkdirat",
					"mmap2",
					"mprotect",
					"munmap",
					"nanosleep",
					"open",
					"openat",
					"pipe",
					"pipe2",
					"poll",
					"pread64",
					"prlimit64",
					"pwrite64",
					"read",
					"readlink",
					"readlinkat",
					"rename",
					"renameat",
					"rt_sigaction",
					"rt_sigprocmask",
					"rt_sigreturn",
					"sched_getaffinity",
					"sched_yield",
					"sendfile64",
					"set_robust_list",
					"set_thread_area",
					"setitimer",
					"sigaltstack",
					"socketcall",
					"splice",
					"stat",
					"stat64",
					"tgkill",
					"time",
					"tkill",
					"uname",
					"unlink",
					"unlinkat",
					"wait4",
					"write",
					"writev",
				},
			},
		},
	})
}
