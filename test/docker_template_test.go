package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/docker"
	"github.com/stretchr/testify/assert"
)

func TestDockerNameSuccess(t *testing.T) {

	tag := "hadenlabs/docker-template:latest"
	otherOptions := []string{
		"--no-cache",
	}
	buildOptions := &docker.BuildOptions{
		Tags:         []string{tag},
		OtherOptions: otherOptions,
	}

	docker.Build(t, "../images/0.0.0", buildOptions)

	opts := &docker.RunOptions{
		Command: []string{},
	}
	output := docker.Run(t, tag, opts)
	assert.NotEmpty(t, output, output)
}
