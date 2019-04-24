package dispatch

import (
	"github.com/sh4nnongoh/scheduler/abc"
	"github.com/sh4nnongoh/scheduler/dispatchWindow"
	"github.com/sh4nnongoh/scheduler/qwe"
)

type Dispatch struct {
	DispatchID       DispatchID
	DispatchWindowID dispatchWindow.DispatchWindowID
	ResourceAbcID    abc.ResourceAbcID
	ResourceQweID    qwe.ResourceQweID
	DateTime         DateTime
	Status           Status
}

func New(id DispatchID, windowID dispatchWindow.DispatchWindowID, abcID abc.ResourceAbcID, qweID qwe.ResourceQweID) *Dispatch {
	return &Dispatch{
		DispatchID:       id,
		DispatchWindowID: windowID,
		ResourceAbcID:    abcID,
		ResourceQweID:    qweID,
		DateTime:         NewDateTime(),
		Status:           "PENDING",
	}
}

type DispatchRepository interface {
	New([]Dispatch) error
	Get([]DispatchID) ([]Dispatch, error)
}
