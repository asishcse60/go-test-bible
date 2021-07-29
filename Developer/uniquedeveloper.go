package Developer

type Developer struct {
	Name string
	Age int
}

func UniqueDevelopers(developers []Developer) []string{
	mp := make(map[string]int)
	for _,dev:=range developers{
		mp[dev.Name] = 1
	}
	var uniqueDevs [] string

	for key, _ :=range mp{
		uniqueDevs = append(uniqueDevs, key)
	}
	return uniqueDevs
}