package grpc

import (
	"context"

	"github.com/inocentini/codepix-go/application/grpc/pb"
	"github.com/inocentini/codepix-go/application/usecase"
)

type PixgRPCService struct {
	PixUseCase usecase.PixUseCase
	pb.UnimplementedPixServiceServer
}

func (p *PixgRPCService) RegisterPixKey(ctx context.Context, in *pb.PixKeyRegistration) (*pb.PixKeyCreatedResult, error) {
	key, err := p.PixUseCase.RegisterKey(in.Key, in.Kind, in.AccountId)
	if err != nil {
		return &pb.PixKeyCreatedResult{
			Status: "not created",
			Error:  err.Error(),
		}, err
	}

	return &pb.PixKeyCreatedResult{
		Id:     key.ID,
		Status: "created",
	}, nil

}

func (p *PixgRPCService) Find(ctx context.Context, in *pb.PixKey) (*pb.PixKeyInfo, error) {
	pixKey, err := p.PixUseCase.FindKey(in.Key, in.Kind)
	if err != nil {
		return &pb.PixKeyInfo{}, nil
	}

	return &pb.PixKeyInfo{
		Id:   pixKey.ID,
		Kind: pixKey.Kind,
		Key:  pixKey.Key,
		Account: &pb.Account{
			AccountId:     pixKey.AccountID,
			AccountNumber: pixKey.Account.Number,
			BankId:        pixKey.Account.BankID,
			BankName:      pixKey.Account.Bank.Name,
			OwnerName:     pixKey.Account.OwnerName,
			CreatedAt:     pixKey.Account.CreatedAt.String(),
		},
		CreatedAt: pixKey.CreatedAt.String(),
	}, nil
}

func NewPixgRPCService(usecase usecase.PixUseCase) *PixgRPCService {
	return &PixgRPCService{
		PixUseCase: usecase,
	}
}
