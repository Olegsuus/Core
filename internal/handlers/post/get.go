package handlers

import (
	"context"
	"fmt"
	postpb "github.com/Olegsuus/Core/settings_grpc/go/core/proto"
	"log/slog"
)

func (h *PostGRPCHandler) GetPost(ctx context.Context, req *postpb.GetPostRequest) (*postpb.GetPostResponse, error) {
	postId := req.GetId()

	post, err := h.postService.GetPost(ctx, postId)
	if err != nil {
		h.l.Debug("ошибка при получении поста", slog.String("err", fmt.Sprintf("%s", err)))
		return nil, err
	}

	return &postpb.GetPostResponse{
		Post: post,
	}, nil
}

func (h *PostGRPCHandler) GetFeed(ctx context.Context, req *postpb.GetFeedRequest) (*postpb.GetFeedResponse, error) {
	subscriberID := req.GetUserId()
	limit := req.GetLimit()
	page := req.GetPage()
	offset := (page - 1) * limit

	posts, err := h.postService.GetFeed(ctx, subscriberID, int(limit), int(offset))
	if err != nil {
		h.l.Debug("ошибка при получении ленты постов", slog.String("error", fmt.Sprintf("%s", err)))
		return nil, err
	}

	return &postpb.GetFeedResponse{
		Posts: posts,
	}, nil
}

func (h *PostGRPCHandler) GetManyPosts(ctx context.Context, req *postpb.GetManyPostsRequest) (*postpb.GetManyPostsResponse, error) {
	limit := req.GetLimit()
	page := req.GetPage()
	offset := (page - 1) * limit

	posts, err := h.postService.GetManyPosts(ctx, int(limit), int(offset), true)
	if err != nil {
		h.l.Debug("ошибка при получении списка постов", slog.String("error:", fmt.Sprintf("%s", err)))
		return nil, err
	}

	return &postpb.GetManyPostsResponse{
		Posts: posts,
	}, nil
}
