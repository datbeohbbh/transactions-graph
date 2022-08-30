package graphAlgo

func (gr *GraphRenderData) UpdateVertices(v *Vertex) {
	v.NodeID = uint64(len(gr.GetVertices()))
	gr.Vertices = append(gr.Vertices, v)
}

func (gr *GraphRenderData) UpdateTxs(tx *Transaction) {
	tx.TxID = uint64(len(gr.GetTransactions()))
	gr.Transactions = append(gr.Transactions, tx)
}

func (gr *GraphRenderData) UpdateLinks(link *Link) {
	gr.Links = append(gr.Links, link)
}

func (gr *GraphRenderData) GetVerticesSize() uint64 {
	return uint64(len(gr.GetVertices()))
}

func (gr *GraphRenderData) GetTransactionsSize() uint64 {
	return uint64(len(gr.GetTransactions()))
}
