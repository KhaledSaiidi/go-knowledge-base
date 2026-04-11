package electionday

import "fmt"

type ElectionResult struct {
	Name  string
	Votes int
}

// NewVoteCounter returns a new vote counter with
// a given number of initial votes.
func NewVoteCounter(initialVotes int) *int {
	return &initialVotes
}

// VoteCount extracts the number of votes from a counter.
func VoteCount(counter *int) int {
	if counter == nil {
		return 0
	}
	return *counter
}

// IncrementVoteCount increments the value in a vote counter.
func IncrementVoteCount(counter *int, increment int) {
	*counter += increment
}

// NewElectionResult creates a new election result.
func NewElectionResult(candidateName string, votes int) *ElectionResult {
	electionResultsPointer := &ElectionResult{
		Name:  candidateName,
		Votes: votes,
	}
	return electionResultsPointer
}

// DisplayResult creates a message with the result to be displayed.
func DisplayResult(result *ElectionResult) string {
	return fmt.Sprintf("%s (%d)", result.Name, result.Votes)
}

// DecrementVotesOfCandidate decrements by one the vote count of a candidate in a map.
func DecrementVotesOfCandidate(results map[string]int, candidate string) {
	// votes, ok := results[candidate]
	// if ok && votes > 0 {
	//     results[candidate]--
	// }
	// Enhanced syntax:
	if votes, ok := results[candidate]; ok && votes > 0 {
		results[candidate]--
	}
}
