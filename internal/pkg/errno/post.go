package errno

import (
	"net/http"

	"github.com/ra1n6ow/gpkg/pkg/errorsx"
)

// ErrPostNotFound 表示未找到指定的博客.
var ErrPostNotFound = &errorsx.ErrorX{Code: http.StatusNotFound, Reason: "NotFound.PostNotFound", Message: "Post not found."}
