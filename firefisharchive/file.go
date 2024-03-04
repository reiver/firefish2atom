package firefisharchive

import (
	"sourcecode.social/reiver/go-opt"
)

// File represents a file in the top-level list of entries in the Firefish archive JSON file.
type File struct {
	ID          opt.Optional[string]  `json:"id"`
	CreatedAt   opt.Optional[string]  `json:"createdAt"`
	Name        opt.Optional[string]  `json:"name"`
	Type        opt.Optional[string]  `json:"type"`
	MD5         opt.Optional[string]  `json:"md5"`
	Size        opt.Optional[uint64]  `json:"size"`
	IsSensitive opt.Optional[bool]    `json:"isSensitive"`
	BlurHash    opt.Optional[string]  `json:"blurhash"`
	Properties struct{
		Width  opt.Optional[uint64] `json:"width"`
		Height opt.Optional[uint64] `json:"height"`
	} `json:"properties"`
        URL          opt.Optional[string] `json:"url"`
        ThumbnailURL opt.Optional[string] `json:"thumbnailUrl"`
	Comment     opt.Optional[string]  `json:"comment"`
	//@TODO: "folderId"
	//@TODO: "folder"
	//@TODO: "userId"
	//@TODO: "user"
}
