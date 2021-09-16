package keeper

import (
	"context"
	"encoding/base64"
	"github.com/cosmonaut/planet/x/blog/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/cosmos-sdk/x/ibc/core/02-client/types"
)

func (k msgServer) SendIbcPost(goCtx context.Context, msg *types.MsgSendIbcPost) (*types.MsgSendIbcPostResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: logic before transmitting the packet

	// Construct the packet
	var packet types.IbcPostPacketData

	packet.Title = msg.Title
	hello := msg.Content
	strbytes := []byte(hello)
	encoded := base64.StdEncoding.EncodeToString(strbytes)
	packet.Content = encoded
	packet.Creator = msg.Sender
	// Transmit the packet
	err := k.TransmitIbcPostPacket(
		ctx,
		packet,
		msg.Port,
		msg.ChannelID,
		clienttypes.ZeroHeight(),
		msg.TimeoutTimestamp,
	)
	if err != nil {
		return nil, err
	}

	return &types.MsgSendIbcPostResponse{}, nil
}
