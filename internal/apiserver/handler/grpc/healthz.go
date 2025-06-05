// Copyright 2024 Ra1n6ow <jeffduuu@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/Ra1n6ow/miniblog.

package grpc

import (
	"context"
	"time"

	emptypb "google.golang.org/protobuf/types/known/emptypb"

	"github.com/ra1n6ow/miniblog/internal/pkg/log"
	apiv1 "github.com/ra1n6ow/miniblog/pkg/api/apiserver/v1"
)

// Healthz 服务健康检查.
func (h *Handler) Healthz(ctx context.Context, rq *emptypb.Empty) (*apiv1.HealthzResponse, error) {
	log.W(ctx).Infow("Healthz handler is called", "method", "Healthz", "status", "healthy")
	return &apiv1.HealthzResponse{
		Status:    apiv1.ServiceStatus_Healthy,
		Timestamp: time.Now().Format(time.DateTime),
	}, nil
}
