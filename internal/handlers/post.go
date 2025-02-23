package handlers

import (
	"context"
	"fmt"
	models "github.com/Olegsuus/Core/internal/service"
	postpb "github.com/Olegsuus/Core/settings_grpc/go/core/proto"
	"log/slog"
	"time"
)

func (h *GRPCHandlers) AddPost(ctx context.Context, req *postpb.AddPostRequest) (*postpb.AddPostResponse, error) {
	post, err := h.service.AddPost(ctx, models.AddPostParam{
		ID:        "",
		Content:   req.GetContent(),
		Title:     req.GetTitle(),
		UserID:    req.GetUserId(),
		CreatedAt: time.Now(),
	})

	if err != nil {
		h.l.Error("ошибка при добавлении нового поста", slog.String("error:", fmt.Sprintf("%s", err)))
		return nil, fmt.Errorf("Service.AddPost: %w", err)
	}
	return &postpb.AddPostResponse{Id: post.ID}, nil
}

func (h *GRPCHandlers) GetPost(ctx context.Context, req *postpb.GetPostRequest) (*postpb.GetPostResponse, error) {
	postId := req.GetId()

	post, err := h.service.GetPost(ctx, postId)
	if err != nil {
		h.l.Error("ошибка при получении поста", slog.String("err", fmt.Sprintf("%s", err)))
		return nil, fmt.Errorf("Service.GetPost: %w", err)
	}

	return &postpb.GetPostResponse{
		Post: postModelsToGRPC(*post),
	}, nil
}

func (h *GRPCHandlers) GetFeed(ctx context.Context, req *postpb.GetFeedRequest) (*postpb.GetFeedResponse, error) {
	subscriberID := req.GetUserId()
	limit := req.GetLimit()
	page := req.GetPage()
	offset := (page - 1) * limit

	posts, err := h.service.GetFeed(ctx, subscriberID, models.GetManyParam{
		Limit: int(limit), Offset: int(offset),
	})

	if err != nil {
		h.l.Error("ошибка при получении ленты постов", slog.String("error", fmt.Sprintf("%s", err)))
		return nil, fmt.Errorf("Service.GetFeed: %w", err)
	}

	var postPB []*postpb.Post
	for _, post := range posts {
		postPB = append(postPB, postModelsToGRPC(post))
	}

	return &postpb.GetFeedResponse{
		Posts: postPB,
	}, nil
}

func (h *GRPCHandlers) GetManyPosts(ctx context.Context, req *postpb.GetManyPostsRequest) (*postpb.GetManyPostsResponse, error) {
	limit := req.GetLimit()
	page := req.GetPage()
	offset := (page - 1) * limit

	posts, err := h.service.GetManyPosts(ctx, models.GetManyParam{
		Limit:  int(limit),
		Offset: int(offset),
		Order:  true,
	})

	if err != nil {
		h.l.Error("ошибка при получении списка постов", slog.String("error:", fmt.Sprintf("%s", err)))
		return nil, fmt.Errorf("Service.GetManyPosts: %w", err)
	}

	var postPB []*postpb.Post
	for _, post := range posts {
		postPB = append(postPB, postModelsToGRPC(post))
	}

	return &postpb.GetManyPostsResponse{
		Posts: postPB,
	}, nil
}

func (h *GRPCHandlers) RemovePost(ctx context.Context, req *postpb.RemovePostRequest) (*postpb.RemovePostResponse, error) {
	err := h.service.RemovePost(ctx, req.GetId())
	if err != nil {
		h.l.Error("ошибка при удалении нового поста", slog.String("error:", fmt.Sprintf("%s", err)))
		return nil, fmt.Errorf("Service.RemovePost: %w", err)
	}
	return &postpb.RemovePostResponse{Success: true}, nil
}
