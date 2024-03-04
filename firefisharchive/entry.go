package firefisharchive

import (
	"sourcecode.social/reiver/go-opt"
)

// Entry represents a single instance of the objects in the ltop-level array in Firefish export JSON file.
type Entry struct {
	ID         opt.Optional[string] `json:"id"`
	Text       opt.Optional[string] `json:"text"`
	CreatedAt  opt.Optional[string] `json:"createdAt"`
	FireIDs  []string               `json:"fileIds"`
	Files    []File                 `json:"files"`
	ReplyID    opt.Optional[string] `json:"replyId"`
	RenoteID   opt.Optional[string] `json:"renoteId"`
	//@TODO: "poll"
	//@TODO: "cw"
	Visibility opt.Optional[string] `json:"visibility"`
	//@TODO: "visibleUserIds"
	LocalOnly  opt.Optional[bool]   `json:"localOnly"`
}
