package app

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"github.com/ArkArk/dcp/internal/comp"
	"github.com/ArkArk/dcp/internal/logger"
	"github.com/ArkArk/dcp/internal/util"
	"github.com/urfave/cli"
)

func getBashComplete() cli.BashCompleteFunc {
	return func(c *cli.Context) {
		completion := comp.GetCompletion(c)

		arg := currentArg(completion)
		logger.Write(logger.DEBUG, "currentArg: %#v", arg)

		containers, err := runningContainers()
		if err != nil {
			logger.Write(logger.ERROR, "%#v", err)
			os.Exit(1)
		}
		logger.Write(logger.DEBUG, "containers: %#v", containers)

		if strings.HasPrefix(arg, "-") {
			completeStd()
			completeOptions()
		} else if strings.Contains(arg, ":") {
			argContainer, argPath := splitArg(arg)
			if util.Contains(containers, argContainer) {
				completeContainerPaths(argContainer, argPath, completion)
			}
		} else {
			completeStd()
			completeLocalPaths(arg)
			completeContainers(containers)
		}
	}
}

// stdin/stdout
func completeStd() {
	fmt.Println("-")
}

// List options
func completeOptions() {
	fmt.Println("-a --archive")
	fmt.Println("-L --follow-link")
	fmt.Println("-h --help")
	fmt.Println("-v --version")
}

// List local files/directories
func completeLocalPaths(arg string) {
	argDir, argFile := splitPath(arg)

	files, err := ioutil.ReadDir(util.IfElse(argDir == "", ".", argDir))
	if err != nil {
		return
	}

	for _, file := range files {
		path := fmt.Sprintf("%s%s", argDir, file.Name())
		if file.IsDir() {
			path = path + "/"
		}
		fmt.Println(path)
	}

	if strings.HasPrefix(argFile, ".") {
		for _, filename := range []string{"./", "../"} {
			path := fmt.Sprintf("%s%s", argDir, filename)
			fmt.Println(path)
		}
	}
}

// List containers
func completeContainers(containers []string) {
	for _, container := range containers {
		fmt.Printf("%s:\n", container)
	}
}

// List container files/directories
func completeContainerPaths(argContainer string, argPath string, completion comp.Completion) {
	argDir, argFile := splitPath(argPath)
	if !strings.HasPrefix(argDir, "/") {
		argDir = "/" + argDir
	}

	out, err := exec.Command("docker", "container", "exec", argContainer, "ls", "-p", argDir).CombinedOutput()

	if err != nil {
		logger.Write(logger.WARN, "%#v", err)
		return
	}

	for _, file := range strings.Fields(string(out)) {
		path := fmt.Sprintf("%s%s", argDir[1:], file)
		if completion.Current == ":" {
			path = ":" + path
		} else if strings.HasPrefix(argPath, "/") {
			path = "/" + path
		}
		fmt.Println(path)
	}

	if strings.HasPrefix(argFile, ".") {
		for _, filename := range []string{"./", "../"} {
			path := fmt.Sprintf("%s%s", argDir[1:], filename)
			if completion.Current == ":" {
				path = ":" + path
			} else if strings.HasPrefix(argPath, "/") {
				path = "/" + path
			}
			logger.Write(logger.WARN, "%v", path)
			fmt.Println(path)
		}
	}
}

func splitArg(arg string) (container, path string) {
	parts := strings.SplitN(arg, ":", 2)
	container = parts[0]
	path = parts[1]
	return
}

func splitPath(path string) (dir, file string) {
	parts := strings.Split(path, "/")
	file = parts[len(parts)-1]
	dir = strings.Join(parts[0:len(parts)-1], "/")
	if dir != "" {
		dir = dir + "/"
	}
	return
}

func runningContainers() ([]string, error) {
	out, err := exec.Command("docker", "ps", `--format="{{.Names}}"`).CombinedOutput()

	if err != nil {
		return []string{}, err
	}

	containers := []string{}
	for _, c := range strings.Fields(string(out)) {
		if len(c) >= 2 && c[0] == '"' && c[len(c)-1] == '"' {
			// Remove double quotes
			c = c[1 : len(c)-1]
		}
		containers = append(containers, c)
	}
	return containers, nil
}

func currentArg(completion comp.Completion) string {
	line := completion.Line
	point := completion.Point
	args := strings.Split(line, " ")
	acc := 0
	for _, arg := range args {
		l := len(arg)
		if point <= acc+l {
			return line[acc:point]
		}
		acc += l + 1
	}
	// unreachable code
	return ""
}
