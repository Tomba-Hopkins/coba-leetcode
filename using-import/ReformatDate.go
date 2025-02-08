package usingimportpackage

import (
	"strconv"
	"strings"
)

// Question :

// Given a date string in the form Day Month Year, where:
// Day is in the set {"1st", "2nd", "3rd", "4th", ..., "30th", "31st"}.
// Month is in the set {"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}.
// Year is in the range [1900, 2100].
// Convert the date string to the format YYYY-MM-DD, where:
// YYYY denotes the 4 digit year.
// MM denotes the 2 digit month.
// DD denotes the 2 digit day.
// Example 1:
// Input: date = "20th Oct 2052"
// Output: "2052-10-20"
// Example 2:
// Input: date = "6th Jun 1933"
// Output: "1933-06-06"
// Example 3:
// Input: date = "26th May 1960"
// Output: "1960-05-26"
// Constraints:
// The given dates are guaranteed to be valid, so no error handling is necessary.

// Answer :
func ReformatDate(date string) string {
    a := strings.Split(date, " ")
    month := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
    m := map[string]string{}
    for i, x := range month {
        s := strconv.Itoa(i + 1)
        m[x] = s
    }
    dig := "0123456789"
    d := map[string]bool{}
    for _, x := range dig {
        d[string(x)] = true
    }
    t := ""
    if d[string(a[0][1])] {
        t = string(a[0][:2])
        } else {
        t = string(a[0][:1])
    }
    b := m[a[1]]

    if len(t) < 2 {
        t = "0" + t
    }
    if len(b) < 2 {
        b = "0" + b
    }    
    r := []string{a[2], b, t}
	return strings.Join(r, "-")
}