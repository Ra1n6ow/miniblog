// Copyright (c) 2024 ra1n6ow <jeffduuu@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/ra1n6ow/miniblog.

package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/ra1n6ow/miniblog/internal/pkg/core"
	"github.com/ra1n6ow/miniblog/internal/pkg/errno"
	"github.com/ra1n6ow/miniblog/internal/pkg/known"
	"github.com/ra1n6ow/miniblog/pkg/token"
)

// Authn 是一个 Gin 中间件，用来在每一个 HTTP 请求的 context, response 中注入 `X-Request-ID` 键值对.
func Authn() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 解析 JWT Token
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
