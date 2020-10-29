package impl

import (
	"testing"

	"github.com/go-openapi/swag"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go.timothygu.me/downtomeet/restapi/operations"
)

func TestGetHello(t *testing.T) {
	params := operations.NewGetHelloParams()

	res := GetHello(params)
	require.IsType(t, (*operations.GetHelloOK)(nil), res)
	assert.Equal(t, res.(*operations.GetHelloOK).Payload.Hello, "world 1")
}

func TestGetHelloError(t *testing.T) {
	params := operations.NewGetHelloParams()
	params.ID = swag.String("error")

	res := GetHello(params)
	require.IsType(t, (*operations.GetHelloDefault)(nil), res)
	assert.Equal(t, res.(*operations.GetHelloDefault).Payload.Message, "ID is \"error\"")
}
