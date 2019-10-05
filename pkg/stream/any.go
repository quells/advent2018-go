package stream

type AnyToAny func(interface{}) interface{}
type AnyFilter func(interface{}) bool

type Any struct {
	values     []interface{}
	transforms []interface{}
}

func OfAnys(a ...interface{}) Any {
	return Any{values: a}
}

func (a Any) Map(f AnyToAny) Any {
	a.transforms = append(a.transforms, f)
	return a
}

func (a Any) Filter(f AnyFilter) Any {
	a.transforms = append(a.transforms, f)
	return a
}

func (a Any) Collect() []interface{} {
	for _, t := range a.transforms {
		switch t.(type) {
		case AnyToAny:
			f := t.(AnyToAny)
			tmp := make([]interface{}, len(a.values))
			for i, v := range a.values {
				tmp[i] = f(v)
			}
			a.values = tmp
			break
		case AnyFilter:
			f := t.(AnyFilter)
			removed := 0
			tmp := make([]interface{}, len(a.values))
			for i, v := range a.values {
				j := i - removed
				if f(v) {
					tmp[j] = v
				} else {
					removed++
				}
			}
			a.values = tmp[:len(tmp)-removed]
			break
		}
	}
	return a.values
}
