package main

import (
	"github.com/docker/machine/libmachine/drivers/plugin"
	"github.com/sergiomacedo/docker-machine-driver-cloudstack"
)

func main() {
	plugin.RegisterDriver(cloudstack.NewDriver("", ""))
}
