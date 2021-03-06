/*
Copyright © 2020 Flanksource
This file is part of Flanksource build tools
*/
package cmd

import (
	"github.com/philipstaffordwood/build-tools/cmd/test"
	"github.com/stretchr/testify/assert"
	"testing"
)


func TestGetRootCmd(t *testing.T) {
	cmd := GetRootCommand()
	assert.NotNil(t,cmd, "We must have a root command")
}

func TestRoot_HasGhSubcommand(t *testing.T) {
	cmd := GetRootCommand()
	test.HasSubcommand(t,cmd,"gh","We want a gh subcommand")
}


