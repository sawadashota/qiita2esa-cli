package esa


func ExistMember(esaMembers []string, qiitaUser string) bool {
	for _, member := range esaMembers {
		if qiitaUser == member {
			return true
		}
	}
	return false
}

