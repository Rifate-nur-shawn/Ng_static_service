package handler

import (
	"context"
	pb "matrimonial-service/internal/api/proto" // Your generated proto package
	"matrimonial-service/internal/domain/model"
	"matrimonial-service/internal/domain/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// LocationHandler struct
type LocationHandler struct {
	pb.UnimplementedLocationServiceServer
	locRepo repository.LocationRepository
}

// NewLocationHandler creates a new gRPC handler for locations
func NewLocationHandler(locRepo repository.LocationRepository) *LocationHandler {
	return &LocationHandler{locRepo: locRepo}
}

// GetDivisions implementation
func (h *LocationHandler) GetDivisions(ctx context.Context, req *pb.GetDivisionsRequest) (*pb.GetDivisionsResponse, error) {
	divisions, err := h.locRepo.GetDivisions(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get divisions: %v", err)
	}

	pbDivisions := make([]*pb.Division, len(divisions))
	for i, d := range divisions {
		pbDivisions[i] = &pb.Division{
			Id:     d.ID,
			NameEn: d.NameEn,
			NameBn: d.NameBn,
		}
	}

	return &pb.GetDivisionsResponse{Divisions: pbDivisions}, nil
}

// GetDistrictsByDivision implementation
func (h *LocationHandler) GetDistrictsByDivision(ctx context.Context, req *pb.GetDistrictsByDivisionRequest) (*pb.GetDistrictsByDivisionResponse, error) {
	districts, err := h.locRepo.GetDistrictsByDivision(ctx, req.GetDivisionId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get districts: %v", err)
	}

	pbDistricts := make([]*pb.District, len(districts))
	for i, d := range districts {
		pbDistricts[i] = &pb.District{
			Id:         d.ID,
			DivisionId: d.DivisionID,
			NameEn:     d.NameEn,
			NameBn:     d.NameBn,
		}
	}

	return &pb.GetDistrictsByDivisionResponse{Districts: pbDistricts}, nil
}

// GetUpazilasByDistrict implementation
func (h *LocationHandler) GetUpazilasByDistrict(ctx context.Context, req *pb.GetUpazilasByDistrictRequest) (*pb.GetUpazilasByDistrictResponse, error) {
	var upazilas []*model.Upazila
	var err error
	
	// If search parameter is provided, use search
	if req.GetSearch() != "" {
		upazilas, err = h.locRepo.SearchUpazilasByDistrict(ctx, req.GetDistrictId(), req.GetSearch())
	} else {
		upazilas, err = h.locRepo.GetUpazilasByDistrict(ctx, req.GetDistrictId())
	}
	
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get upazilas: %v", err)
	}

	pbUpazilas := make([]*pb.Upazila, len(upazilas))
	for i, u := range upazilas {
		pbUpazilas[i] = &pb.Upazila{
			Id:         u.ID,
			DistrictId: u.DistrictID,
			NameEn:     u.NameEn,
			NameBn:     u.NameBn,
		}
	}

	return &pb.GetUpazilasByDistrictResponse{Upazilas: pbUpazilas}, nil
}

// GetUpazilaById implementation
func (h *LocationHandler) GetUpazilaById(ctx context.Context, req *pb.GetUpazilaByIdRequest) (*pb.GetUpazilaByIdResponse, error) {
	upazila, err := h.locRepo.GetUpazilaById(ctx, req.GetUpazilaId())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "upazila not found: %v", err)
	}

	pbUpazila := &pb.Upazila{
		Id:         upazila.ID,
		DistrictId: upazila.DistrictID,
		NameEn:     upazila.NameEn,
		NameBn:     upazila.NameBn,
	}

	return &pb.GetUpazilaByIdResponse{Upazila: pbUpazila}, nil
}