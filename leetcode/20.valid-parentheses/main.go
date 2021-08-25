package main

func isValid(s string) bool {
	var stacks []byte

	for i := 0; i < len(s); i++ {
		if s[i] == '[' || s[i] == '{' || s[i] == '(' {
			stacks = append(stacks, s[i])
		} else if (s[i] == '}' && len(stacks) > 0 && stacks[len(stacks)-1] == '{') ||
			(s[i] == ']' && len(stacks) > 0 && stacks[len(stacks)-1] == '[') ||
			(s[i] == ')' && len(stacks) > 0 && stacks[len(stacks)-1] == '(') {
			stacks = stacks[:len(stacks)-1]
		} else {
			return false
		}
	}

	return len(stacks) == 0
}

func main() {

}
