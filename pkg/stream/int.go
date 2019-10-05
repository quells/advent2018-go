package stream

type IntToInt func(int) int
type IntFilter func(int) bool

type Int struct {
	values     []int
	transforms []interface{}
	repeat     bool
}

func OfInts(x ...int) Int {
	return Int{values: x}
}

func (x Int) Repeated() Int {
	x.repeat = true
	return x
}

func (x Int) Map(f IntToInt) Int {
	x.transforms = append(x.transforms, f)
	return x
}

func (x Int) Filter(f IntFilter) Int {
	x.transforms = append(x.transforms, f)
	return x
}

func (x Int) Collect() (results []int) {
	for _, t := range x.transforms {
		switch t.(type) {
		case IntToInt:
			f := t.(IntToInt)
			tmp := make([]int, len(x.values))
			for i, v := range x.values {
				tmp[i] = f(v)
			}
			x.values = tmp
			break
		case IntFilter:
			f := t.(IntFilter)
			removed := 0
			tmp := make([]int, len(x.values))
			for i, v := range x.values {
				j := i - removed
				if f(v) {
					tmp[j] = v
				} else {
					removed++
				}
			}
			x.values = tmp[:len(tmp)-removed]
			break
		}
	}
	return x.values
}

func (x Int) MustReduce(agg int, f func(a, b int) int) int {
	values := x.Collect()
	for _, v := range values {
		agg = f(agg, v)
	}
	return agg
}

func (x Int) Reduce(agg int, f func(a, b int) (int, error)) int {
	var err error
	values := x.Collect()
	for _, v := range values {
		agg, err = f(agg, v)
		if err != nil {
			return agg
		}
	}
	if x.repeat {
		return OfInts(values...).Repeated().Reduce(agg, f)
	}
	return agg
}

func (x Int) Sum() int {
	return x.MustReduce(0, func(a, b int) int { return a + b })
}
