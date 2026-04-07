package techpalace
import "strings"
import "fmt"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	starsNum := strings.Repeat("*", numStarsPerLine)
    k := fmt.Sprintf("%s\n%s\n%s", starsNum, welcomeMsg, starsNum)
    fmt.Sprintln(k)
    return k
    }

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	lines := strings.Split(oldMsg, "\n")
	if len(lines) < 2 {
		return strings.TrimSpace(strings.ReplaceAll(oldMsg, "*", ""))
	}

	messageLine := lines[1]
	messageLine = strings.ReplaceAll(messageLine, "*", "")
	messageLine = strings.TrimSpace(messageLine)

	return messageLine
}