package dao

import (
	"context"
	"log"
	"time"
)

func (dao *DAO) UpdateDB(pctx context.Context, vertex *Vertex, txEdge *TxEdge, events []*Event) error {
	ctx, cancel := context.WithTimeout(pctx, 3*time.Second)
	defer cancel()
	upd := func() <-chan error {
		errc := make(chan error)
		go func() {
			defer close(errc)
			found, err := dao.ExistedAddress(ctx, "vertex", vertex.Address)
			if err != nil {
				errc <- err
			}
			if found {
				err = dao.UpdateVertex(ctx, vertex, txEdge, events)
				if err != nil {
					errc <- err
				}
			} else {
				err = dao.InsertVertex(ctx, vertex, txEdge, events)
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
		//log.Printf("vertex added/updated: %s - msg: %v", vertex.Address, err)
		return err
	case <-ctx.Done():
		log.Printf("Fail on on updateDB: %v", ctx.Err())
		return ctx.Err()
	}
}
