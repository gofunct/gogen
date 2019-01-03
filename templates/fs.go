//+build vfsgen

package templates

import "net/http"

var VirtualFS http.FileSystem = http.Dir("./_data")
