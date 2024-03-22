package main

import "testing"

func TestRandomBytes(t *testing.T) {
	t.Parallel()

	t.Run("it should generate 16 bytes", func(t *testing.T) {
		t.Parallel()

		generatedBytes, err := RandomBytes(&RandStrct{NumberOfBytes: 16})
		if err != nil {
			t.Fatalf("RandomBytes returned an error: %s", err)
		}

		if len(generatedBytes) != 32 {
			t.Fatalf("number of bytes does not match the requested number: %d instead of %d", len(generatedBytes)/2, 16)
		}
	})

	t.Run("it should generate 100000 bytes", func(t *testing.T) {
		t.Parallel()

		generatedBytes, err := RandomBytes(&RandStrct{NumberOfBytes: 100000})
		if err != nil {
			t.Fatalf("RandomBytes returned an error: %s", err)
		}

		if len(generatedBytes) != 200000 {
			t.Fatalf("number of bytes does not match the requested number: %d instead of %d", len(generatedBytes)/2, 100000)
		}
	})
}
