package main

import (
	"os"
	
	"github.com/ipfs/kubo/cmd/ipfs/kubo"
)

func main() {
	os.Exit(mainRet())
}

func mainRet() (exitCode int) {
	return kubo.Start(kubo.BuildDefaultEnv)
}
