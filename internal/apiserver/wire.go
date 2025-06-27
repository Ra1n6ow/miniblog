//go:build wireinject
// +build wireinject

package apiserver

import (
	"github.com/google/wire"

	"github.com/ra1n6ow/miniblog/internal/apiserver/biz"
	"github.com/ra1n6ow/miniblog/internal/apiserver/pkg/validation"
	"github.com/ra1n6ow/miniblog/internal/apiserver/store"
	ginmw "github.com/ra1n6ow/miniblog/internal/pkg/middleware/gin"
	"github.com/ra1n6ow/miniblog/internal/pkg/server"
	"github.com/ra1n6ow/miniblog/pkg/auth"
)

func InitializeWebServer(*Config) (server.Server, error) {
	wire.Build(
		wire.NewSet(NewWebServer, wire.FieldsOf(new(*Config), "ServerMode")),
		wire.Struct(new(ServerConfig), "*"), // * 表示注入全部字段
		wire.NewSet(store.ProviderSet, biz.ProviderSet),
		ProvideDB, // 提供数据库实例
		validation.ProviderSet,
		wire.NewSet(
			wire.Struct(new(UserRetriever), "*"),
			wire.Bind(new(ginmw.UserRetriever), new(*UserRetriever)),
		),
		auth.ProviderSet,
	)
	return nil, nil
}
