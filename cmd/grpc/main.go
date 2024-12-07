package main

import (
	"context"
	"log"
	"net"

	"example.com/internal/lyrics/dto"
	lyricsR "example.com/internal/lyrics/repository/implementation"
	"example.com/internal/lyrics/service"
	lyricsS "example.com/internal/lyrics/service/implementation"
	pb "example.com/pkg/protos"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const port = ":9000"

type server struct {
	pb.UnimplementedLyricsServiceServer
	service service.LyricsService
}

func (s *server) GetAllLyrics(ctx context.Context, in *pb.Empty) (*pb.LyricsResponseList, error) {
	lyricsDtos := s.service.GetAllLyrics()
	result := &pb.LyricsResponseList{}

	for _, lyricsDto := range lyricsDtos {
		result.Lyrics = append(result.Lyrics, &pb.LyricsResponse{
			Id:        lyricsDto.Id,
			SongId:    lyricsDto.SongId,
			Content:   lyricsDto.Content,
			CreatedAt: lyricsDto.CreatedAt,
			UpdatedAt: lyricsDto.UpdatedAt,
		})
	}

	return result, nil
}

func (s *server) GetLyricsBySongId(ctx context.Context, in *pb.LyricsBySongIdRequest) (*pb.LyricsResponse, error) {
	lyricsDto := s.service.GetLyricsBySongId(in.SongId)
	if lyricsDto == nil {
		return nil, status.Errorf(codes.NotFound, "Lyrics not found")
	}

	return &pb.LyricsResponse{
		Id:        lyricsDto.Id,
		SongId:    lyricsDto.SongId,
		Content:   lyricsDto.Content,
		CreatedAt: lyricsDto.CreatedAt,
		UpdatedAt: lyricsDto.UpdatedAt,
	}, nil
}

func (s *server) CreateLyrics(ctx context.Context, in *pb.CreateLyricsRequest) (*pb.LyricsResponse, error) {
	result := s.service.AddLyrics(&dto.CreateLyricsDto{
		SongId:  in.SongId,
		Content: in.Content,
	})

	return &pb.LyricsResponse{
		Id:        result.Id,
		SongId:    result.SongId,
		Content:   result.Content,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	}, nil
}

func (s *server) UpdateLyrics(ctx context.Context, in *pb.UpdateLyricsRequest) (*pb.LyricsResponse, error) {
	result := s.service.UpdateLyrics(&dto.UpdateLyricsDto{
		Content: in.Content,
	}, in.SongId)
	if result == nil {
		return nil, status.Errorf(codes.NotFound, "Lyrics not found")
	}

	return &pb.LyricsResponse{
		Id:        result.Id,
		SongId:    result.SongId,
		Content:   result.Content,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	}, nil
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln("Error loading .env file")
	}

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	lyricsService := lyricsS.NewLyricsService(lyricsR.NewLyricsRepository())
	pb.RegisterLyricsServiceServer(grpcServer, &server{service: lyricsService})

	log.Printf("server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
