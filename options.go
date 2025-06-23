package ufs

import "log"

// options.go

type Options struct {
	ShowError      bool
	ReturnReadable bool
}

type UFS struct {
	opts Options
}

var dufs *UFS = &UFS{
	opts: Options{
		ShowError:      true,
		ReturnReadable: false,
	},
}

func NewUfs(opts *Options) *UFS {
	if opts == nil {
		opts = &Options{}
	}
	return &UFS{opts: *opts}
}

// NewOptions creates a new Options instance with default values.
func NewOptions() *Options {
	return &Options{
		ShowError:      false,
		ReturnReadable: true,
	}
}

func (ufs *UFS) SetOptions(opts *Options) {
	if opts == nil {
		ufs.opts = *NewOptions()
	} else {
		ufs.opts = *opts
	}
}

func (ufs *UFS) handleError(err error, operation ...string) {
	if ufs.opts.ShowError {
		if len(operation) > 0 {
			// Log the error with operation context
			log.Printf("%s: %v", operation[0], err)
		} else {
			log.Println(err)
		}
	}

	// Simply do nothing if ShowError is false
}

func (ufs *UFS) handleMistakeWarning(mesage string) {
	if ufs.opts.ShowError {
		log.Println(mesage)
	}
}
