package middleware

import "github.com/julienschmidt/httprouter"

type Middleware func(httprouter.Handle) httprouter.Handle

func BuildChain(f httprouter.Handle, m ...Middleware) httprouter.Handle {
	if len(m) == 0 {
		return f
	}
	return f
	// return m[0](BuildChain(f, m[1:cap(m)]...))
}
