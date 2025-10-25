package models

import (
	"testing"
	"time"
)

func TestCalculateAge(t *testing.T) {
	tests := []struct {
		name     string
		dob      time.Time
		expected int
	}{
		{
			name:     "Born in 1990, should be 34 in 2024",
			dob:      time.Date(1990, 5, 10, 0, 0, 0, 0, time.UTC),
			expected: 34, // This will need to be updated based on current year
		},
		{
			name:     "Born in 2000, should be 24 in 2024",
			dob:      time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
			expected: 24, // This will need to be updated based on current year
		},
		{
			name:     "Born today, should be 0",
			dob:      time.Now(),
			expected: 0,
		},
		{
			name:     "Born yesterday, should be 0",
			dob:      time.Now().AddDate(0, 0, -1),
			expected: 0,
		},
		{
			name:     "Born one year ago, should be 1",
			dob:      time.Now().AddDate(-1, 0, 0),
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user := &User{DOB: tt.dob}
			age := user.CalculateAge()
			
			// For tests with specific years, we need to calculate expected age dynamically
			if tt.name == "Born in 1990, should be 34 in 2024" {
				expectedAge := time.Now().Year() - 1990
				if time.Now().YearDay() < tt.dob.YearDay() {
					expectedAge--
				}
				if age != expectedAge {
					t.Errorf("CalculateAge() = %v, want %v", age, expectedAge)
				}
			} else if tt.name == "Born in 2000, should be 24 in 2024" {
				expectedAge := time.Now().Year() - 2000
				if time.Now().YearDay() < tt.dob.YearDay() {
					expectedAge--
				}
				if age != expectedAge {
					t.Errorf("CalculateAge() = %v, want %v", age, expectedAge)
				}
			} else {
				if age != tt.expected {
					t.Errorf("CalculateAge() = %v, want %v", age, tt.expected)
				}
			}
		})
	}
}

func TestToResponse(t *testing.T) {
	dob := time.Date(1990, 5, 10, 0, 0, 0, 0, time.UTC)
	user := &User{
		ID:   1,
		Name: "John Doe",
		DOB:  dob,
	}

	response := user.ToResponse()

	if response.ID != 1 {
		t.Errorf("Expected ID 1, got %d", response.ID)
	}

	if response.Name != "John Doe" {
		t.Errorf("Expected name 'John Doe', got '%s'", response.Name)
	}

	if response.DOB != "1990-05-10" {
		t.Errorf("Expected DOB '1990-05-10', got '%s'", response.DOB)
	}

	// Age should be calculated correctly
	expectedAge := user.CalculateAge()
	if response.Age != expectedAge {
		t.Errorf("Expected age %d, got %d", expectedAge, response.Age)
	}
}



