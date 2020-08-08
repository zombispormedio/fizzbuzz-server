package fizzbuzz

const defaultInitRange = 1
const defaultEndRange = 10

// CreateFizzbuzzRange returns fizzbuzz range
func CreateFizzbuzzRange(irange ...int) []string {

	var init int
	var end int

	irangeLen := len(irange)
	if irangeLen >= 2 {
		init = irange[0]
		end = irange[1]
	} else if irangeLen == 1 {
		init = irange[0]
		end = defaultEndRange
	} else {
		init = defaultInitRange
		end = defaultEndRange
	}

	var orange []string

	for i := init; i <= end; i++ {
		orange = append(orange, ItoFizzBuzz(i))
	}

	return orange
}
