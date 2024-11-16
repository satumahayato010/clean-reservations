package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewReservation(t *testing.T) {
	tests := []struct {
		name        string
		id          string
		userID      string
		evseUID     string
		startTime   time.Time
		endTime     time.Time
		status      Status
		expectError bool
		errorMsg    string
	}{
		{
			name:        "valid inputs",
			id:          "res1",
			userID:      "user1",
			evseUID:     "evse1",
			startTime:   time.Now(),
			endTime:     time.Now().Add(2 * time.Hour),
			status:      Active,
			expectError: false,
		},
		{
			name:        "empty inputs",
			id:          "",
			userID:      "",
			evseUID:     "",
			startTime:   time.Now(),
			endTime:     time.Now().Add(2 * time.Hour),
			status:      Active,
			expectError: true,
			errorMsg:    "id, userID, and evseUID cannot bi empty",
		},
		{
			name:        "startTime after endTime",
			id:          "res1",
			userID:      "user1",
			evseUID:     "evse1",
			startTime:   time.Now(),
			endTime:     time.Now().Add(-2 * time.Hour),
			status:      Active,
			expectError: true,
			errorMsg:    "startTime must be before endTime",
		},
		{
			name:        "invalid status",
			id:          "res1",
			userID:      "user1",
			evseUID:     "evse1",
			startTime:   time.Now(),
			endTime:     time.Now().Add(2 * time.Hour),
			status:      Status(999),
			expectError: true,
			errorMsg:    "invalid status",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reservation, err := NewReservation(tt.id, tt.userID, tt.evseUID, tt.startTime, tt.endTime, tt.status)

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, reservation)
				assert.EqualError(t, err, tt.errorMsg)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, reservation)
				assert.Equal(t, tt.id, reservation.id)
				assert.Equal(t, tt.userID, reservation.userID)
				assert.Equal(t, tt.evseUID, reservation.evseUID)
				assert.Equal(t, tt.startTime, reservation.startTime)
				assert.Equal(t, tt.endTime, reservation.endTime)
			}
		})
	}
}
