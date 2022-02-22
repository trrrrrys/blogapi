package rest

type RestHandler interface {
	ContentHandler
}

func NewRestHandler(ch ContentHandler) RestHandler {
	return &restHandler{ch}
}

type restHandler struct {
	ContentHandler
}
