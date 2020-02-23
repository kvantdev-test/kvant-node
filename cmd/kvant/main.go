package main

import (
	"github.com/kvant-node/cmd/kvant/cmd"
	"github.com/kvant-node/cmd/utils"
)

func main() {
	rootCmd := cmd.RootCmd

	rootCmd.AddCommand(
		cmd.RunNode,
		cmd.ShowNodeId,
		cmd.ShowValidator,
//		cmd.ManagerCommand,
//		cmd.ManagerConsole,
		cmd.VerifyGenesis,
		cmd.Version)

	rootCmd.PersistentFlags().StringVar(&utils.MinterHome, "home-dir", "", "base dir (default is $HOME/.kvant)")
	rootCmd.PersistentFlags().StringVar(&utils.MinterConfig, "config", "", "path to config (default is $(home-dir)/config/config.toml)")
	rootCmd.PersistentFlags().Bool("pprof", false, "enable pprof")
	rootCmd.PersistentFlags().String("pprof-addr", "0.0.0.0:6060", "pprof listen addr")

	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
