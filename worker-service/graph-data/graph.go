package graph

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"go.mongodb.org/mongo-driver/bson"
)

func (gd *GraphData) GetAllVertex(ctx context.Context, projection *ExcludeFields) (*Vertices, error) {
	excludeFields := []any{}
	for _, ex := range projection.ExcludeField {
		excludeFields = append(excludeFields, bson.D{{ex.Field, ex.Opt}})
	}
	result, err := gd.graphDB.GetAllVertex(ctx, excludeFields)
	if err != nil {
		return nil, err
	}

	vertices := Vertices{}
	for _, b := range result {
		vertex := Vertex{}
		err = json.Unmarshal(b, &vertex)
		if err != nil {
			return nil, err
		}
		vertices.Vertices = append(vertices.Vertices, &vertex)
	}
	return &vertices, nil
}

func (gd *GraphData) GetVertexByAddress(ctx context.Context, vertexRequest *VertexRequest) (*Vertex, error) {
	address := common.HexToAddress(vertexRequest.Address).Hex()
	projection := vertexRequest.Exclude
	excludeFields := []any{}
	if projection != nil {
		for _, ex := range projection.ExcludeField {
			excludeFields = append(excludeFields, bson.D{{ex.Field, ex.Opt}})
		}
	}

	result, err := gd.graphDB.GetVertexByAddress(ctx, address, excludeFields)
	if err != nil {
		return nil, fmt.Errorf("failed on get vertex of address %s: %v", address, err)
	}

	vertex := Vertex{}
	err = json.Unmarshal(result, &vertex)
	if err != nil {
		return nil, fmt.Errorf("failed on unmarshal vertex result: %v", err)
	}
	return &vertex, nil
}
