package cmd

import (
	"github.com/spf13/cobra"

	"github.com/lin-snow/ech0/internal/cli"
)

// serveCmd 是启动 Web 和 SSH 服务的命令
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "启动 Web 和 SSH 服务",
	Run: func(cmd *cobra.Command, args []string) {
		cli.DoSSH()
		cli.DoServeWithBlock()
	},
}

// init 函数用于初始化根命令和子命令
func init() {
	rootCmd.AddCommand(serveCmd)
}
