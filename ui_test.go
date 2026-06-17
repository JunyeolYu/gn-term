package main

import (
	"testing"
	"time"
)

func TestFormatRelativeTime(t *testing.T) {
	now := time.Date(2026, 6, 17, 12, 0, 0, 0, time.FixedZone("KST", 9*60*60))

	tests := []struct {
		name        string
		publishedAt time.Time
		want        string
	}{
		{
			name:        "less than minute",
			publishedAt: now.Add(-30 * time.Second),
			want:        "just now",
		},
		{
			name:        "single minute",
			publishedAt: now.Add(-1 * time.Minute),
			want:        "1 minute ago",
		},
		{
			name:        "minutes",
			publishedAt: now.Add(-45 * time.Minute),
			want:        "45 minutes ago",
		},
		{
			name:        "single hour",
			publishedAt: now.Add(-1 * time.Hour),
			want:        "1 hour ago",
		},
		{
			name:        "hours",
			publishedAt: now.Add(-3*time.Hour - 20*time.Minute),
			want:        "3 hours ago",
		},
		{
			name:        "future",
			publishedAt: now.Add(5 * time.Minute),
			want:        "just now",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := formatRelativeTime(now, tt.publishedAt); got != tt.want {
				t.Errorf("formatRelativeTime() = %q, want %q", got, tt.want)
			}
		})
	}
}
