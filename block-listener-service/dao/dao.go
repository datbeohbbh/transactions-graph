package dao

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (dao *DAO) GetAllVertex(ctx context.Context, noneParams *emptypb.Empty) (*Vertices, error) {
	return nil, nil
}

func (dao *DAO) GetVertexByAddress(ctx context.Context, vertexRequest *VertexRequest) (*Vertex, error) {
	return nil, nil
}
