package banner

import (
	"github.com/fatih/color"
	"io"
	"os"
)


var defaultRenderOptions = renderOption{
	output: os.Stdout,
	ft: nil,
	bold: false,
	color: color.FgCyan,
	normFunc: func(s string) string {return s},

}

type renderOption struct {
	output io.Writer
	ft 	*font
	bold bool
	color color.Attribute
	normFunc func(string) string
}

type Option interface {
	apply(*renderOption)
}

type funcRenderOption struct {
	f func(option *renderOption)
}

func (fdo *funcRenderOption) apply(opt *renderOption) {
	fdo.f(opt)
}

func newFuncRenderOption(f func(option *renderOption)) *funcRenderOption {
	return &funcRenderOption{
		f: f,
	}
}

func WithCustomWriter(w io.Writer) Option {
	return newFuncRenderOption(func(o *renderOption) {
		o.output = w
	})
}

func WithFont(src string) Option {
	return newFuncRenderOption(func(o *renderOption) {
		f, err := LoadFontFile(src)
		if err != nil {
			f, _ = LoadFont(defaultFont)
		}
		o.ft = f
	})
}

func WithBoldness(flag bool)  Option {
	return newFuncRenderOption(func(o *renderOption) {
		o.bold = flag
	})
}


func WithNormalisation(f func(string)string)  Option {
	return newFuncRenderOption(func(o *renderOption) {
		o.normFunc = f
	})
}


func WithColor(color color.Attribute)  Option {
	return newFuncRenderOption(func(o *renderOption) {
		o.color = color
	})
}
