// Copyright Gidi233 <qpbtyfh@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/Gidi233/Gd-Blog.

package main

import (
	"os"

	_ "go.uber.org/automaxprocs/maxprocs" // 程序自动设置 GOMAXPROCS 以匹配 Linux 容器 CPU 配额。

	"github.com/Gidi233/Gd-Blog/internal/GdBlog"
)

// program entry function.
func main() {
	command := GdBlog.NewGdBlogCommand()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
