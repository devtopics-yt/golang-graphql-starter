package store

import (
	"context"
	"devtopics-gql/graph/model"
	"net/http"
)

type Store struct {
	Todos []*model.Todo
}

func NewStore() *Store {
	todos := make([]*model.Todo, 0)

	return &Store{
		Todos: todos,
	}
}

// Write TODO - helper func that updates todos state
func (s *Store) AddTodo(t *model.NewTodo) error {
	s.Todos = append(s.Todos, &model.Todo{
		ID:   "1",
		Text: t.Text,
		Done: false,
		User: &model.User{
			ID:   t.UserID,
			Name: "devtopics",
		},
	})

	return nil
}

type StoreKeyType string

var StoreKey StoreKeyType = "STORE"

// WithStore middle - inject store into context
func WithStore(store *Store, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get context and add our store to context
		reqWithStore := r.WithContext(context.WithValue(r.Context(), StoreKey, store))

		next.ServeHTTP(w, reqWithStore)
	})
}

// GetStoreFromContext - retrieves store from request context
func GetStoreFromContext(ctx context.Context) *Store {
	store, ok := ctx.Value(StoreKey).(*Store)
	if !ok {
		panic("couldn't find the store")
	}

	return store
}
