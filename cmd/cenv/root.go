package cmd

import (
	"flag"
	"fmt"
	"os"

	"github.com/hellozimi/cenv/pkg/version"
)

const usage = `cenv <env-file> -- <command> [<args>...]`

func Execute() error {
	vflag := flag.Bool("v", false, "prints current cenv version")
	flag.Parse()
	if *vflag {
		fmt.Println(version.Version())
		return nil
	}

	args := os.Args[1:]
	if len(args) < 3 {
		fmt.Println(usage)
		flag.PrintDefaults()
		return nil
	}
	envFile, err := absPath(args[0])
	if err != nil {
		return err
	}

	cmdStrings := argsAfterTerminator()
	cmdArgs := cmdStrings[1:]
	cmd := cmdStrings[0]
	env, err := getenv(envFile)
	if err != nil {
		return err
	}

	env = append(os.Environ(), env...)

	return execCommand(cmd, cmdArgs, env)
}
