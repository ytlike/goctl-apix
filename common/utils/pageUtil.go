package utils

type Pager struct {
	Offset int32
	Limit  int32
}

// Paging 构建分页对象
func Paging(page, size, totalSize int32) *Pager {
	if page < 1 {
		page = 1
	}
	if page*size > totalSize {
		if totalSize%size == 0 {
			page = totalSize / size
		} else {
			page = totalSize/size + 1
		}
	}
	pager := &Pager{}
	pager.Offset = (page - 1) * size
	pager.Limit = size
	return pager
}
