package handler

import (
	"encoding/json"

	"../gql_schema"
	"../models"
	"../resolver"
	"../src/routes"
	"github.com/gin-gonic/gin"
	graphql "github.com/graph-gophers/graphql-go"
)

type Handler struct {
	Schema *graphql.Schema
}

func GraphqlHandler() gin.HandlerFunc {

	db := models.GetDB()

	dbResolver := resolver.DBStruct{
		Db: db,
	}

	resolver := resolver.Resolver{
		DBStruct: &dbResolver,
	}

	schemaString, err := gql_schema.String()

	if err != nil {
		return func(c *gin.Context) {
			routes.SendResponse(c, 500, err.Error())
			return
		}
	}

	h := Handler{
		Schema: graphql.MustParseSchema(schemaString, &resolver),
	}

	return func(c *gin.Context) {
		h.ServeHTTP(c)
	}
}

func (h *Handler) ServeHTTP(c *gin.Context) {
	var params struct {
		Query         string                 `json:"query"`
		OperationName string                 `json:"operationName"`
		Variables     map[string]interface{} `json:"variables"`
	}

	if err := json.NewDecoder(c.Request.Body).Decode(&params); err != nil {
		routes.SendResponse(c, 500, err.Error())
		return
	}

	response := h.Schema.Exec(c.Request.Context(), params.Query, params.OperationName, params.Variables)
	responseJSON, err := json.Marshal(response)
	if err != nil {
		routes.SendResponse(c, 500, err.Error())
		return
	}

	routes.SendResponse(c, 200, string(responseJSON))
	return
}
