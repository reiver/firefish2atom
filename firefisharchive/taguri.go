package firefisharchive

import (
	"fmt"
)

// tagURI returns a tag-uri.
//
// This is used as the value of the <id> in the Atom <entry>.
//
// 'date' is assumed to have a format of YYYY-MM-DD.
// A longer date formats (that starts with YYYY-MM-DD) can be passed to it,
// and tagURI will truncate it to just that part.
func tagURI(internetDomain string, date string, id string) string {
	return fmt.Sprintf("tag:%s:%s:%s", internetDomain, date[:10], id)
}
