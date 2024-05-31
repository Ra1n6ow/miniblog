// Copyright (c) 2024 ra1n6ow <jeffduuu@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/ra1n6ow/miniblog.

package user

import (
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"

	"github.com/ra1n6ow/miniblog/internal/pkg/core"
	"github.com/ra1n6ow/miniblog/internal/pkg/errno"
	"github.com/ra1n6ow/miniblog/internal/pkg/log"
	v1 "github.com/ra1n6ow/miniblog/pkg/api/miniblog/v1"
)

func (ctrl *UserController) Login(c *gin.Context) {
	log.C(c).Infow("Login user function called")
	var req v1.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		core.WriteResponse(c, errno.ErrBind, nil)
		return
	}
	if _, err := govalidator.ValidateStruct(req); err != nil {
		core.WriteResponse(c, errno.ErrInvalidParameter.SetMessage(err.Error()), nil)
		return
	}
	rsp, err := ctrl.b.Users().Login(c, &req)
	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}
	core.WriteResponse(c, nil, rsp)
}
