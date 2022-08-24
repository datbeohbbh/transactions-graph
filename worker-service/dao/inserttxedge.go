package dao

import "context"

func (dao *DAO) InsertTxEdge(ctx context.Context, txEdge *TxEdge) (any, error) {
	edgeColl := dao.GetCollection("edge")
	result, err := edgeColl.InsertOne(ctx, *txEdge)
	if err != nil {
		return nil, err
	}
	return result.InsertedID, nil
}
