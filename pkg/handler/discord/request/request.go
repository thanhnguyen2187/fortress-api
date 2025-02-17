package request

import (
	"github.com/dwarvesf/fortress-api/pkg/handler/discord/errs"
)

type BraineryReportInput struct {
	View      string `json:"view" binding:"required"`
	ChannelID string `json:"channelID" binding:"required"`
}

func (input BraineryReportInput) Validate() error {
	if len(input.View) == 0 {
		return errs.ErrEmptyReportView
	}

	if len(input.ChannelID) == 0 {
		return errs.ErrEmptyChannelID
	}
	return nil
}
