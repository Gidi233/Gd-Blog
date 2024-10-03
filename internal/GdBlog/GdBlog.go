// Copyright Gidi233 <qpbtyfh@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/Gidi233/Gd-Blog.

package GdBlog

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"google.golang.org/grpc"

	"github.com/Gidi233/Gd-Blog/internal/GdBlog/controller/v1/user"
	"github.com/Gidi233/Gd-Blog/internal/GdBlog/store"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/Gidi233/Gd-Blog/internal/pkg/known"
	"github.com/Gidi233/Gd-Blog/internal/pkg/log"
	mw "github.com/Gidi233/Gd-Blog/internal/pkg/middleware"
	pb "github.com/Gidi233/Gd-Blog/pkg/proto/GdBlog/v1"
	"github.com/Gidi233/Gd-Blog/pkg/token"
	"github.com/Gidi233/Gd-Blog/pkg/version/verflag"
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

	cmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "The path to the Gd-Blog configuration file. Empty string for no configuration file.")

	cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	verflag.AddFlags(cmd.PersistentFlags())

	return cmd
}

func run() error {
	if err := initStore(); err != nil {
		return err
	}

	token.Init(viper.GetString("jwt-secret"), known.XUsernameKey)

	gin.SetMode(viper.GetString("runmode"))

	g := gin.New()

	mws := []gin.HandlerFunc{gin.Recovery(), mw.NoCache, mw.Cors, mw.Secure, mw.RequestID()}
	g.Use(mws...)

	if err := installRouters(g); err != nil {
		return err
	}
	// 创建并运行 HTTP 服务器
	httpsrv := startInsecureServer(g)

	// 创建并运行 HTTPS 服务器
	httpssrv := startSecureServer(g)

	// 创建并运行 GRPC 服务器
	grpcsrv := startGRPCServer()

	quit := make(chan os.Signal, 1)
	// SIGKILL 信号，不能被捕获，所以不需要添加
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Infow("Shutting down server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := httpsrv.Shutdown(ctx); err != nil {
		log.Errorw("Insecure Server forced to shutdown", "err", err)
		return err
	}
	if err := httpssrv.Shutdown(ctx); err != nil {
		log.Errorw("Secure Server forced to shutdown", "err", err)
		return err
	}

	grpcsrv.GracefulStop()

	log.Infow("Server exiting")

	return nil
}

// startInsecureServer 创建并运行 HTTP 服务器.
func startInsecureServer(g *gin.Engine) *http.Server {
	httpsrv := &http.Server{Addr: viper.GetString("addr"), Handler: g}

	log.Infow("Start to listening the incoming requests on http address", "addr", viper.GetString("addr"))
	go func() {
		if err := httpsrv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalw(err.Error())
		}
	}()

	return httpsrv
}

// startSecureServer 创建并运行 HTTPS 服务器.
func startSecureServer(g *gin.Engine) *http.Server {
	httpssrv := &http.Server{Addr: viper.GetString("tls.addr"), Handler: g}

	log.Infow("Start to listening the incoming requests on https address", "addr", viper.GetString("tls.addr"))
	cert, key := viper.GetString("tls.cert"), viper.GetString("tls.key")
	if cert != "" && key != "" {
		go func() {
			if err := httpssrv.ListenAndServeTLS(cert, key); err != nil && !errors.Is(err, http.ErrServerClosed) {
				log.Fatalw(err.Error())
			}
		}()
	}

	return httpssrv
}

// startGRPCServer 创建并运行 GRPC 服务器.
func startGRPCServer() *grpc.Server {
	lis, err := net.Listen("tcp", viper.GetString("grpc.addr"))
	if err != nil {
		log.Fatalw("Failed to listen", "err", err)
	}

	grpcsrv := grpc.NewServer()
	pb.RegisterGdBlogServer(grpcsrv, user.New(store.S, nil))

	log.Infow("Start to listening the incoming requests on grpc address", "addr", viper.GetString("grpc.addr"))
	go func() {
		if err := grpcsrv.Serve(lis); err != nil {
			log.Fatalw(err.Error())
		}
	}()

	return grpcsrv
}
