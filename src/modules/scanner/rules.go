package scanner

var (
	equal    = []string{}
	contains = []string{}
	internal = []string{}
)

func addNewRule(k string, r []string) {
	switch k {
	case "equal":
		equal = r
	case "contains":
		contains = r
	case "internal":
		internal = r
	}
}
