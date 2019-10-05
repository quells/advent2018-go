package stream

type StringToString func(string) string
type StringFilter func(string) bool

type String struct {
	values     []string
	transforms []interface{}
}

func OfStrings(s ...string) String {
	return String{values: s}
}

func (s String) Map(f StringToString) String {
	s.transforms = append(s.transforms, f)
	return s
}

func (s String) Filter(f StringFilter) String {
	s.transforms = append(s.transforms, f)
	return s
}

func (s String) Collect() (results []string) {
	for _, t := range s.transforms {
		switch t.(type) {
		case StringToString:
			f := t.(StringToString)
			tmp := make([]string, len(s.values))
			for i, v := range s.values {
				tmp[i] = f(v)
			}
			s.values = tmp
			break
		case StringFilter:
			f := t.(StringFilter)
			removed := 0
			tmp := make([]string, len(s.values))
			for i, v := range s.values {
				j := i - removed
				if f(v) {
					tmp[j] = v
				} else {
					removed++
				}
			}
			s.values = tmp[:len(tmp)-removed]
			break
		}
	}
	return s.values
}

type StringToInt func(string) int

func (s String) ToInts(f StringToInt) Int {
	ss := s.Collect()
	x := make([]int, len(ss))
	for i, s := range ss {
		x[i] = f(s)
	}
	return OfInts(x...)
}
