package graphqlschemas

import (
	"database/sql"
	graphqlobjects "igrwijaya-go-template/internal/application/graphql_objects"
	"igrwijaya-go-template/internal/domain/common"
	"igrwijaya-go-template/internal/domain/todo"
	"time"

	"github.com/graphql-go/graphql"
)

type TodoGraphql interface {
	Query() *graphql.Object
	Mutation() *graphql.Object
}

type todoGraphql struct {
	todoRepo todo.TodoRepository
}

func (q *todoGraphql) Mutation() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			/* Create new product item
			http://localhost:8080/product?query=mutation+_{create(name:"Inca Kola",info:"Inca Kola is a soft drink that was created in Peru in 1935 by British immigrant Joseph Robinson Lindley using lemon verbena (wiki)",price:1.99){id,name,info,price}}
			*/
			"create": &graphql.Field{
				Type:        graphqlobjects.TodoObjectGraph(),
				Description: "Create new Todo",
				Args: graphql.FieldConfigArgument{
					"title": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"description": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					todo := todo.Todo{
						Title:       params.Args["title"].(string),
						Description: params.Args["description"].(string),
						AuditableEntity: common.AuditableEntity{
							CreatedBy: 0,
							CreatedAt: time.Now(),
							UpdatedBy: 0,
							UpdatedAt: sql.NullTime{
								Time:  time.Now(),
								Valid: true,
							},
						},
					}

					createdTodo, errCreate := q.todoRepo.Create(todo)

					if errCreate != nil {
						return nil, errCreate
					}

					return createdTodo, nil
				},
			},

			/* Update product by id
			   http://localhost:8080/product?query=mutation+_{update(id:1,price:3.95){id,name,info,price}}
			*/
			"update": &graphql.Field{
				Type:        graphqlobjects.TodoObjectGraph(),
				Description: "Update Todo by Id",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"title": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"description": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					id, _ := params.Args["id"].(uint)
					title, titleOk := params.Args["title"].(string)
					description, descOk := params.Args["description"].(string)

					todo, err := q.todoRepo.Read(id)

					if err != nil {
						return nil, err
					}

					if titleOk {
						todo.Title = title
					}

					if descOk {
						todo.Description = description
					}

					updatedTodo, err := q.todoRepo.Update(todo)

					return updatedTodo, nil
				},
			},

			/* Delete product by id
			   http://localhost:8080/product?query=mutation+_{delete(id:1){id,name,info,price}}
			*/
			"delete": &graphql.Field{
				Type:        graphqlobjects.TodoObjectGraph(),
				Description: "Delete Todo by Id",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					id, _ := params.Args["id"].(uint)

					todo, err := q.todoRepo.Read(id)

					if err != nil {
						return nil, err
					}

					errDel := q.todoRepo.Delete(id)

					if errDel != nil {
						return nil, errDel
					}

					return todo, nil
				},
			},
		},
	})
}

func (q *todoGraphql) Query() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "TodoQuery",
			Fields: graphql.Fields{
				/* Get (read) single product by id
				   http://localhost:8080/product?query={product(id:1){name,info,price}}
				*/
				"todo": &graphql.Field{
					Type:        graphqlobjects.TodoObjectGraph(),
					Description: "Get Todo by Id",
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{
							Type: graphql.Int,
						},
					},
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						id, ok := p.Args["id"].(uint)
						if ok {
							todo, err := q.todoRepo.Read(id)
							if err != nil {
								return nil, err
							}

							return todo, nil
						}

						return nil, nil
					},
				},
				/* Get (read) product list
				   http://localhost:8080/product?query={list{id,name,info,price}}
				*/
				"list": &graphql.Field{
					Type:        graphql.NewList(graphqlobjects.TodoObjectGraph()),
					Description: "Get Todo list",
					Args: graphql.FieldConfigArgument{
						"page": &graphql.ArgumentConfig{
							Type: graphql.Int,
						},
						"limit": &graphql.ArgumentConfig{
							Type: graphql.Int,
						},
					},
					Resolve: func(params graphql.ResolveParams) (interface{}, error) {
						page, ok := params.Args["page"].(int)
						if !ok {
							return nil, nil
						}

						limit, ok := params.Args["limit"].(int)
						if !ok {
							return nil, nil
						}

						todos, err := q.todoRepo.GetAll(page, limit)
						if err != nil {
							return nil, err
						}

						return todos, nil
					},
				},
			},
		})
}

func NewTodoGraphql(todoRepo todo.TodoRepository) TodoGraphql {
	return &todoGraphql{
		todoRepo: todoRepo,
	}
}
