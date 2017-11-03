package common

import (
	"fmt"
	"os/exec"
)

func MustRunCommand(name string, args ...string) {
	output, err := exec.Command(name, args...).CombinedOutput()
	if err != nil {
		fmt.Println(output)
		panic(err)
	}
}

func MustFmtFile(filename string) {
	MustRunCommand("gofmt", "-w", filename)
}

func MustGoreturns(filename string) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("----------------------------------------")
			fmt.Println("Make sure you have installed goreturns")
			fmt.Println("go get -u -v sourcegraph.com/sqs/goreturns")
			fmt.Println("----------------------------------------")
			panic(err)
		}
	}()
	MustRunCommand("goreturns", "-w", filename)
}
