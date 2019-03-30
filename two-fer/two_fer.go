/* Twofer implements "One for you and one for me"
If given name is not empty, It will echo "One for <name>, one for me"
*/
package twofer

// ShareWith will echo out "One for <name>, one for me" with given name as a param, or "you" if the name is empty
func ShareWith(name string) string {

	if len(name) > 0 {
		return "One for " + name + ", one for me."
	} else {
		return "One for you, one for me."
	}
}
