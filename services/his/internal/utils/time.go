package utils

import "time"

func ParseDatePtr(value string) (*time.Time, error) {
	if value == "" {
		return nil, nil
	}

	t, err := time.Parse("2006-01-02", value)
	if err != nil {
		return nil, err
	}

	return &t, nil
}
