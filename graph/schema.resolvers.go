package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"graphql/graph/generated"
	"graphql/graph/model"
	"time"
)

// CreateMovie is the resolver for the createMovie field.
func (r *mutationResolver) CreateMovie(ctx context.Context, input model.NewMovie) (*model.Movie, error) {

	movie := model.Movie{
		Title: input.Title,
		URL:   input.URL,
	}

	docId := fmt.Sprintf("id_%d", time.Now().Unix())
	movie.ID = docId
	movie.ReleaseDate = time.Now().Format("2006-01-02")
	mres := r.DB.InsertIntoBucket(docId, "bagnbrand.graphql.movie", &movie)
	if mres.Status != "success" {
		return nil, fmt.Errorf("%+v", mres.Errors)
	}
	// _, err := r.DB.Model(&movie).Insert()
	// if err != nil {
	// 	return nil, fmt.Errorf("error inserting new movie: %v", err)
	// }
	return &movie, nil
}

// Movies is the resolver for the movies field.
func (r *queryResolver) Movies(ctx context.Context) ([]*model.Movie, error) {

	var movies []*model.Movie

	//docId := fmt.Sprintf("101_%v", movie.ID)
	mres := r.DB.Query("SELECT id,title,url,releaseDate FROM bagnbrand.graphql.movie;")
	if mres.Status != "success" {

		fmt.Println("ERR", mres.Errors, fmt.Sprintf("%+v", mres.Errors))
		return nil, fmt.Errorf("%v", "something went wrong")
	}
	rows := mres.GetRows()
	for _, row := range rows {
		var movie = &model.Movie{}
		movie.ID = fmt.Sprint(row["id"])
		movie.ReleaseDate = fmt.Sprint(row["releaseDate"])
		movie.Title = fmt.Sprint(row["title"])
		movie.URL = fmt.Sprint(row["url"])
		movies = append(movies, movie)
	}
	// err := r.DB.Model(&movies).Select()
	// if err != nil {
	// 	return nil, err
	// }
	return movies, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
