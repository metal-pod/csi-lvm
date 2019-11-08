package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

const (
	flagLVName         = "lvname"
	flagLVSize         = "lvsize"
	flagVGName         = "vgname"
	flagDevicesPattern = "devices"
	flagDirectory      = "directory"
	flagLVMType        = "lvmtype"
	flagBlockMode      = "block"
)

func cmdNotFound(c *cli.Context, command string) {
	panic(fmt.Errorf("Unrecognized command: %s", command))
}

func onUsageError(c *cli.Context, err error, isSubcommand bool) error {
	panic(fmt.Errorf("Usage error, please check your command"))
}

func main() {
	p := cli.NewApp()
	p.Usage = "LVM Provisioner Pod"
	p.Commands = []cli.Command{
		createLVCmd(),
		deleteLVCmd(),
	}
	p.CommandNotFound = cmdNotFound
	p.OnUsageError = onUsageError

	log.Println("starting csi-lvm-provisioner")

	if err := p.Run(os.Args); err != nil {
		log.Fatalf("Critical error: %v", err)
	}
}
