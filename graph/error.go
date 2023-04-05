package graph

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

const (
	INVALID_INPUT = "BAD_REQUEST"
	NOT_FOUND     = "NOT_FOUND"
)

func AddInputError(ctx context.Context, invalidInput map[string]interface{}) {
	graphql.AddError(ctx, &gqlerror.Error{
		Path:    graphql.GetPath(ctx),
		Message: "Multiple input are invalid",
		Extensions: map[string]interface{}{
			"code":         "INVALID_INPUT",
			"invalidInput": invalidInput,
		},
	})
}

func AddError(ctx context.Context, code string) {
	switch code {
	case NOT_FOUND:
		graphql.AddError(ctx, &gqlerror.Error{
			Path:    graphql.GetPath(ctx),
			Message: "Data Not Found",
			Extensions: map[string]interface{}{
				"code": NOT_FOUND,
			},
		})
	}
}
