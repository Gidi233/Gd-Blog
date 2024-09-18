// Copyright Gidi233 <qpbtyfh@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/Gidi233/Gd-Blog.

package user

import (
	"context"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/Gidi233/Gd-Blog/internal/pkg/log"
	pb "github.com/Gidi233/Gd-Blog/pkg/proto/GdBlog/v1"
)

// ListUser 返回用户列表，只有 root 用户才能获取用户列表.
func (ctrl *UserController) ListUser(ctx context.Context, r *pb.ListUserRequest) (*pb.ListUserResponse, error) {
	log.C(ctx).Infow("ListUser function called")

	resp, err := ctrl.b.Users().List(ctx, int(r.Offset), int(r.Limit))
	if err != nil {
		return nil, err
	}

	users := make([]*pb.UserInfo, 0, len(resp.Users))
	for _, u := range resp.Users {
		createdAt, _ := time.Parse("2006-01-02 15:04:05", u.CreatedAt)
		updatedAt, _ := time.Parse("2006-01-02 15:04:05", u.UpdatedAt)
		users = append(users, &pb.UserInfo{
			Username:  u.Username,
			Nickname:  u.Nickname,
			Email:     u.Email,
			Phone:     u.Phone,
			PostCount: u.PostCount,
			CreatedAt: timestamppb.New(createdAt),
			UpdatedAt: timestamppb.New(updatedAt),
		})
	}

	ret := &pb.ListUserResponse{
		TotalCount: resp.TotalCount,
		Users:      users,
	}

	return ret, nil
}
