package main

import (
	"os"
	"os/exec"
	"path/filepath"
)

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}

func runLoggedMust(cmd *exec.Cmd) {
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	panicIfErr(err)
}

func main() {
	{
		path := filepath.Join("cmd", "gen", "gen.js")
		cmd := exec.Command("bun", path)
		runLoggedMust(cmd)
	}
	{
		cmd := exec.Command("go", "test", ".")
		runLoggedMust(cmd)
	}
}
