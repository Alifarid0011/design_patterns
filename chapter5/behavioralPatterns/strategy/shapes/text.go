package shapes

import (
	"design_patterns/chapter5/behavioralPatterns/strategy"
)

type TextSquare struct {
	strategy.PrintOutput
}

//	func (t *TextSquare) Print() error {
//		r := bytes.NewReader([]byte("Circle"))
//		io.Copy(t.Writer, r)
//		return nil
//	}
func (t *TextSquare) Print() error {
	t.Writer.Write([]byte("Circle"))
	return nil
}
