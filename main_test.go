package main

import (
	"testing"
)

func TestIsHeavy(t *testing.T) {
	tests := []struct {
		name     string
		mass     int
		expected bool
	}{
		{"Light_package", 15, false},
		{"Heavy_package", 25, true},
		{"Heavy_boundary_package", 20, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isHeavy(tt.mass)
			if result != tt.expected {
				t.Errorf("isHeavy(%d): got %t, want %t", tt.mass, result, tt.expected)
			}
		})
	}
}

func TestIsBulky(t *testing.T) {
	tests := []struct {
		name     string
		width    int
		height   int
		length   int
		expected bool
	}{
		{"Small_dimensions", 10, 10, 10, false},
		{"Bulky_by_width", 160, 10, 10, true},
		{"Bulky_by_heigh", 10, 160, 10, true},
		{"Bulky_by_length", 10, 10, 160, true},
		{"Bulky_by_volume", 100, 100, 100, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isBulky(tt.width, tt.height, tt.length)
			if result != tt.expected {
				t.Errorf("isBulky(%d, %d, %d): got %t, want %t", tt.width, tt.height, tt.length, result, tt.expected)
			}
		})
	}
}

func TestSort(t *testing.T) {
	tests := []struct {
		name     string
		width    int
		height   int
		length   int
		mass     int
		expected packageType
	}{
		{"Standard_package", 10, 10, 10, 10, stardard},
		{"Special_by_volume", 100, 100, 100, 10, special},
		{"Special_by_weight", 10, 10, 10, 25, special},
		{"Rejected_package", 100, 100, 100, 25, rejected},
		{"Rejected_by_width_and_weight", 200, 10, 10, 25, rejected},
		{"Invalid_by_mass", 10, 10, 10, -3, invalid},
		{"Invalid_by_width", -10, 10, 10, 10, invalid},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := sort(tt.width, tt.height, tt.length, tt.mass)
			if result != tt.expected {
				t.Errorf("sort(%d, %d, %d, %d): got %s, want %s", tt.width, tt.height, tt.length, tt.mass, result, tt.expected)
			}
		})
	}
}
