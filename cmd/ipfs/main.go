package main

import (
	"github.com/ipfs/kubo/cmd/ipfs/kubo"
)

func mainRet() (exitCode int) {
	return kubo.Start(kubo.BuildDefaultEnv)
}
