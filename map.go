package common


//MapSearch: Search key by value and return key
func MapSearch(search interface{}, target map[interface{}]interface{})(interface{}, bool){
	var result interface{}
	ok := false
	for k,v := range target{
		if v == search {
			result = k
			ok = true
			break
		}
	}
	return result, ok
}
