// Copyright Gidi233 <qpbtyfh@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/Gidi233/Gd-Blog.

package middleware

import (
	"github.com/gin-gonic/gin"

	"github.com/Gidi233/Gd-Blog/internal/pkg/core"
	"github.com/Gidi233/Gd-Blog/internal/pkg/errno"
	"github.com/Gidi233/Gd-Blog/internal/pkg/known"
	"github.com/Gidi233/Gd-Blog/pkg/token"
)

// Authn 是认证中间件，用来从 gin.Context 中提取 token 并验证 token 是否合法，
// 如果合法则将 token 中的用户名存放在 gin.Context 的 XUsernameKey 键中.
func Authn() gin.HandlerFunc {
	return func(c *gin.Context) {
		username, err := token.ParseRequest(c)
		if err != nil {
			core.WriteResponse(c, errno.ErrTokenInvalid, nil)
			c.Abort()

			return
		}

		c.Set(known.XUsernameKey, username)
		c.Next()
	}
}
