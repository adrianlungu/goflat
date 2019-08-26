package goflat

func Flatten(s []interface{}) []int {

	var result = make([]int, 0)

	for _, i := range s {

		switch t := i.(type) {
		case []interface{}:
			f := Flatten(t)
			result = append(result, f...)
		case int:
			result = append(result, t)
		}

	}

	return result

}


