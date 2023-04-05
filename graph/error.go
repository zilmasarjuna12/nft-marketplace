package graph

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

const (
	INVALID_INPUT         = "BAD_REQUEST"
	INTERNAL_SERVER_ERROR = "INTERNAL_SERVER_ERROR"
	ITEM_EMPTY            = "ITEM_EMPTY"
	USER_NOT_NOUND        = "USER_NOT_FOUND"
	NOT_FOUND             = "NOT_FOUND"
	NOT_ACCEPTABLE        = "NOT_ACCEPTABLE"
)

func AddInputError(ctx context.Context, invalidInput interface{}) {
	graphql.AddError(ctx, &gqlerror.Error{
		Path:    graphql.GetPath(ctx),
		Message: "input are invalid",
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
	case USER_NOT_NOUND:
		graphql.AddError(ctx, &gqlerror.Error{
			Path:    graphql.GetPath(ctx),
			Message: "User Not Found",
			Extensions: map[string]interface{}{
				"code": USER_NOT_NOUND,
			},
		})
	case ITEM_EMPTY:
		graphql.AddError(ctx, &gqlerror.Error{
			Path:    graphql.GetPath(ctx),
			Message: "Item Is't Available",
			Extensions: map[string]interface{}{
				"code": ITEM_EMPTY,
			},
		})
	case NOT_ACCEPTABLE:
		graphql.AddError(ctx, &gqlerror.Error{
			Path:    graphql.GetPath(ctx),
			Message: "Not Acceptable",
			Extensions: map[string]interface{}{
				"code": NOT_ACCEPTABLE,
			},
		})
	default:
		graphql.AddError(ctx, &gqlerror.Error{
			Path:    graphql.GetPath(ctx),
			Message: "Internal Server Error",
			Extensions: map[string]interface{}{
				"code": INTERNAL_SERVER_ERROR,
			},
		})
	}

}
