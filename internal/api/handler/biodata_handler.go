package handler

import (
	"context"
	pb "matrimonial-service/internal/api/proto"
	"matrimonial-service/internal/domain/model"
	"matrimonial-service/internal/domain/repository"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// StaticDataHandler struct
type StaticDataHandler struct {
	pb.UnimplementedStaticDataServiceServer
	staticRepo repository.StaticDataRepository
}

// NewStaticDataHandler creates a new gRPC handler for static data
func NewStaticDataHandler(staticRepo repository.StaticDataRepository) *StaticDataHandler {
	return &StaticDataHandler{staticRepo: staticRepo}
}

// Helper to convert domain model to proto
func toProtoOptions(options []*model.StaticOption) []*pb.StaticOption {
	pbOptions := make([]*pb.StaticOption, len(options))
	for i, opt := range options {
		pbOptions[i] = &pb.StaticOption{
			Id:   opt.ID,
			Name: opt.Name,
		}
	}
	return pbOptions
}

func (h *StaticDataHandler) GetReligions(ctx context.Context, req *pb.GetReligionsRequest) (*pb.GetReligionsResponse, error) {
	religions, err := h.staticRepo.GetReligions(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get religions: %v", err)
	}
	return &pb.GetReligionsResponse{Religions: toProtoOptions(religions)}, nil
}

func (h *StaticDataHandler) GetProfessions(ctx context.Context, req *pb.GetProfessionsRequest) (*pb.GetProfessionsResponse, error) {
	professions, err := h.staticRepo.GetProfessions(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get professions: %v", err)
	}
	return &pb.GetProfessionsResponse{Professions: toProtoOptions(professions)}, nil
}

func (h *StaticDataHandler) GetMaritalStatuses(ctx context.Context, req *pb.GetMaritalStatusesRequest) (*pb.GetMaritalStatusesResponse, error) {
	statuses, err := h.staticRepo.GetMaritalStatuses(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get marital statuses: %v", err)
	}
	return &pb.GetMaritalStatusesResponse{MaritalStatuses: toProtoOptions(statuses)}, nil
}