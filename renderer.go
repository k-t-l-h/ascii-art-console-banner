package banner

import (
	"github.com/fatih/color"
	"unicode/utf8"
)

type Renderer interface {
	Render(string)
}

type AsciiRender struct {
	opts renderOption
}

func NewAsciiRender(opt ...Option) Renderer {
	opts := defaultRenderOptions
	for _, option := range opt {
		option.apply(&opts)
	}

	if opts.ft == nil {
		opts.ft, _ = LoadFont(defaultFont)
	}

	return &AsciiRender{opts: opts}
}


func (a* AsciiRender) Render(text string)  {

	text = a.opts.normFunc(text)
	allSet := make([][]string, utf8.RuneCountInString(text))

	textSet := []rune(text)
	for index, char := range textSet {
		sl := a.opts.ft.getRuneRepresentation(char)
		allSet[index] = sl
	}

	cl := color.New()
	cl.Add(a.opts.color)
	if a.opts.bold {
		cl = cl.Add(color.Bold)
	}
	cl.SetWriter(a.opts.output)


	for i := 0; i < a.opts.ft.height; i++ {
		line := ""
		for _, set := range allSet {
			line += set[i]
		}
		cl.Println(line)
	}
}

