// Copyright Gidi233 <qpbtyfh@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/Gidi233/Gd-Blog.

package core

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Gidi233/Gd-Blog/internal/pkg/errno"
)

// 发生错误时的返回消息.
type ErrResponse struct {
	// 错误码.
	Code string `json:"code"`

	//错误信息.
	Message string `json:"message"`
}

// 将错误、响应数据写入 HTTP 响应主体。
func WriteResponse(c *gin.Context, err error, data interface{}) {
	if err != nil {
		hcode, code, message := errno.Decode(err)
		c.JSON(hcode, ErrResponse{
			Code:    code,
			Message: message,
		})

		return
	}

	c.JSON(http.StatusOK, data)
}
