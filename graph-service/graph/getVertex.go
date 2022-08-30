package graph

import (
	"context"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func (graph *GraphData) GetVertexByAddress(ctx context.Context, request *Query) (*Vertex, error) {
	addr := request.GetAddress()
	vertex, err := graph.dao.GetVertexByAddress(ctx, addr)
	if err != nil {
		return nil, err
	}

	v := Vertex{
		Address:   vertex.Address,
		Type:      vertex.Type,
		TxEdges:   vertex.TxEdges,
		CreatedAt: timestamppb.New(vertex.CreatedAt),
		UpdatedAt: timestamppb.New(vertex.UpdatedAt),
	}

	return &v, nil
}
