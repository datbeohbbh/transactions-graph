package data

import "context"

func (graphdb *GraphDB) InsertTxEdge(ctx context.Context, txEdge *TxEdge) (any, error) {
	edgeColl := graphdb.GetCollection("edge")
	result, err := edgeColl.InsertOne(ctx, *txEdge)
	if err != nil {
		return nil, err
	}
	return result.InsertedID, nil
}
