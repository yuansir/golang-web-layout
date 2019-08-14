package utils

import (
	"golang-echo-layout/config"
	"math"
	"reflect"
	"strconv"
	"strings"
)

// SplitErrors split err error to string slice
func SplitErrors(err error) []string {
	return strings.Split(err.Error(), ";")
}

func InArray(needle interface{}, haystak interface{}) (exists bool, index int) {
	exists = false
	index = -1

	switch reflect.TypeOf(haystak).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(haystak)
		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(needle, s.Index(i).Interface()) {
				index = i
				exists = true
				return
			}
		}
	}

	return
}

func Pagination(payload map[string]interface{}, total int) (pagination map[string]interface{}) {
	pagination = map[string]interface{}{
		"total": total,
	}
	if limit, ok := payload["per_page"]; ok {
		pagination["per_page"], _ = strconv.Atoi(limit.(string))
	} else {
		pagination["per_page"] = config.Conf.App.PageSize
	}

	if page, ok := payload["page"]; ok {
		pagination["current_page"], _ = strconv.Atoi(page.(string))
	} else {
		pagination["current_page"] = 1
	}

	pagination["total_pages"] = int(math.Ceil(float64(total) / float64(pagination["per_page"].(int))))
	return
}