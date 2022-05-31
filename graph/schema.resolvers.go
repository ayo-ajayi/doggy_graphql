package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	//"fmt"

	"github.com/ayo-ajayi/gqlgen-todos/database"
	"github.com/ayo-ajayi/gqlgen-todos/graph/generated"
	"github.com/ayo-ajayi/gqlgen-todos/graph/model"
)

var db = database.Connect()
func (r *mutationResolver) CreateDog(ctx context.Context, input *model.NewDog) (*model.Dog, error) {
	//panic(fmt.Errorf("not implemented"))
	return db.Save(input), nil
}

func (r *queryResolver) Dog(ctx context.Context, id string) (*model.Dog, error) {
	//panic(fmt.Errorf("not implemented"))
	return db.FindByID(id), nil
}

func (r *queryResolver) Dogs(ctx context.Context) ([]*model.Dog, error) {
	//panic(fmt.Errorf("not implemented"))
	return db.All(), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

