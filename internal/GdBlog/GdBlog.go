// Copyright Gidi233 <qpbtyfh@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/Gidi233/Gd-Blog.

package GdBlog

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Gidi233/Gd-Blog/internal/pkg/log"
	"github.com/Gidi233/Gd-Blog/pkg/version/verflag"
	"github.com/gin-gonic/gin"
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
			// 如果 `--version=true`，则打印版本并退出
			verflag.PrintAndExitIfRequested()
			log.Init(logOptions())
			defer log.Sync()

			return run()
		},
		//
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				return fmt.Errorf("%q does not take any arguments, got %q", cmd.CommandPath(), args)
			}

			return nil
		},
	}

	cobra.OnInitialize(initConfig)

	cmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "The path to the miniblog configuration file. Empty string for no configuration file.")

	cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	verflag.AddFlags(cmd.PersistentFlags())

	return cmd
}

func run() error {
	gin.SetMode(viper.GetString("runmode"))

	g := gin.New()

	g.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"code": 10003, "message": "Page not found."})
	})

	g.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	httpsrv := &http.Server{Addr: viper.GetString("addr"), Handler: g}

	log.Infow("Start to listening the incoming requests on http address", "addr", viper.GetString("addr"))
	if err := httpsrv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalw(err.Error())
	}
	return nil
}
