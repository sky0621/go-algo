package util

func Bsort(targets []int) {
	for ix := 0; ix < len(targets)-1; ix++ {
		for idx := 0; idx < len(targets)-1; idx++ {
			if targets[idx] > targets[idx+1] {
				tmp := targets[idx+1]
				targets[idx+1] = targets[idx]
				targets[idx] = tmp
			}
		}
	}
}
