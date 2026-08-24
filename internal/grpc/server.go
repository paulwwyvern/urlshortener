package grpc

import (
	"context"
	"errors"

	"github.com/paulwwyvern/urlshortener/api/proto"
	"github.com/paulwwyvern/urlshortener/internal/model/dto"
	"github.com/paulwwyvern/urlshortener/internal/model/errs"
	"github.com/paulwwyvern/urlshortener/pkg/httphelpers/httpuser"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ShortenerService interface {
	GetURL(ctx context.Context, shortURL string) (string, error)
	GetUserURLs(ctx context.Context, userID int32) ([]dto.GetUserURLResponse, error)
	GenerateURL(ctx context.Context, userID int32, url string) (string, error)
}

type Server struct {
	proto.UnimplementedShortenerServiceServer

	logger  *zap.Logger
	service ShortenerService
}

func NewServer(logger *zap.Logger, service ShortenerService) *Server {
	return &Server{
		logger:  logger,
		service: service,
	}
}

func (s *Server) ShortenURL(ctx context.Context, req *proto.URLShortenRequest) (*proto.URLShortenResponse, error) {
	userID := httpuser.GetUserID(ctx)
	url := req.GetUrl()

	shortURL, err := s.service.GenerateURL(ctx, userID, url)
	if err != nil {
		if errors.Is(err, errs.ErrOriginalURLAlreadyExists) {
			return nil, status.Error(codes.AlreadyExists, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	response := proto.URLShortenResponse_builder{
		Result: shortURL,
	}
	return response.Build(), nil
}

func (s *Server) ExpandURL(ctx context.Context, req *proto.URLExpandRequest) (*proto.URLExpandResponse, error) {
	shortUrl := req.GetId()

	url, err := s.service.GetURL(ctx, shortUrl)
	if err != nil {
		if errors.Is(err, errs.ErrShortURLNotFound) || errors.Is(err, errs.ErrShortURLGone) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	response := proto.URLExpandResponse_builder{
		Result: url,
	}
	return response.Build(), nil
}

func (s *Server) ListUserURLs(ctx context.Context, _ *emptypb.Empty) (*proto.UserURLsResponse, error) {
	userID := httpuser.GetUserID(ctx)

	urls, err := s.service.GetUserURLs(ctx, userID)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var urlsData []*proto.URLData

	for _, url := range urls {
		urlData := proto.URLData_builder{
			ShortUrl:    url.ShortURL,
			OriginalUrl: url.OriginalURL,
		}
		urlsData = append(urlsData, urlData.Build())
	}

	response := proto.UserURLsResponse_builder{
		Url: urlsData,
	}
	return response.Build(), nil

}
