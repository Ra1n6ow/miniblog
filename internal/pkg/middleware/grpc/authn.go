package grpc

import (
	"context"

	"google.golang.org/grpc"

	"github.com/ra1n6ow/miniblog/internal/apiserver/model"
	"github.com/ra1n6ow/miniblog/internal/pkg/contextx"
	"github.com/ra1n6ow/miniblog/internal/pkg/errno"
	"github.com/ra1n6ow/miniblog/internal/pkg/known"
	"github.com/ra1n6ow/miniblog/internal/pkg/log"
	"github.com/ra1n6ow/miniblog/pkg/token"
)

// UserRetriever 用于根据用户名获取用户信息的接口.
type UserRetriever interface {
	// GetUser 根据用户 ID 获取用户信息
	GetUser(ctx context.Context, userID string) (*model.UserM, error)
}

// AuthnInterceptor 是一个 gRPC 拦截器，用于进行认证.
func AuthnInterceptor(retriever UserRetriever) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		// 解析 JWT Token
		userID, err := token.ParseRequest(ctx)
		if err != nil {
			log.Errorw("Failed to parse request", "err", err)
			return nil, errno.ErrTokenInvalid.WithMessage("%s", err.Error())
		}

		log.Debugw("Token parsing successful", "userID", userID)

		user, err := retriever.GetUser(ctx, userID)
		if err != nil {
			return nil, errno.ErrUnauthenticated.WithMessage("%s", err.Error())
		}

		// 将用户信息存入上下文
		//nolint:staticcheck
		ctx = context.WithValue(ctx, known.XUsername, user.Username)
		//nolint:staticcheck
		ctx = context.WithValue(ctx, known.XUserID, userID)

		// 供 log 和 contextx 使用
		ctx = contextx.WithUserID(ctx, user.UserID)
		ctx = contextx.WithUsername(ctx, user.Username)

		// 继续处理请求
		return handler(ctx, req)
	}
}
