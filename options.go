package ufs

import (
	"fmt"
	"log"

	"github.com/utsav-56/ulog"
)

// options.go

type Options struct {
	ShowError      bool
	ReturnReadable bool
	prettifyError  bool // If true, prettify the error messages
}

type UFS struct {
	opts Options
}

var dufs *UFS = &UFS{
	opts: Options{
		ShowError:      true,
		ReturnReadable: false,
		prettifyError:  true, // Default to prettifying errors
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

		errMessage := err.Error()

		if ufs.opts.prettifyError {
			ulog.Error(errMessage, operation...)
			return
		}

		if len(operation) > 0 {
			// Log the error with operation context
			log.Printf("%s: %v", operation[0], errMessage)
		} else {
			log.Println(errMessage)
		}
	}

	// Simply do nothing if ShowError is false
}

func (ufs *UFS) handleMistakeWarning(mesage string) {
	if ufs.opts.ShowError {
		if ufs.opts.prettifyError {
			ulog.Warning(mesage)
			return
		}
		log.Println(mesage)
	}
}

// wrapError is a helper function to wrap errors with function names
func (ufs *UFS) wrapError(err error, functionName string) error {
	if err != nil {
		return fmt.Errorf("%s: %w", functionName, err)
	}
	return nil
}
