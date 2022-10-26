// Checking lists for equality or one being a sublist of another
package sublist

// Relation is the comparison between lists
type Relation string

// Possible relations
const (
	RelationEqual     Relation = "equal"
	RelationSublist   Relation = "sublist"
	RelationSuperlist Relation = "superlist"
	RelationUnequal   Relation = "unequal"
)

// Sublist checks difference of two lists and
// returns equal, sublist, superlist or unequal according
// to their relation to each other.
func Sublist(l1, l2 []int) Relation {
	if isEqual(l1, l2) {
		return RelationEqual
	} else if contains(l1, l2) {
		return RelationSuperlist
	} else if contains(l2, l1) {
		return RelationSublist
	}
	return RelationUnequal
}

func isEqual(l1, l2 []int) bool {
	if len(l1) != len(l2) {
		return false
	}
	for i, v := range l1 {
		if v != l2[i] {
			return false
		}
	}
	return true
}

// Returns true if l1 contains l2
func contains(l1, l2 []int) bool {
	if len(l2) > len(l1) {
		return false
	}
outer:
	for i := range l1 {
		if i+len(l2) > len(l1) {
			return false
		}
		for j, itemOfL2 := range l2 {
			if itemOfL2 != l1[i+j] {
				continue outer
			}
		}
		return true
	}
	return false
}
