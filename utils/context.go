package utils

import "context"

func FirstContextOrBackground(ctx []context.Context) context.Context {
	if ctx == nil || len(ctx) == 0 || ctx[0] == nil {
		return context.Background()
	}
	return ctx[0]
}
