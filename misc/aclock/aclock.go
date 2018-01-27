package main

import (
	"fmt"
	"time"
	"os"
	"strings"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: aclock <format>")
		os.Exit(1)
	}
	for true {
		t := time.Now()
		z := time.FixedZone("myzone", 0)
		t = t.In(z)
		tF := strings.Replace(os.Args[1], "1136239445", "%s", -1)
		tF = t.Format(tF)
		tF = strings.Replace(tF, "%s", strconv.FormatInt(t.Unix(), 10), -1)
		fmt.Println(tF)
		time.Sleep(1 * time.Second)
	}
}
