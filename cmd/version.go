/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"runtime/debug"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "打印版本和构建信息",
	Long:  `显示当前二进制文件的版本、Commit Hash 以及构建时间等信息。`,
	Run: func(cmd *cobra.Command, args []string) {
		// 尝试读取构建信息
		if info, ok := debug.ReadBuildInfo(); ok {
			fmt.Printf("Module:  %s\n", info.Main.Path)
			fmt.Printf("Version: %s\n", info.Main.Version)
			
			var revision, time string
			for _, setting := range info.Settings {
				switch setting.Key {
				case "vcs.revision":
					revision = setting.Value
				case "vcs.time":
					time = setting.Value
				}
			}
			
			if revision != "" {
				fmt.Printf("Commit:  %s\n", revision)
			}
			if time != "" {
				fmt.Printf("Built:   %s\n", time)
			}
		} else {
			fmt.Println("无法读取构建信息 (需要 Go 1.18+ 构建)")
		}
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
