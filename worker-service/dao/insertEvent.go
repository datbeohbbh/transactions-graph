package dao

import "context"

func (dao *DAO) InsertEvent(ctx context.Context, events []*Event) ([]any, error) {
	eventColl := dao.GetCollection("event")
	eventConv := []any{}
	for _, pe := range events {
		eventConv = append(eventConv, *pe)
	}
	if len(eventConv) == 0 {
		return []any{}, nil
	}
	result, err := eventColl.InsertMany(ctx, eventConv)
	if err != nil {
		return []any{}, err
	}
	return result.InsertedIDs, nil
}
