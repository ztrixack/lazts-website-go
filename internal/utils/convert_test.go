package utils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestConvertShowDayMonth(t *testing.T) {
	testTime := time.Date(2022, time.March, 15, 0, 0, 0, 0, time.UTC)
	expected := "15 Mar"
	result := ToDayMonth(testTime)
	assert.Equal(t, expected, result, "Date should be formatted as 'DD MMM'")
}

func TestConvertShowDate(t *testing.T) {
	// Setup test cases
	tests := []struct {
		name     string
		from     time.Time
		to       time.Time
		expected string
	}{
		{
			name:     "same day",
			from:     time.Date(2023, time.April, 10, 0, 0, 0, 0, time.UTC),
			to:       time.Date(2023, time.April, 10, 0, 0, 0, 0, time.UTC),
			expected: "วันจันทร์ที่ 10 เมษายน 2023",
		},
		{
			name:     "same month, different days",
			from:     time.Date(2023, time.April, 10, 0, 0, 0, 0, time.UTC),
			to:       time.Date(2023, time.April, 15, 0, 0, 0, 0, time.UTC),
			expected: "วันจันทร์ที่ 10 - วันเสาร์ที่ 15 เมษายน 2023",
		},
		{
			name:     "different months, same year",
			from:     time.Date(2023, time.April, 30, 0, 0, 0, 0, time.UTC),
			to:       time.Date(2023, time.May, 1, 0, 0, 0, 0, time.UTC),
			expected: "วันอาทิตย์ที่ 30 เมษายน - วันจันทร์ที่ 1 พฤษภาคม 2023",
		},
		{
			name:     "different years",
			from:     time.Date(2022, time.December, 31, 0, 0, 0, 0, time.UTC),
			to:       time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC),
			expected: "วันเสาร์ที่ 31 ธันวาคม 2022 - วันอาทิตย์ที่ 1 มกราคม 2023",
		},
	}

	// Execute test cases
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := ToYearMonthDayRange(tc.from, tc.to)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestCountryToFlagEmoji(t *testing.T) {
	assert.Equal(t, "\U0001F1F9\U0001F1ED", ToFlagEmoji("Thailand"), "Should return Thai flag emoji")
	assert.Equal(t, "Mars", ToFlagEmoji("Mars"), "Should return the country name if not recognized")
}
