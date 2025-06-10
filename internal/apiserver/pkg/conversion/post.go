package conversion

import (
	"github.com/ra1n6ow/gpkg/pkg/core"

	"github.com/ra1n6ow/miniblog/internal/apiserver/model"
	apiv1 "github.com/ra1n6ow/miniblog/pkg/api/apiserver/v1"
)

// PostModelToPostV1 将模型层的 PostM（博客模型对象）转换为 Protobuf 层的 Post（v1 博客对象）.
func PostModelToPostV1(postModel *model.PostM) *apiv1.Post {
	var protoPost apiv1.Post
	_ = core.CopyWithConverters(&protoPost, postModel)
	return &protoPost
}

// PostV1ToPostModel 将 Protobuf 层的 Post（v1 博客对象）转换为模型层的 PostM（博客模型对象）.
func PostV1ToPostModel(protoPost *apiv1.Post) *model.PostM {
	var postModel model.PostM
	_ = core.CopyWithConverters(&postModel, protoPost)
	return &postModel
}
