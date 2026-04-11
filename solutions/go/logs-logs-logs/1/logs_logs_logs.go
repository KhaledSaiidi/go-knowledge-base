package logs

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, r := range log {
		switch r {
		case '❗':
			return "recommendation"
		case '🔍':
			return "search"
		case '☀':
			return "weather"
		}
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	result := []rune(log)
	for i, r := range result {
		if r == oldRune{
            result[i] = newRune
        }
    }
return string(result)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	count := 0
   	for range log {
    	count++
    }
    return count <= limit
}
