package xlsx
import (
	"regexp"
)

type Row struct {
	Cells        []*Cell
	Hidden       bool
	Sheet        *Sheet
	Height       float64
	OutlineLevel uint8
	isCustom     bool
}

func (r *Row) SetHeight(ht float64) {
	r.Height = ht
	r.isCustom = true
}

func (r *Row) SetHeightCM(ht float64) {
	r.Height = ht * 28.3464567 // Convert CM to postscript points
	r.isCustom = true
}

func (r *Row) AddCell() *Cell {
	cell := NewCell(r)
	r.Cells = append(r.Cells, cell)
	r.Sheet.maybeAddCol(len(r.Cells))
	return cell
}

// Contains(regex string) - checks whether a given regex is contained in a xlsx.Cell of the given
// xlsx.Row. It returns true the regex matches a cell in the given row.
func (r *Row) Contains(regex string) bool{
	var rst bool = false
	for _, cell := range r.Cells{
		if found, _ := regexp.Match(regex, []byte(cell.Value)); found{
			rst = true
			break
		}
	}
	return rst
}

// IdxMatchedCell(regex string) - whether a given regex is contained in a xlsx.Cell of the given
// xlsx.Row. It returns its index if found, otherwise -1.
func (r *Row) IdxMatchedCell(regex string) int{
	var rst int = -1
	for idx, cell := range r.Cells{
		if found, _ := regexp.Match(regex, []byte(cell.Value)); found{
			rst = idx
			break
		}
	}
	return rst
}
