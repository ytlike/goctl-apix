package utils

type MYMAP[KEY int | string, VALUE int | string | struct{} | bool | byte] map[KEY]VALUE

func RemoveRepeatElement[string comparable](list []string) []string {
	// 创建一个临时map用来存储数组元素
	temp := make(map[string]struct{})
	index := 0
	// 将元素放入map中
	for _, v := range list {
		temp[v] = struct{}{}
	}
	tempList := make([]string, len(temp))
	for key := range temp {
		tempList[index] = key
		index++
	}
	return tempList
}

// DiffArray 求两个切片的差集
func DiffArray[T int | string](a []T, b []T) []T {
	var diffArray []T
	var temp MYMAP[T, struct{}] = map[T]struct{}{}
	for _, val := range b {
		if _, ok := temp[val]; !ok {
			temp[val] = struct{}{}
		}
	}

	for _, val := range a {
		if _, ok := temp[val]; !ok {
			diffArray = append(diffArray, val)
		}
	}

	return diffArray
}

// IntersectArray 求两个切片的交集
func IntersectArray[T int | string](a []T, b []T) []T {
	var inter []T
	var mp MYMAP[T, bool] = map[T]bool{}

	for _, s := range a {
		if _, ok := mp[s]; !ok {
			mp[s] = true
		}
	}
	for _, s := range b {
		if _, ok := mp[s]; ok {
			inter = append(inter, s)
		}
	}

	return inter
}
