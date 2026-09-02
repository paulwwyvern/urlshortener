package selector

import (
	"context"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/selector"
)

func UnprotectedMethodMatcher(unprotected map[string]bool) selector.Matcher {

	return selector.MatchFunc(func(_ context.Context, c interceptors.CallMeta) bool {
		return !unprotected[c.FullMethod()]
	})
}
