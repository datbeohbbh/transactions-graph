package graphAlgo

import (
	"context"
	"fmt"
	"time"

	"github.com/datbeohbbh/go-utils/queue"
	grdao "github.com/datbeohbbh/transactions-graph/graph-service/dao"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Bfs struct {
	from              string
	depthLimit        uint64
	txCompletedBefore time.Time
	ctx               context.Context
	dao               grdao.IDAO
	curDepth          map[string]uint64
	vertexID          map[string]map[string]uint64
	link              map[uint64]map[uint64][]uint64
}

func NewBfs(ctx_ context.Context, dao_ grdao.IDAO, from_ string, depth_ uint64, txCompletedBefore_ time.Time) *Bfs {
	return &Bfs{
		ctx:               ctx_,
		dao:               dao_,
		from:              from_,
		depthLimit:        depth_,
		txCompletedBefore: txCompletedBefore_, // debug purpose
		curDepth:          make(map[string]uint64),
		vertexID:          make(map[string]map[string]uint64),
		link:              make(map[uint64]map[uint64][]uint64),
	}
}

func (bfs *Bfs) Execute() (*GraphRenderData, error) {
	from := bfs.from // common.HexToAddress(bfs.from).Hex()
	ctx := bfs.ctx
	dao := bfs.dao
	if exist, _ := bfs.dao.ExistedAddress(ctx, "vertex", from); !exist {
		return nil, fmt.Errorf("address %s does not exists", from)
	}
	result := GraphRenderData{}

	type Pair struct {
		Address string
		Type    string
	}

	fromVertex, err := dao.GetVertexByAddress(ctx, from)
	if err != nil {
		return nil, err
	}

	result.UpdateVertices(&Vertex{
		Depth:   1,
		Address: fromVertex.Address,
		Type:    fromVertex.Type,
	})

	queue := queue.New[Pair]()
	queue.Push(Pair{
		Address: from,
		Type:    fromVertex.Type,
	})

	bfs.curDepth[from] = 1
	if len(bfs.vertexID[fromVertex.Address]) == 0 {
		bfs.vertexID[fromVertex.Address] = make(map[string]uint64)
	}
	bfs.vertexID[fromVertex.Address][fromVertex.Type] = result.GetVerticesSize() - 1

	for !queue.Empty() {
		top := queue.Front()
		queue.Pop()

		vertex, err := dao.GetVertexByAddress(ctx, top.Address)
		if err != nil {
			return nil, err
		}

		txObjectIds := vertex.TxEdges
		for _, ids := range txObjectIds {
			tx, err := dao.GetTxByObjectID(ctx, ids)
			if err != nil {
				return nil, err
			}
			if tx.Direct == grdao.OUT {
				continue
			}

			var updTx bool = false
			if bfs.curDepth[tx.Address] == 0 &&
				bfs.curDepth[top.Address]+1 <= bfs.depthLimit &&
				tx.CreatedAt.Before(bfs.txCompletedBefore) {

				v, err := dao.GetVertexByAddress(ctx, tx.Address)
				if err != nil {
					result.UpdateTxs(&Transaction{
						TxHash:    tx.TxHash,
						Direction: "IN",
						CreateAt:  timestamppb.New(tx.CreatedAt),
					})

					result.UpdateVertices(&Vertex{
						Depth:   bfs.curDepth[top.Address] + 1,
						Address: tx.Address,
						Type:    "unknown type",
					})

					if len(bfs.vertexID[tx.Address]) == 0 {
						bfs.vertexID[tx.Address] = make(map[string]uint64)
					}
					bfs.vertexID[tx.Address]["unknown type"] = result.GetVerticesSize() - 1

					txID := result.GetTransactionsSize() - 1
					fromID := bfs.vertexID[top.Address][top.Type]
					toID := bfs.vertexID[tx.Address]["unknown type"]

					if len(bfs.link[fromID]) == 0 {
						bfs.link[fromID] = make(map[uint64][]uint64)
					}

					bfs.link[fromID][toID] = append(bfs.link[fromID][toID], txID)
					continue
				}

				bfs.curDepth[v.Address] = bfs.curDepth[top.Address] + 1
				result.UpdateVertices(&Vertex{
					Depth:   bfs.curDepth[v.Address],
					Address: v.Address,
					Type:    v.Type,
				})

				queue.Push(Pair{
					Address: v.Address,
					Type:    v.Type,
				})

				if len(bfs.vertexID[v.Address]) == 0 {
					bfs.vertexID[v.Address] = make(map[string]uint64)
				}
				bfs.vertexID[v.Address][v.Type] = result.GetVerticesSize() - 1
				updTx = true
			} else if bfs.curDepth[tx.Address] != 0 {
				updTx = true
			}

			if updTx {
				v, err := dao.GetVertexByAddress(ctx, tx.Address)
				if err != nil {
					return nil, err
				}

				result.UpdateTxs(&Transaction{
					TxHash:    tx.TxHash,
					Direction: "IN",
					CreateAt:  timestamppb.New(tx.CreatedAt),
				})
				txID := result.GetTransactionsSize() - 1
				fromID := bfs.vertexID[top.Address][top.Type]
				toID := bfs.vertexID[v.Address][v.Type]

				if len(bfs.link[fromID]) == 0 {
					bfs.link[fromID] = make(map[uint64][]uint64)
				}
				bfs.link[fromID][toID] = append(bfs.link[fromID][toID], txID)
			}
		}
	}

	for fromID, mp := range bfs.link {
		for toID, txIds := range mp {
			link := Link{
				From:      fromID,
				To:        toID,
				EdgesInfo: txIds,
			}
			result.UpdateLinks(&link)
		}
	}

	return &result, nil
}
