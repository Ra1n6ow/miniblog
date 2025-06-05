// Copyright 2024 Ra1n6ow <jeffduuu@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/Ra1n6ow/miniblog.

package errno

import (
	"net/http"

	"github.com/ra1n6ow/gpkg/pkg/errorsx"
)

// ErrPostNotFound 表示未找到指定的博客.
var ErrPostNotFound = &errorsx.ErrorX{Code: http.StatusNotFound, Reason: "NotFound.PostNotFound", Message: "Post not found."}
