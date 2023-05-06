package easy

import (
	"fmt"
	"strings"
)

// Reference: https://leetcode.com/problems/unique-email-addresses/
func Leetcode_Num_Unique_Emails() {
	fmt.Println("Input: emails = ['test.email+alex@leetcode.com','test.e.mail+bob.cathy@leetcode.com','testemail+david@lee.tcode.com']")
	fmt.Println("Output:", numUniqueEmails([]string{"test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"}))
	fmt.Println("Input: emails = ['a@leetcode.com','b@leetcode.com','c@leetcode.com']")
	fmt.Println("Output:", numUniqueEmails([]string{"a@leetcode.com", "b@leetcode.com", "c@leetcode.com"}))
}

func numUniqueEmails(emails []string) int {
	m := make(map[string]bool)
	for _, v := range emails {
		parts := strings.Split(v, "@")
		email := strings.Split(strings.Join(strings.Split(parts[0], "."), ""), "+")[0] + "@" + parts[1]
		m[email] = true
	}
	return len(m)
}
