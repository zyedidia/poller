// +build windows

package poller

import "os"

type FD struct {
	os.File
}
