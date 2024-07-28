package main

import (
	fmt "fmt"
	"github.com/monimesl/openapi-go/openapi3"
	"log"
	"net/http"
	"time"
)

func main() {
	reflector := openapi3.Reflector{}
	reflector.Spec = &openapi3.Spec{Openapi: "3.0.3"}
	reflector.Spec.Info.
		WithTitle("Things API").
		WithVersion("1.2.3").
		WithDescription("Put something here")

	type req struct {
		ID string `path:"id" example:"XXX-XXXXX"`
	}

	type resp struct {
		ID     string `json:"id" example:"XXX-XXXXX"`
		Amount uint   `json:"amount"`
		Items  []struct {
			Count uint   `json:"count"`
			Name  string `json:"name" description:"item name"`
		} `json:"items"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	putOp, err := reflector.NewOperationContext(http.MethodPut, "/things/{id}")
	handleError(err)

	putOp.AddReqStructure(new(req))
	putOp.AddRespStructure(new(resp))

	reflector.AddOperation(putOp)

	getOp, err := reflector.NewOperationContext(http.MethodGet, "/things/{id}")
	handleError(err)

	getOp.AddReqStructure(new(req))
	//getOp.AddRespStructure(new(resp), func(cu *openapi.ContentUnit) { cu.HTTPStatus = http.StatusOK })

	reflector.AddOperation(getOp)

	schema, err := reflector.Spec.MarshalYAML()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(schema))
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
