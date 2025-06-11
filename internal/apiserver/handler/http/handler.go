// Copyright 2024 Ra1n6ow <jeffduuu@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/Ra1n6ow/miniblog.

package http

import "github.com/ra1n6ow/miniblog/internal/apiserver/biz"

// Handler 处理博客模块的请求.
type Handler struct {
	biz biz.IBiz
}

// NewHandler 创建新的 Handler 实例.
func NewHandler(biz biz.IBiz) *Handler {
	return &Handler{
		biz: biz,
	}
}
