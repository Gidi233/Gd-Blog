// Copyright Gidi233 <qpbtyfh@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/Gidi233/Gd-Blog.

package post

import (
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"

	"github.com/Gidi233/Gd-Blog/internal/pkg/core"
	"github.com/Gidi233/Gd-Blog/internal/pkg/errno"
	"github.com/Gidi233/Gd-Blog/internal/pkg/known"
	"github.com/Gidi233/Gd-Blog/internal/pkg/log"
	v1 "github.com/Gidi233/Gd-Blog/pkg/api/GdBlog/v1"
)

// Create 创建一条博客.
func (ctrl *PostController) Create(c *gin.Context) {
	log.C(c).Infow("Create post function called")

	var r v1.CreatePostRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, errno.ErrBind, nil)

		return
	}

	if _, err := govalidator.ValidateStruct(r); err != nil {
		core.WriteResponse(c, errno.ErrInvalidParameter.SetMessage(err.Error()), nil)

		return
	}

	resp, err := ctrl.b.Posts().Create(c, c.GetString(known.XUsernameKey), &r)
	if err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, resp)
}
