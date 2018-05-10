package main

import (
	"fmt"
	"flag"
	"os"
	"os/user"
	"os/exec"
	"syscall"
	"strings"
	"strconv"
)

var (
	err error
	env = os.Environ()

	fSetsid bool
	fPreserveEnv bool
	fLogin bool
)

func parseArgs() {
	flag.Usage = func() {
		fmt.Printf("%v: [options] user[:group] command [args...]\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
	}

	flag.BoolVar(&fSetsid, "s", false, "Create new session (setsid)")
	flag.BoolVar(&fPreserveEnv, "p", false, "Preserve environment")
	flag.BoolVar(&fLogin, "l", false, "Run the command as a login shell, envrionment is cleared and \"-p\" flag is ignored")
	flag.Parse()

	if fLogin {
		fPreserveEnv = false
	}

	if len(flag.Args()) < 2 {
		flag.Usage()
	}
}

func envDel(key string) {
	var envNew = []string{}
	for _, e := range env {
		pair := strings.Split(e, "=")
		if pair[0] != key {
			envNew = append(envNew, e)
		}
	}
	env = envNew
}

func envGet(key string) (string, bool) {
	for _, e := range env {
		pair := strings.Split(e, "=")
		if pair[0] == key && len(pair) > 1 {
			return pair[1], true
		}
	}
	return "", false
}

func envSet(key, val string) {
	envDel(key)
	env = append(env, strings.Join([]string{key, val}, "="))
}

func main() {
	var (
		strArr []string
		uid uint32
		gid uint32
		groups []uint32
		pArgs []string
		pCmd string
		p *os.Process
		ps *os.ProcessState
	)

	parseArgs()

	if fLogin {
		var v, ok = envGet("TERM")
		env = []string{}
		if ok {
			envSet("TERM", v)
		}
		envSet("PATH", "/usr/local/sbin:/usr/local/bin:/sbin:/bin:/usr/sbin:/usr/bin")
	}

	strArr = strings.Split(flag.Args()[0], ":")

	if i, err := strconv.Atoi(strArr[0]); err == nil {
		uid = uint32(i)
	} else if u, err := user.Lookup(strArr[0]); err == nil {
		i, _ = strconv.Atoi(u.Uid)
		uid = uint32(i)
	} else {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	if len(strArr) > 1 && strArr[1] != "" {
		if i, err := strconv.Atoi(strArr[1]); err == nil {
			gid = uint32(i)
		} else if g, err := user.LookupGroup(strArr[1]); err == nil {
			i, _ = strconv.Atoi(g.Gid)
			gid = uint32(i)
		} else {
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}
	} else {
		gid = uid
	}

	if u, err := user.LookupId(strconv.FormatUint(uint64(uid), 10)); err == nil {
		if ug, err := u.GroupIds(); err == nil {
			for _, v := range ug {
				if i, err := strconv.Atoi(v); err == nil {
					groups = append(groups, uint32(i))
				}
			}
		}
		if !fPreserveEnv {
			envSet("HOME", u.HomeDir)
			envSet("USER", u.Username)
		}
	}

	pArgs = make([]string, len(flag.Args()[1:]))
	copy(pArgs, flag.Args()[1:])

	pCmd, err = exec.LookPath(pArgs[0])
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	if fLogin {
		pArgs[0] = "-"
	}

	p, err = os.StartProcess(pCmd, pArgs, &os.ProcAttr{
		Env: env,
		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
		Sys: &syscall.SysProcAttr{
			Credential: &syscall.Credential{
				Uid: uid,
				Gid: gid,
				Groups: groups,
			},
			Setsid: fSetsid,
		},
	})
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	ps, err = p.Wait()
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	if ps.Exited() {
		os.Exit(ps.Sys().(syscall.WaitStatus).ExitStatus())
	}
}
