package validate

import (
	"github.com/gogf/gf/v2/util/gconv"
)

// 包含判断

// InSliceExistStr 判断字符或切片字符是否存在指定字符
func InSliceExistStr(elems any, search string) bool {
	switch elems.(type) {
	case []string:
		elem := gconv.Strings(elems)
		for i := 0; i < len(elem); i++ {
			if gconv.String(elem[i]) == search {
				return true
			}
		}
	default:
		return gconv.String(elems) == search
	}
	return false
}

// InSlice 元素是否存在于切片中
func InSlice[K comparable](slice []K, key K) bool {
	for _, v := range slice {
		if v == key {
			return true
		}
	}
	return false
}
