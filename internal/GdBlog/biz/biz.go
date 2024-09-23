// Copyright Gidi233 <qpbtyfh@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/Gidi233/Gd-Blog.

package biz

//go:generate mockgen -destination mock_biz.go -package biz github.com/Gidi233/Gd-Blog/internal/GdBlog/biz IBiz

import (
	"github.com/Gidi233/Gd-Blog/internal/GdBlog/biz/post"
	"github.com/Gidi233/Gd-Blog/internal/GdBlog/biz/user"
	"github.com/Gidi233/Gd-Blog/internal/GdBlog/store"
)

type IBiz interface {
	Posts() post.PostBiz
	Users() user.UserBiz
}

var _ IBiz = (*biz)(nil)

type biz struct {
	ds store.IStore
}

var _ IBiz = (*biz)(nil)

func NewBiz(ds store.IStore) *biz {
	return &biz{ds: ds}
}

func (b *biz) Users() user.UserBiz {
	return user.New(b.ds)
}

func (b *biz) Posts() post.PostBiz {
	return post.New(b.ds)
}
