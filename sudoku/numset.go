package sudoku

// NumSet is a set where the possible members of the set are integers
// between 1 and a specified maximum number.
type NumSet struct {
	a   []bool
	cnt int
}

func NewNumSet(maxNum int) NumSet {
	return NumSet{
		a:   make([]bool, maxNum),
		cnt: 0,
	}
}

func (set *NumSet) Fill() {
	for i := range len(set.a) {
		set.a[i] = true
	}
	set.cnt = len(set.a)
}

func (set *NumSet) SetNum(num int) {
	if num < 1 || num > len(set.a) {
		panic("SetNum called with number out of range")
	}
	set.a[num-1] = true
	set.cnt++
}

func (set *NumSet) ClearNum(num int) {
	if num < 1 || num > len(set.a) {
		panic("ClearNum called with number out of range")
	}
	set.a[num-1] = false
	set.cnt--
}

func (set *NumSet) Has(num int) bool {
	return num > 0 && num < len(set.a)+1 && set.a[num-1]
}

func (set *NumSet) Count() int {
	return set.cnt
}

func (set *NumSet) MaxNum() int {
	return len(set.a)
}
