package graph

import (
	"context"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func (graph *GraphData) GetVertexByAccountType(ctx context.Context, query *Query) (*Vertices, error) {
	vertices, err := graph.dao.GetVertexByAccountType(ctx, query.GetAccountType())
	if err != nil {
		return nil, err
	}

	var ret Vertices
	for _, vertex := range vertices {
		elem := Vertex{
			Address:   vertex.Address,
			Type:      vertex.Type,
			TxEdges:   vertex.TxEdges,
			CreatedAt: timestamppb.New(vertex.CreatedAt),
			UpdatedAt: timestamppb.New(vertex.UpdatedAt),
		}
		ret.Vertices = append(ret.Vertices, &elem)
	}
	return &ret, nil
}
