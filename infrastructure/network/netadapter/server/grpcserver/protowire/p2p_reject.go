package protowire

import (
	"github.com/harbi-network/harbid/app/appmessage"
	"github.com/pkg/errors"
)

func (x *HarbidMessage_Reject) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "HarbidMessage_Reject is nil")
	}
	return x.Reject.toAppMessage()
}

func (x *RejectMessage) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "RejectMessage is nil")
	}
	return &appmessage.MsgReject{
		Reason: x.Reason,
	}, nil
}

func (x *HarbidMessage_Reject) fromAppMessage(msgReject *appmessage.MsgReject) error {
	x.Reject = &RejectMessage{
		Reason: msgReject.Reason,
	}
	return nil
}
