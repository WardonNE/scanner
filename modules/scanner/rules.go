package scanner

var (
	equal = []string{
		"#",
		"###",
		"/",
		"",
	}
	contains = []string{
		".jpg",
		".png",
		".gif",
		".jpeg",
		".zip",
		".rar",
		".pdf",
		".doc",
		".docx",
		".xls",
		".xlsx",
		"tel",
		"mailto",
		"javascript",
		"/uploads",
		"/themes",
		"/privatefile",
		".mp3",
		".mp4",
		".gz",
		".bmp",
		".wav",
	}
	self = []string{
		"www.",
		".com",
		"https",
		"http",
	}
)

func AddNewRule(k string, r string) {
	switch k {
	case "equal":
		equal = append(equal, r)
	case "contains":
		contains = append(contains, r)
	case "self":
		self = append(self, r)
	}
}

func ListRules(k string) []string {
	switch k {
	case "equal":
		return equal
	case "contains":
		return contains
	case "self":
		return self
	default:
		return []string{}
	}
}

func ListAllRules() map[string][]string {
	var rules = make(map[string][]string)
	rules["equal"] = equal
	rules["contains"] = contains
	rules["self"] = self
	return rules
}
