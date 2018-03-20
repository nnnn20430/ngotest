package main

import (
	"fmt"
	"flag"
	"os"
	"os/user"
	"syscall"
	"strings"
	"strconv"
)

func main() {
	var (
		setsid bool;
		strArr []string;
		uid uint32;
		gid uint32;
		groups []uint32;
		pArgs []string;
		p *os.Process;
		err error;
	)
	flag.BoolVar(&setsid, "s", false, "create new session (setsid)")
	flag.Parse()
	if len(flag.Args()) < 2 {
		fmt.Printf("%v: [-s] uid[:gid] command [args]\n", os.Args[0])
		os.Exit(1)
	}
	strArr = strings.Split(flag.Args()[0], ":")
	if i, err := strconv.Atoi(strArr[0]); err == nil {
		uid = uint32(i)
	} else {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
	if len(strArr) > 1 && strArr[1] != "" {
		if i, err := strconv.Atoi(strArr[1]); err == nil {
			gid = uint32(i)
		} else {
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}
	} else {
		gid = uid
	}
	if u, err := user.LookupId(strArr[0]); err == nil {
		if ug, err := u.GroupIds(); err == nil {
			for _, v := range ug {
				if i, err := strconv.Atoi(v); err == nil {
					groups = append(groups, uint32(i))
				}
			}
		}
	}
	pArgs = make([]string, len(flag.Args()[1:]))
	copy(pArgs, flag.Args()[1:])
	p, err = os.StartProcess(pArgs[0], pArgs, &os.ProcAttr{
		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
		Sys: &syscall.SysProcAttr{
			Credential: &syscall.Credential{
				Uid: uid,
				Gid: gid,
				Groups: groups,
			},
			Setsid: setsid,
		},
	})
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
	p.Wait()
}
