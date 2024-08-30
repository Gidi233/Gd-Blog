// Copyright Gidi233 <qpbtyfh@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/Gidi233/Gd-Blog.

package GdBlog

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

func NewGdBlogCommand() *cobra.Command {
	cmd := &cobra.Command{
		// specify the name of the command
		Use: "Gd-Blog",
		// specify the short description of the command
		Short: "Gd-Blog is a simple blog system",
		// specify the long description of the command
		Long: `Gd-Blog is a simple but not easy blog system,
Find more Gd-Blog information at:
	https://github.com/Gidi233/Gd-Blog#readme`,

		// when error occurs, the command will not print usage information
		SilenceUsage: true,
		// specify the run function to execute when cmd.Execute() is called
		// if the function fails, an error message will be returned
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
		//
		Args: func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				if len(arg) > 0 {
					return fmt.Errorf("%q does not take any arguments, got %q", cmd.CommandPath(), args)
				}
			}

			return nil
		},
	}

	// 以下设置，使得 initConfig 函数在每个命令运行时都会被调用以读取配置
	cobra.OnInitialize(initConfig)

	// 在这里您将定义标志和配置设置。

	// Cobra 支持持久性标志(PersistentFlag)，该标志可用于它所分配的命令以及该命令下的每个子命令
	cmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "The path to the miniblog configuration file. Empty string for no configuration file.")

	// Cobra 也支持本地标志，本地标志只能在其所绑定的命令上使用
	cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	return cmd
}

func run() error {
	// 打印所有的配置项及其值
	settings, _ := json.Marshal(viper.AllSettings())
	fmt.Println(string(settings))
	// 打印 db -> username 配置项的值
	fmt.Println(viper.GetString("db.username"))
	return nil
}
