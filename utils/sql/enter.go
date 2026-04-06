package sql

import "fmt"

func ConvertSliceOrderSql(list []uint) (s string) {
	for i, u := range list {
		if i == len(list)-1 {
			s += fmt.Sprintf("id = %d desc", u)
			break
		}
		s += fmt.Sprintf("id = %d desc, ", u)
	}
	return
}
