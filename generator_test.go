package main

import (
	"github.com/cuisongliu/container-install/install/command"
	"testing"
)

func TestDocker_fetch(t *testing.T) {
	d := &command.Docker{}
	d.Fetch()
}

func TestContainerd_fetch(t *testing.T) {
	d := &command.Containerd{}
	d.Fetch()
}
