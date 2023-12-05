package paginatejson

import (
	"math"
	"paginatejson/lib"
)

func OffsetByPageSize(listOutstd []interface{}, page, size int) (resultListOutstd []interface{}, pagination Page) {
	totalData := len(listOutstd)

	cutStartIdx, cutLength, totalPage, currentPage, minPage, maxPage, resSize := paginate(totalData, size, page)

	// set result
	resultListOutstd = listOutstd[cutStartIdx:cutLength]

	pagination.Items = resultListOutstd
	pagination.TotalPages = int64(totalPage)
	pagination.Page = int64(currentPage)
	pagination.MaxPage = int64(maxPage)
	pagination.First = page == minPage
	pagination.Last = page == maxPage
	pagination.Size = int64(resSize)
	pagination.Total = int64(totalData)
	pagination.Visible = int64(len(resultListOutstd))
	return
}

func paginate(totalData, size, page int) (cutStartIdx, cutLength, totalPage, currentPage, minPage, maxPage, finalSize int) {
	isValidTotalData := validateTotalData(totalData)
	if !isValidTotalData {
		return
	}

	finalSize = formatSize(size, totalData)

	totalPage = genTotalPage(totalData, finalSize)

	isValidTotalPage := validateTotalPage(totalPage)
	if !isValidTotalPage {
		return
	}

	currentPage, minPage, maxPage = formatPage(totalPage, page)

	cutStartIdx, cutLength = genCutStartLength(totalData, currentPage, finalSize)

	return
}

func validateTotalData(totalData int) (isValid bool) {
	if lib.IsEmptyInt(totalData) {
		return
	}

	isValid = true
	return
}

func formatSize(size, totalData int) (resultSize int) {
	minSize := 0
	maxSize := totalData

	if size < minSize || size > maxSize {
		size = maxSize
	}

	resultSize = size
	return
}

func genTotalPage(totalData, size int) (totalPage int) {
	float64TotalData := float64(totalData)
	float64Size := float64(size)

	divideResult := float64TotalData / float64Size
	roundUp := math.Ceil(divideResult)
	totalPage = int(roundUp)
	return
}

func validateTotalPage(totalPage int) (isValid bool) {
	if lib.IsEmptyInt(totalPage) {
		return
	}

	isValid = true
	return
}

func formatPage(totalPage, page int) (resultPage, minPage, maxPage int) {
	minPage = 0
	maxPage = totalPage - 1

	if page < minPage {
		page = minPage
	} else if page > maxPage {
		page = maxPage
	}

	resultPage = page
	return
}

func genCutStartLength(totalData, page, size int) (startIdx, len int) {
	maxLen := totalData

	startIdx = (page * size)

	len = (page * size) + size
	if len > maxLen {
		len = maxLen
	}

	return
}
