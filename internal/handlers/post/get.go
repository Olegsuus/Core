package handlers

import (
	"context"
	"fmt"
	"github.com/Olegsuus/Core/internal/models"
	postpb "github.com/Olegsuus/Core/settings_grpc/go/core/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log/slog"
)

func (h *PostGRPCHandler) GetManyPosts(ctx context.Context, req *postpb.GetManyPostsRequest) (*postpb.GetManyPostsResponse, error) {
	limit := req.GetLimit()
	page := req.GetPage()

	offset := (page - 1) * limit

	settings := models.GetManyPostSettings{
		Limit:  int(limit),
		Offset: int(offset),
	}

	posts, err := h.psP.ServiceGetMany(ctx, settings)
	if err != nil {
		h.l.Debug("ошибка при получении списка постов", slog.String("error:", fmt.Sprintf("%w", err)))
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

func (h *PostGRPCHandler) GetFeed(ctx context.Context, req *postpb.GetFeedRequest) (*postpb.GetFeedResponse, error) {
	subscriberID := req.GetUserId()
	limit := req.GetLimit()
	page := req.GetPage()
	offset := (page - 1) * limit

	settings := models.GetManyPostSettings{
		Limit:  int(limit),
		Offset: int(offset),
	}

	posts, err := h.psP.ServiceGetFeed(ctx, subscriberID, settings)
	if err != nil {
		h.l.Debug("ошибка при получении ленты постов", slog.String("error", fmt.Sprintf("%w", err)))
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

	return &postpb.GetFeedResponse{
		Posts: pbPosts,
	}, nil
}
