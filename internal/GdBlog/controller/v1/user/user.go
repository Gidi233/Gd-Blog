// Copyright Gidi233 <qpbtyfh@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/Gidi233/Gd-Blog.

package user

import (
	"github.com/Gidi233/Gd-Blog/internal/GdBlog/biz"
	"github.com/Gidi233/Gd-Blog/internal/GdBlog/store"
	"github.com/Gidi233/Gd-Blog/pkg/auth"
	pb "github.com/Gidi233/Gd-Blog/pkg/proto/GdBlog/v1"
)

// UserController 是 user 模块在 Controller 层的实现，用来处理用户模块的请求.
type UserController struct {
	a *auth.Authz
	b biz.IBiz
	pb.UnimplementedGdBlogServer
}

func New(ds store.IStore, a *auth.Authz) *UserController {
	return &UserController{a: a, b: biz.NewBiz(ds)}
}
