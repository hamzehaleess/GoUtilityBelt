// Package time provides utility functions for time and date operations
package time

import (
        "fmt"
        "time"
)

// FormatDuration formats a duration into a human-readable string
func FormatDuration(d time.Duration) string {
        if d < time.Minute {
                return fmt.Sprintf("%.0fs", d.Seconds())
        }
        if d < time.Hour {
                totalSeconds := int(d.Seconds())
                minutes := totalSeconds / 60
                seconds := totalSeconds % 60
                return fmt.Sprintf("%dm %ds", minutes, seconds)
        }
        if d < 24*time.Hour {
                totalMinutes := int(d.Minutes())
                hours := totalMinutes / 60
                minutes := totalMinutes % 60
                return fmt.Sprintf("%dh %dm", hours, minutes)
        }
        days := int(d.Hours() / 24)
        hours := int(d.Hours()) % 24
        return fmt.Sprintf("%dd %dh", days, hours)
}

// ParseDuration parses a duration string with custom format
func ParseDuration(s string) (time.Duration, error) {
        return time.ParseDuration(s)
}

// IsWeekend checks if a given time is on a weekend
func IsWeekend(t time.Time) bool {
        weekday := t.Weekday()
        return weekday == time.Saturday || weekday == time.Sunday
}

// BusinessDays calculates the number of business days between two dates
func BusinessDays(start, end time.Time) int {
        if start.After(end) {
                start, end = end, start
        }
        
        days := 0
        for d := start; d.Before(end) || d.Equal(end); d = d.AddDate(0, 0, 1) {
                if !IsWeekend(d) {
                        days++
                }
        }
        return days
}

// TimeAgo returns a human-readable string representing how long ago a time was
func TimeAgo(t time.Time) string {
        now := time.Now()
        diff := now.Sub(t)
        
        if diff < time.Minute {
                return "just now"
        }
        if diff < time.Hour {
                minutes := int(diff.Minutes())
                if minutes == 1 {
                        return "1 minute ago"
                }
                return fmt.Sprintf("%d minutes ago", minutes)
        }
        if diff < 24*time.Hour {
                hours := int(diff.Hours())
                if hours == 1 {
                        return "1 hour ago"
                }
                return fmt.Sprintf("%d hours ago", hours)
        }
        
        days := int(diff.Hours() / 24)
        if days == 1 {
                return "1 day ago"
        }
        if days < 30 {
                return fmt.Sprintf("%d days ago", days)
        }
        if days < 365 {
                months := days / 30
                if months == 1 {
                        return "1 month ago"
                }
                return fmt.Sprintf("%d months ago", months)
        }
        
        years := days / 365
        if years == 1 {
                return "1 year ago"
        }
        return fmt.Sprintf("%d years ago", years)
}

// StartOfDay returns the start of the day for a given time
func StartOfDay(t time.Time) time.Time {
        return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

// EndOfDay returns the end of the day for a given time
func EndOfDay(t time.Time) time.Time {
        return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 999999999, t.Location())
}

// AddBusinessDays adds business days to a given time
func AddBusinessDays(t time.Time, days int) time.Time {
        result := t
        for i := 0; i < days; {
                result = result.AddDate(0, 0, 1)
                if !IsWeekend(result) {
                        i++
                }
        }
        return result
}

// IsLeapYear checks if a given year is a leap year
func IsLeapYear(year int) bool {
        return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

// DaysInMonth returns the number of days in a given month and year
func DaysInMonth(year int, month time.Month) int {
        return time.Date(year, month+1, 0, 0, 0, 0, 0, time.UTC).Day()
}
