package server

import (
	"testing"

	"context"
	"github.com/stretchr/testify/assert"
	"github.com/tessthedog/slack"
)

func TestAnnounce(t *testing.T) {
	ctx := context.Background()
	req := &slack.AnnounceRequest{}

	res, err := cli.Announce(ctx, req)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}
