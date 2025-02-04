package handlers

import (
	"context"
	"github.com/Olegsuus/Core/internal/domain/dto"
	postpb "github.com/Olegsuus/Core/settings_grpc/go/core/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (h *PostGRPCHandler) GetManyPosts(ctx context.Context, req *postpb.GetManyPostsRequest) (*postpb.GetManyPostsResponse, error) {
	limit := req.GetLimit()
	page := req.GetPage()

	offset := (page - 1) * limit

	settings := dto.GetManyPostSettings{
		Limit:    int(limit),
		Offset:   int(offset),
		SortDesc: req.GetSortDesc(),
	}

	posts, err := h.psP.GetMany(ctx, settings)
	if err != nil {
		return nil, err
	}

	var pbPosts []*postpb.Post
	for _, p := range posts {
		pbPosts = append(pbPosts, &postpb.Post{
			Id:        p.ID,
			Title:     p.Title,
			Content:   p.Content,
			CreatedAt: timestamppb.New(p.CreatedAt),
		})
	}

	return &postpb.GetManyPostsResponse{
		Posts: pbPosts,
	}, nil
}
