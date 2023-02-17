package server

import (
	"errors"

	"context"

	"github.com/SebasEscobar/slack"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s SlackServer) Announce(ctx context.Context, r *slack.AnnounceRequest) (*emptypb.Empty, error) {
	return nil, errors.New("not yet implemented")
}
