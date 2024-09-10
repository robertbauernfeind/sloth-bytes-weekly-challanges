package lovesmelovesmenot

import "strings"

func LovesMe(itr int) string {
	var str = []string{}
	for i := 1; i <= itr; i++ {
		if i%2 != 0 {
			if i == itr {
				str = append(str, "LOVES ME")
				break
			}
			str = append(str, "Loves me")
		}
		if i%2 == 0 {
			if i == itr {
				str = append(str, "LOVES ME NOT")
				break
			}
			str = append(str, "Loves me not")
		}
	}
	return strings.Join(str, ", ")
}
