package service


import (
	"context"
	grpc_svn_protos "github.com/fatelei/grpc-svn/internal"
	"github.com/fatelei/grpc-svn/internal/common"
	controller2 "github.com/fatelei/grpc-svn/pkg/controller/impl"
)

type GrpcSvnServiceServerImpl struct {
	grpc_svn_protos.UnimplementedGrpcSvnServiceServer
	ctl *controller2.GitSvnImpl
}

func NewGrpcSvnServiceServer() *GrpcSvnServiceServerImpl {
	ctl := controller2.NewGivSvnController()
	return &GrpcSvnServiceServerImpl{ctl: ctl}
}

var _ grpc_svn_protos.GrpcSvnServiceServer = (*GrpcSvnServiceServerImpl)(nil)

func (p *GrpcSvnServiceServerImpl) UpdateRepo(
	ctx context.Context, param *grpc_svn_protos.UpdateRepoRequest) (*grpc_svn_protos.UpdateRepoResponse, error) {
	err := p.ctl.Update(ctx, param.SvnUrl)
	if err == nil {
		res := &grpc_svn_protos.UpdateRepoResponse{}
		res.Status = &common.ResponseStatus{
			Code:    common.ResponseStatus_SUCCESS,
			Message: "",
		}
		return res, nil
	}
	return nil, err
}

func (p *GrpcSvnServiceServerImpl) Clone(ctx context.Context, param *grpc_svn_protos.CloneRepoRequest) (
	*grpc_svn_protos.CloneResponse, error) {
	err := p.ctl.Clone(ctx, param.SvnUrl, param.MailSuffix)
	if err == nil {
		response := &grpc_svn_protos.CloneResponse{}
		response.Status = &common.ResponseStatus{
			Code:    common.ResponseStatus_SUCCESS,
			Message: "",
		}
		return response, nil
	}
	return nil, err
}
