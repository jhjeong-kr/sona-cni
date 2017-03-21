package main

import (
	"encoding/json"
	"os"

	log "plugin-to-SONA/log"

	//	"github.com/containernetworking/cni/pkg/invoke"
	"github.com/containernetworking/cni/pkg/skel"
	"github.com/containernetworking/cni/pkg/version"
	//	"github.com/containernetworking/cni/pkg/types"
)

func cmdAdd(args *skel.CmdArgs) error {
	log.Info("cni: add")
	return nil
}

func cmdDel(args *skel.CmdArgs) error {
	log.Info("cni: delete")
	return nil
}

func main() {
	log.Info(os.Args[0], " is invoked")
	if e := skel.PluginMainWithError(cmdAdd, cmdDel, version.All); e != nil {
		log.Error(os.Args[0], " returns with the following: ", prettyString(e))
		if err := printWithReturn(e); err != nil {
			log.Error("Error writing error JSON to stdout: ", err)
		}
	} else {
		log.Info(os.Args[0], " returns the result: ")
	}
	os.Exit(1)
}

// This function is modified from .../cni/pkg/types.Result object
func printWithReturn(obj interface{}) error {
	data, err := json.MarshalIndent(obj, "", "    ")
	if err != nil {
		return err
	}
	_, err = os.Stdout.Write(append(data, ([]byte("\n"))...))
	return err
}

func prettyString(obj interface{}) string {
	data, err := json.Marshal(obj)
	if err != nil {
		return "conversion error"
	}
	return string(data)
}
