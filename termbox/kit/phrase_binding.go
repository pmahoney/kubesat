package kit

import (
	"fmt"
	"time"
)

type PhraseBinding struct {
	Phrases   map[string]string

	// The currently entered phrase prefix
	current   string

	lastInput time.Time
}

func (pb PhraseBinding) Insert(now time.Time, ch rune) (PhraseBinding, string, bool) {
	pb2 := PhraseBinding{
		Phrases: pb.Phrases,
		current: fmt.Sprintf("%s%c", pb.Current(now), ch),
		lastInput: now,
	}

	if value, ok := pb2.Phrases[pb2.current]; ok {
		pb2.current = ""
		return pb2, value, ok
	}

	return pb2, "", false
}

func (pb PhraseBinding) Current(now time.Time) string {
	if now.Sub(pb.lastInput) > (5 * time.Second) {
		return ""
	}
	return pb.current
}
