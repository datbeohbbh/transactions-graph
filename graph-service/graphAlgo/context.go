package graphAlgo

type GraphRenderContext struct {
	algo GraphAlgo
}

func New() *GraphRenderContext {
	return &GraphRenderContext{}
}

func (ctx *GraphRenderContext) SetGraphAlgo(algo_ GraphAlgo) {
	ctx.algo = algo_
}

func (ctx *GraphRenderContext) Execute() (*GraphRenderData, error) {
	return ctx.algo.Execute()
}
