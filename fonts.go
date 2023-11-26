package banner

import (
	"errors"
	"io"
	"os"
	"strconv"
	"strings"
)

const fontPrefix = "flf2a"
const printableOffset = 32

type font struct {
	hardblank string
	height int
	baseline int
	fontSlice []string
}

func LoadFontFile(scr string) (*font, error) {
	if _, err := os.Stat(scr); err != nil {
		return nil, err
	}

	file, err := os.Open(scr)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	buf := new(strings.Builder)
	_, err = io.Copy(buf, file)
	if err != nil {
		return nil, err
	}

	return LoadFont(buf.String())
}

//          flf2a$ 6 5 20 15 3 0 143 229    NOTE: The first five characters in
//            |  | | | |  |  | |  |   |     the entire file must be "flf2a".
//           /  /  | | |  |  | |  |   \
//  Signature  /  /  | |  |  | |   \   Codetag_Count
//    Hardblank  /  /  |  |  |  \   Full_Layout*
//         Height  /   |  |   \  Print_Direction
//         Baseline   /    \   Comment_Lines
//          Max_Length      Old_Layout*
func LoadFont(content string) (*font, error) {
	lines := strings.Split(content, "\n")
	if len(lines) < 1 {
		return nil, errors.New("no data")
	}

	metadata := strings.Split(lines[0], " ")
	if !strings.HasPrefix(metadata[0], fontPrefix) {
		return nil, errors.New("font file does not contain flf2a signature")
	}
	hardblank := strings.TrimPrefix(metadata[0],fontPrefix)
	fontHeight, err := strconv.Atoi(metadata[1])
	if err!= nil {
		return nil, errors.New("font file does not contain font size")
	}
	commentLines, err := strconv.Atoi(metadata[5])
	if err!= nil {
		return nil, errors.New("font file does not contain comment lines size")
	}
	return &font{
		hardblank: hardblank,
		height: fontHeight,
		fontSlice: lines[commentLines+1:],
	}, nil
}

func (f *font) getRuneRepresentation(char rune) []string {

	beginRow := (int(char) - printableOffset) * f.height
	lines := make([]string,  f.height)

	//if there is no such symbol in a font we just skip it
	if beginRow > len(f.fontSlice) {
		return lines
	}

	for i := 0; i <  f.height; i++ {
		row := f.fontSlice[beginRow+i]
		row = strings.Replace(row, "@", "", -1)
		row = strings.Replace(row, f.hardblank, " ", -1)
		lines[i] = row
	}

	return lines
}