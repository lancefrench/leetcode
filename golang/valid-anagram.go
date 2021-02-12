// In order to use sort.Sort with runes we need to implement
// a sort.Interface https://golang.org/pkg/sort/#Interface

type runes []rune

func (r runes) Len() int           { return len(r) }
func (r runes) Less(i, j int) bool { return r[i] < r[j] }
func (r runes) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }

func sortString(w string) string {
	r := []rune(w)
	sort.Sort(runes(r))
	return string(r)
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sortedStr1 := sortString(s)
	sortedStr2 := sortString(t)
	return sortedStr1 == sortedStr2
}