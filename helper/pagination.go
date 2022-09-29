package helper

func BuildPagination(reqPage, reqLimit int) (int, int, int) {
	page := 1
	if reqPage > 0 {
		page = reqPage
	}

	limit := 10
	if reqLimit > 0 {
		limit = reqLimit
	}

	offset := limit * (page - 1)

	return page, limit, offset
}
