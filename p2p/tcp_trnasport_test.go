package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTCPTransport(t *testing.T){
	listenAddr := "400"
	tr := InitCPTransport(listenAddr)
	assert.Equal(t,tr.listenAddr,listenAddr)
}