package turnbased

import (
	"strconv"
	"bytes"
)

type Transcript string

// CellAddress is coded as "CharDigit" where Char is X and Digit is Y, e.g. A1, B2, C3, etc.
type CellAddress string

type Size CellAddress

func (ca CellAddress) X() int {
	return int(rune(ca[0]) - 'A')
}

func NewSize(width, height int) Size {
	var s bytes.Buffer
	s.WriteRune('A' +rune(width)-1)
	if height <= 9 {
		s.WriteRune(rune(height-1))
	} else {
		s.WriteString(strconv.Itoa(height))
	}
	return Size(s.String())
}

func NewCellAddress(x, y int) CellAddress {
	return CellAddress([]rune{
		'A'+rune(x),
		'1'+rune(y),
	})
}

func (ca CellAddress) IsXY(x, y int) bool {
	return len(ca) != 0 && ca.X() == x && ca.Y() == y
}

func (ca CellAddress) XY() (x, y int) {
	return ca.X(), ca.Y()
}

func (wh Size) Width() (width int) {
	return CellAddress(wh).X()+1
}

func (wh Size) Height() (height int) {
	return CellAddress(wh).Y()+1
}

func (wh Size) WidthHeight() (width, height int) {
	return wh.Width(), wh.Height()
}

func (ca CellAddress) Y() int {
	if len(ca) == 2 {
		return int(rune(ca[1]) - '1')
	}
	if y, err := strconv.Atoi(string(ca[1:])); err != nil {
		panic(err)
	} else {
		return y-1
	}
}

func(t Transcript) Cells() (cells []CellAddress) {
	l := len(t)
	cells = make([]CellAddress, l/2)
	var c int
	for i := 0; i < l; i+=2  {
		cells[c] = CellAddress(t[i:i+2])
		c++
	}
	return
}

func (t Transcript) Count() int {
	if l := len(t); l%2 != 0 {
		panic("length is not even")
	} else {
		return l/2
	}
}