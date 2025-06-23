package ufs

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
