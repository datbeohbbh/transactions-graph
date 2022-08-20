package data

import (
	"context"
	"log"
	"time"
)

func (graphdb *GraphDB) UpdateDB(pctx context.Context, vertex *Vertex, txEdge *TxEdge, events []*Event) error {
	ctx, cancel := context.WithTimeout(pctx, 3*time.Second)
	defer cancel()
	upd := func() <-chan error {
		errc := make(chan error)
		go func() {
			defer close(errc)
			found, err := graphdb.ExistedVertex(ctx, vertex.Address)
			if err != nil {
				errc <- err
			}
			if found {
				err = graphdb.UpdateVertex(ctx, vertex, txEdge, events)
				if err != nil {
					errc <- err
				}
			} else {
				err = graphdb.InsertVertex(ctx, vertex, txEdge, events)
				if err != nil {
					errc <- err
				}
			}
			errc <- nil
		}()
		return errc
	}

	select {
	case err := <-upd():
		log.Printf("vertex added/updated: %s - msg: %v", vertex.Address, err)
		return err
	case <-ctx.Done():
		log.Printf("Fail on on updateDB: %v", ctx.Err())
		return ctx.Err()
	}
}
