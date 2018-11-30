package kit

import (
	"testing"
	"time"
)

func TestPhraseBinding(t *testing.T) {
	phrases := make(map[string]string)
	phrases["hello"] = "hello"
	phrases["del"] = "delete"
	pb := PhraseBinding{Phrases: phrases}

	{
		var value string
		var ok bool
		now := time.Now()
		pb, _, _ = pb.Insert(now, 'h')
		pb, _, _ = pb.Insert(now, 'e')
		pb, _, _ = pb.Insert(now, 'l')
		pb, _, _ = pb.Insert(now, 'l')
		pb, value, ok = pb.Insert(now, 'o')
		if !ok {
			t.Fatalf("expected value, got %s, %v", value, ok)
		}
		if value != "hello" {
			t.Fatalf("expected '%s' but got '%s'", "hello", value)
		}

		if pb.Current(now) != "" {
			t.Fatalf("expected current to be cleared after success but was '%s'", pb.Current(now))
		}
	}

	{
		now := time.Now()
		pb, _, _ = pb.Insert(now.Add(1*time.Second), 'd')
		pb, _, _ = pb.Insert(now.Add(2*time.Second), 'e')
		_, value, ok := pb.Insert(now.Add(3*time.Second), 'l')
		if !ok {
			t.Fatalf("expected value, got '%s', %v", value, ok)
		}
		if value != "delete" {
			t.Fatalf("expected '%s' but got '%s'", "delete", value)
		}
	}
}

func TestPhraseBindingTimeout(t *testing.T) {
	phrases := make(map[string]string)
	phrases["hello"] = "hello"
	phrases["del"] = "delete"
	pb := PhraseBinding{Phrases: phrases}

	now := time.Now()
	then := now.Add(6 * time.Second)

	pb, _, _ = pb.Insert(now, 'h')
	pb, _, _ = pb.Insert(now, 'e')
	pb, _, _ = pb.Insert(now, 'l')
	pb, _, _ = pb.Insert(now, 'l')
	pb, value, ok := pb.Insert(then, 'o')
	if ok {
		t.Fatalf("expected no value but got %s, %v", value, ok)
	}

	if pb.Current(then) != "o" {
		t.Fatalf("expected '%s' but got '%s'", "o", pb.Current(then))
	}
}
