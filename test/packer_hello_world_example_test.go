package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wahlfeld/terratest/modules/docker"
	"github.com/wahlfeld/terratest/modules/packer"
)

func TestPackerHelloWorldExample(t *testing.T) {
	packerOptions := &packer.Options{
		// website::tag::1:: The path to where the Packer template is located
		Template: "../examples/packer-hello-world-example/build.pkr.hcl",
	}

	// website::tag::2:: Build the Packer template. This template will create a Docker image.
	packer.BuildArtifact(t, packerOptions)

	// website::tag::3:: Run the Docker image, read the text file from it, and make sure it contains the expected output.
	opts := &docker.RunOptions{Command: []string{"cat", "/test.txt"}}
	output := docker.Run(t, "gruntwork/packer-hello-world-example", opts)
	assert.Equal(t, "Hello, World!", output)
}
