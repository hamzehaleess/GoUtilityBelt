package time

import (
        "testing"
        "time"
)

func TestParseDuration(t *testing.T) {
        tests := []struct {
                input    string
                expected time.Duration
                hasError bool
        }{
                {"1h", time.Hour, false},
                {"30m", 30 * time.Minute, false},
                {"45s", 45 * time.Second, false},
                {"1h30m", time.Hour + 30*time.Minute, false},
                {"invalid", 0, true},
        }

        for _, test := range tests {
                result, err := ParseDuration(test.input)
                if test.hasError {
                        if err == nil {
                                t.Errorf("ParseDuration(%q) expected error, got nil", test.input)
                        }
                } else {
                        if err != nil {
                                t.Errorf("ParseDuration(%q) unexpected error: %v", test.input, err)
                        }
                        if result != test.expected {
                                t.Errorf("ParseDuration(%q) = %v; expected %v", test.input, result, test.expected)
                        }
                }
        }
}

func TestFormatDuration(t *testing.T) {
        tests := []struct {
                input    time.Duration
                expected string
        }{
                {30 * time.Second, "30s"},
                {90 * time.Second, "1m 30s"},
                {3661 * time.Second, "1h 1m"},
                {25 * time.Hour, "1d 1h"},
        }

        for _, test := range tests {
                result := FormatDuration(test.input)
                if result != test.expected {
                        t.Errorf("FormatDuration(%v) = %q; expected %q", test.input, result, test.expected)
                }
        }
}

func TestIsWeekend(t *testing.T) {
        tests := []struct {
                input    time.Time
                expected bool
        }{
                {time.Date(2023, 7, 1, 0, 0, 0, 0, time.UTC), true},  // Saturday
                {time.Date(2023, 7, 2, 0, 0, 0, 0, time.UTC), true},  // Sunday
                {time.Date(2023, 7, 3, 0, 0, 0, 0, time.UTC), false}, // Monday
                {time.Date(2023, 7, 4, 0, 0, 0, 0, time.UTC), false}, // Tuesday
        }

        for _, test := range tests {
                result := IsWeekend(test.input)
                if result != test.expected {
                        t.Errorf("IsWeekend(%v) = %v; expected %v", test.input, result, test.expected)
                }
        }
}

func TestBusinessDays(t *testing.T) {
        start := time.Date(2023, 7, 3, 0, 0, 0, 0, time.UTC)  // Monday
        end := time.Date(2023, 7, 7, 0, 0, 0, 0, time.UTC)    // Friday
        
        result := BusinessDays(start, end)
        expected := 5 // Monday through Friday
        
        if result != expected {
                t.Errorf("BusinessDays(%v, %v) = %d; expected %d", start, end, result, expected)
        }
}

func TestTimeAgo(t *testing.T) {
        now := time.Now()
        tests := []struct {
                input    time.Time
                expected string
        }{
                {now.Add(-30 * time.Second), "just now"},
                {now.Add(-2 * time.Minute), "2 minutes ago"},
                {now.Add(-1 * time.Hour), "1 hour ago"},
                {now.Add(-25 * time.Hour), "1 day ago"},
                {now.Add(-8 * 24 * time.Hour), "8 days ago"},
        }

        for _, test := range tests {
                result := TimeAgo(test.input)
                if result != test.expected {
                        t.Errorf("TimeAgo(%v) = %q; expected %q", test.input, result, test.expected)
                }
        }
}

func TestAddBusinessDays(t *testing.T) {
        // Test with a Monday (not weekend)
        monday := time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC) // January 1, 2024 is a Monday
        
        tests := []struct {
                input    time.Time
                days     int
                expected time.Time
        }{
                {monday, 0, monday},
                {monday, 1, monday.AddDate(0, 0, 1)}, // Tuesday
                {monday, 5, monday.AddDate(0, 0, 7)}, // Next Monday (skipping weekend)
        }

        for _, test := range tests {
                result := AddBusinessDays(test.input, test.days)
                if !result.Equal(test.expected) {
                        t.Errorf("AddBusinessDays(%v, %d) = %v; expected %v", test.input, test.days, result, test.expected)
                }
        }
}

func TestStartOfDay(t *testing.T) {
        input := time.Date(2023, 7, 1, 15, 30, 45, 123456789, time.UTC)
        expected := time.Date(2023, 7, 1, 0, 0, 0, 0, time.UTC)
        
        result := StartOfDay(input)
        if !result.Equal(expected) {
                t.Errorf("StartOfDay(%v) = %v; expected %v", input, result, expected)
        }
}

func TestEndOfDay(t *testing.T) {
        input := time.Date(2023, 7, 1, 15, 30, 45, 123456789, time.UTC)
        expected := time.Date(2023, 7, 1, 23, 59, 59, 999999999, time.UTC)
        
        result := EndOfDay(input)
        if !result.Equal(expected) {
                t.Errorf("EndOfDay(%v) = %v; expected %v", input, result, expected)
        }
}

func TestIsLeapYear(t *testing.T) {
        tests := []struct {
                input    int
                expected bool
        }{
                {2020, true},
                {2021, false},
                {2000, true},
                {1900, false},
                {2024, true},
        }

        for _, test := range tests {
                result := IsLeapYear(test.input)
                if result != test.expected {
                        t.Errorf("IsLeapYear(%d) = %v; expected %v", test.input, result, test.expected)
                }
        }
}

func TestDaysInMonth(t *testing.T) {
        tests := []struct {
                year     int
                month    time.Month
                expected int
        }{
                {2023, time.January, 31},
                {2023, time.February, 28},
                {2024, time.February, 29}, // Leap year
                {2023, time.April, 30},
                {2023, time.December, 31},
        }

        for _, test := range tests {
                result := DaysInMonth(test.year, test.month)
                if result != test.expected {
                        t.Errorf("DaysInMonth(%d, %v) = %d; expected %d", test.year, test.month, result, test.expected)
                }
        }
}
