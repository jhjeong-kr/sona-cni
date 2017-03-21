package main

import (
	log "plugin-to-SONA/log"

	//	"github.com/containernetworking/cni/pkg/invoke"
	"github.com/containernetworking/cni/pkg/skel"
	"github.com/containernetworking/cni/pkg/version"
	//	"github.com/containernetworking/cni/pkg/types"
)

func cmdAdd(args *skel.CmdArgs) error {
	return nil
}

func cmdDel(args *skel.CmdArgs) error {
	return nil
}

func main() {
	log.Info("sona cni is invoked")
	//	config.ParseCommandLine()
	skel.PluginMain(cmdAdd, cmdDel, version.All)
	log.Info("sona cni is terminated")
}
