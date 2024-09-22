// Copyright Gidi233 <qpbtyfh@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/Gidi233/Gd-Blog.

package post

import (
	"github.com/gin-gonic/gin"

	"github.com/Gidi233/Gd-Blog/internal/pkg/core"
	"github.com/Gidi233/Gd-Blog/internal/pkg/known"
	"github.com/Gidi233/Gd-Blog/internal/pkg/log"
)

// DeleteCollection 批量删除博客.
func (ctrl *PostController) DeleteCollection(c *gin.Context) {
	log.C(c).Infow("Batch delete post function called")

	postIDs := c.QueryArray("postID")
	if err := ctrl.b.Posts().DeleteCollection(c, c.GetString(known.XUsernameKey), postIDs); err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, nil)
}
