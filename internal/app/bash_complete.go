package app

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/ArkArk/dcp/internal/comp"
	"github.com/ArkArk/dcp/internal/docker"
	"github.com/ArkArk/dcp/internal/logger"
	"github.com/ArkArk/dcp/internal/util"
	"github.com/urfave/cli"
)

func getBashComplete() cli.BashCompleteFunc {
	return func(c *cli.Context) {
		completion := comp.GetCompletion(c)

		arg := currentArg(completion)
		logger.Write(logger.DEBUG, "currentArg: %#v", arg)

		containers, err := runningContainerNames()
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

	result, err := docker.Exec(argContainer, []string{"ls", "-p", argDir})

	if err != nil {
		logger.Write(logger.WARN, "%#v", err)
		return
	}
	if result.ExitCode != 0 {
		logger.Write(logger.DEBUG, "%#v", result)
		logger.Write(logger.ERROR, "%#v", result.Stderr)
		return
	}

	for _, file := range strings.Fields(result.Stdout) {
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

func runningContainerNames() ([]string, error) {
	containers, err := docker.ContainerList()
	if err != nil {
		return nil, err
	}

	var containerNames []string
	for _, container := range containers {
		for _, name := range container.Names {
			// E.g. name == "/foo"
			containerNames = append(containerNames, name[1:])
		}
	}

	return containerNames, nil
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
