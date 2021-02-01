package cmd

import (
	"bufio"
	"io/ioutil"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"regexp"
	"strings"
	"syscall"
)

func execCommand(command string, args []string, env []string) error {
	cmd := exec.Command(command, args...)
	cmd.Env = env
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	ch := make(chan os.Signal, 1)
	signal.Notify(ch)

	if err := cmd.Start(); err != nil {
		return err
	}

	go func() {
		for {
			sig := <-ch
			cmd.Process.Signal(sig)
		}
	}()

	if err := cmd.Wait(); err != nil {
		cmd.Process.Signal(os.Kill)
		return err
	}

	status := cmd.ProcessState.Sys().(syscall.WaitStatus)
	os.Exit(status.ExitStatus())
	return nil
}

// argsAfterTerminator returns all arguments after the
// terminator '--' as a slice of strings.
func argsAfterTerminator() []string {
	for i, s := range os.Args {
		if s == "--" {
			return os.Args[i+1:]
		}
	}
	return nil
}

// getcwd returns the current directory the executable
// is executed from.
func getcwd() (string, error) {
	return os.Getwd()
}

// getenv reads an envfile and returns the environments
// variables as slice of strings with the format ["key=value"].
//   1. Lines prefixed with # will be skipped
//   2. Lines prefixed with 'export' will be normalized
func getenv(path string) ([]string, error) {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	sc := bufio.NewScanner(strings.NewReader(string(f)))
	var lines []string
	re := regexp.MustCompile("^export\\s")
	for sc.Scan() {
		line := sc.Text()
		if strings.HasPrefix(line, "#") {
			continue
		}
		line = re.ReplaceAllString(line, "")
		lines = append(lines, line)
	}
	return lines, nil
}

// absPath returns the absolute path of a file.
// If the input file path is absolute it will just return it.
// If the path isn't it tries to resolve it from the path
// of where this executed from.
func absPath(file string) (string, error) {
	if filepath.IsAbs(file) {
		return filepath.Clean(file), nil
	}
	cwd, err := getcwd()
	if err != nil {
		return "", err
	}

	return filepath.Clean(filepath.Join(cwd, file)), nil
}
