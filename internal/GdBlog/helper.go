// Copyright Gidi233 <qpbtyfh@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/Gidi233/Gd-Blog.

package GdBlog

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/Gidi233/Gd-Blog/internal/pkg/log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	// 默认目录.
	recommendedHomeDir = ".GdBlog"

	// 默认配置文件名.
	defaultConfigName = "GdBlog.yaml"
)

// initConfig 设置需要读取的配置文件名、环境变量，并读取配置文件内容到 viper 中.
func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(filepath.Join(home, recommendedHomeDir))

		viper.AddConfigPath(".")

		viper.SetConfigType("yaml")

		viper.SetConfigName(defaultConfigName)
	}

	viper.AutomaticEnv()

	viper.SetEnvPrefix("MINIBLOG")

	replacer := strings.NewReplacer(".", "_", "-", "_")
	viper.SetEnvKeyReplacer(replacer)

	if err := viper.ReadInConfig(); err != nil {
		log.Errorw("Failed to read viper configuration file", "err", err)
	}
	log.Debugw("Using config file", "file", viper.ConfigFileUsed())

}

// logOptions 从 viper 中读取日志配置
func logOptions() *log.Options {
	return &log.Options{
		DisableCaller:     viper.GetBool("log.disable-caller"),
		DisableStacktrace: viper.GetBool("log.disable-stacktrace"),
		Level:             viper.GetString("log.level"),
		Format:            viper.GetString("log.format"),
		OutputPaths:       viper.GetStringSlice("log.output-paths"),
	}
}
