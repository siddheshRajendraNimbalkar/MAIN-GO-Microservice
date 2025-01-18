package main

import "github.com/99designs/gqlgen/graphql"

type Server struct {
	// accountClient *account.Client
	// productClient *product.Client
	// orderClient   *order.Client
}

func NewGraphQLServer(accountURL, productURL, orderURL string) (*Server, error) {
	// accountClient, err := account.NewClient(accountURL)

	// if err != nil {
	// 	return nil, err
	// }
	// productClient, err := product.NewClient(productURL)
	// if err != nil {
	// 	accountClient.Close()
	// 	return nil, err
	// }

	// orderClient, err := order.NewClient(orderURL)
	// if err != nil {
	// 	accountClient.Close()
	// 	productClient.Close()
	// 	return nil, err
	// }

	return &Server{
		// accountClient,
		// productClient,
		// orderClient,
	}, nil
}

// func (s *Server) Mutation() MutationResolver {
// 	return &mutationResolver{
// 		server: s,
// 	}
// }

// func (s *Server) Query() QuaryResolver {
// 	return &quaryResolver{
// 		server: s,
// 	}
// }

// func (s *Server) Account() AccountResolver {
// 	return &accountResolver{
// 		server: s,
// 	}
// }

func (s *Server) ToExecutableSchema() graphql.ExecutableSchema {
	return NewExecutableSchema(Config{
		Resolvers: s,
	})
}
