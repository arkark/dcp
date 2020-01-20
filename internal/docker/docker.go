package docker

import (
	"bytes"
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
)

type ExecResult struct {
	ExitCode int
	Stdout   string
	Stderr   string
}

func ContainerList() ([]types.Container, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return nil, err
	}
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		return nil, err
	}
	return containers, nil
}

func Exec(containerName string, cmd []string) (ExecResult, error) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return ExecResult{}, err
	}

	execConfig := types.ExecConfig{
		AttachStdout: true,
		AttachStderr: true,
		Cmd:          cmd,
	}
	cresp, err := cli.ContainerExecCreate(ctx, containerName, execConfig)
	if err != nil {
		return ExecResult{}, err
	}
	execID := cresp.ID

	aresp, err := cli.ContainerExecAttach(ctx, execID, types.ExecStartCheck{})
	if err != nil {
		return ExecResult{}, err
	}
	defer aresp.Close()

	var stdoutBuf, stderrBuf bytes.Buffer
	outputDone := make(chan error, 1)

	go func() {
		_, err = stdcopy.StdCopy(&stdoutBuf, &stderrBuf, aresp.Reader)
		outputDone <- err
	}()

	select {
	case err := <-outputDone:
		if err != nil {
			return ExecResult{}, err
		}
		break
	case <-ctx.Done():
		return ExecResult{}, ctx.Err()
	}

	iresp, err := cli.ContainerExecInspect(ctx, execID)
	if err != nil {
		return ExecResult{}, err
	}

	return ExecResult{
		ExitCode: iresp.ExitCode,
		Stdout:   stdoutBuf.String(),
		Stderr:   stderrBuf.String(),
	}, nil
}
