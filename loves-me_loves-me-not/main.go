package lovesmelovesmenot

import "strings"

func LovesMe(itr int) string {
	var parts = []string{"Loves me not", "Loves me"}
	var str = []string{}
	for i := 1; i <= itr; i++ {
		var appendString = parts[i%2]
		if i == itr {
			appendString = strings.ToUpper(appendString)
		}
		str = append(str, appendString)
	}
	return strings.Join(str, ", ")
}
