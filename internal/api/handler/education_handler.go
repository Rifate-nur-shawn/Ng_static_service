package handler

import (
	"context"
	pb "matrimonial-service/internal/api/proto"
	"matrimonial-service/internal/domain/repository"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// EducationHandler struct
type EducationHandler struct {
	pb.UnimplementedEducationServiceServer
	eduRepo repository.EducationRepository
}

// NewEducationHandler creates a new gRPC handler for education
func NewEducationHandler(eduRepo repository.EducationRepository) *EducationHandler {
	return &EducationHandler{eduRepo: eduRepo}
}

// GetQualifications implementation
func (h *EducationHandler) GetQualifications(ctx context.Context, req *pb.GetQualificationsRequest) (*pb.GetQualificationsResponse, error) {
	quals, err := h.eduRepo.GetQualifications(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get qualifications: %v", err)
	}

	pbQuals := make([]*pb.Qualification, len(quals))
	for i, q := range quals {
		pbQuals[i] = &pb.Qualification{
			Id:          q.ID,
			Name:        q.Name,
			CommonName:  q.CommonName,
			LevelId:     q.LevelID,
			StreamId:    q.StreamID,
		}
	}

	return &pb.GetQualificationsResponse{Qualifications: pbQuals}, nil
}