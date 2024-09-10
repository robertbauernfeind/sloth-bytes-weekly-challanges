package lovesmelovesmenot

import "strings"

func LovesMe(itr int) string {
	var str = []string{}
	for i := 1; i <= itr; i++ {
		var appendString = ""
		if i%2 != 0 {
			appendString = "Loves me"
		}
		if i%2 == 0 {
			appendString = "Loves me not"
		}

		if i == itr {
			appendString = strings.ToUpper(appendString)
		}
		str = append(str, appendString)
	}
	return strings.Join(str, ", ")
}
