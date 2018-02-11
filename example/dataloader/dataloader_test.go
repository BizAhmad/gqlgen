package dataloader

import (
	"net/http/httptest"
	"testing"

	"github.com/vektah/graphql-go"
	"github.com/vektah/graphql-go/client"
	"github.com/vektah/graphql-go/introspection"
)

func TestTodo(t *testing.T) {
	srv := httptest.NewServer(LoaderMiddleware(graphql.Handler(NewExecutor(&Resolver{}))))
	c := client.New(srv.URL)

	t.Run("create a new todo", func(t *testing.T) {
		c.MustPost(`{
		  customers {
			name
			address {
			  street
			  
			}
			orders {
			  id
              amount
			  items {
				name
			  }
			}
		  }
		}`, nil)
	})

	t.Run("introspection", func(t *testing.T) {
		// Make sure we can run the graphiql introspection query without errors
		c.MustPost(introspection.Query, nil)
	})
}
