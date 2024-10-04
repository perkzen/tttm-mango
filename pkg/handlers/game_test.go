package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestMoveHandler(t *testing.T) {
	tests := []struct {
		name       string
		query      string
		wantStatus int
		wantBody   string
	}{
		// Invalid queries
		{
			name:       "Invalid player symbol",
			query:      "?gid=1234&size=3&playing=Z&moves=",
			wantStatus: http.StatusBadRequest,
			wantBody:   "Error:Sorry. Can't do it bro.",
		},
		// 3x3 grid
		{
			name:       "Valid move with 3x3 grid",
			query:      "?gid=1234&size=3&playing=O&moves=X-0-0",
			wantStatus: http.StatusOK,
			wantBody:   "Move:O-0-1",
		},
		{
			name:       "Valid move with 3x3 grid, middle move",
			query:      "?gid=1234&size=3&playing=X&moves=O-1-1",
			wantStatus: http.StatusOK,
			wantBody:   "Move:X-0-0", // Assuming AI prefers the corner
		},
		{
			name:       "AI winning move on 3x3 grid",
			query:      "?gid=1234&size=3&playing=X&moves=X-0-0_O-1-0_X-0-1_O-1-1",
			wantStatus: http.StatusOK,
			wantBody:   "Move:X-0-2", // X should win by placing in (0,2)
		},
		{
			name:       "Block opponent from winning",
			query:      "?gid=1234&size=3&playing=X&moves=O-0-0_O-0-1",
			wantStatus: http.StatusOK,
			wantBody:   "Move:X-0-2", // Assuming AI plays to block
		},
		{
			name:       "AI winning move on 3x3 grid",
			query:      "?gid=1234&size=3&playing=X&moves=X-0-0_O-1-0_X-0-1_O-1-1",
			wantStatus: http.StatusOK,
			wantBody:   "Move:X-0-2", // X should win by placing in (0,2)
		},
		{
			name:       "No moves made yet",
			query:      "?gid=1234&size=3&playing=X&moves=",
			wantStatus: http.StatusOK,
			wantBody:   "Move:X-0-0", // First move is usually a corner
		},

		// 5x5 grid
		{
			name:       "Valid move with 5x5 grid",
			query:      "?gid=1234&size=5&playing=O&moves=X-0-0",
			wantStatus: http.StatusOK,
			wantBody:   "Move:O-0-4", // Assuming AI takes the next available move
		},
		{
			name:       "Should block opponent from winning",
			query:      "?gid=1234&size=5&playing=X&moves=O-0-0_O-0-1_O-0-2_O-0-3_O-0-4_X-1-0_X-1-1_X-1-2_X-1-3",
			wantStatus: http.StatusOK,
			wantBody:   "Move:X-1-4", // Assuming AI blocks opponent from winning
		},
		{
			name:       "AI should win",
			wantStatus: http.StatusOK,
			query:      "?gid=1234&size=5&playing=X&moves=O-0-0_O-0-1_O-0-2_O-0-3_O-0-4_X-1-0_X-1-1_X-1-2_X-1-3_X-1-4",
			wantBody:   "Move:X-2-0",
		},

		// 7x7 grid
		{
			name:       "Valid move with 7x7 grid",
			query:      "?gid=1234&size=7&playing=O&moves=X-0-0",
			wantStatus: http.StatusOK,
			wantBody:   "Move:O-0-6", // Assuming AI takes the next available move
		},
		{
			name:       "Should block opponent from winning",
			query:      "?gid=1234&size=7&playing=X&moves=O-0-0_O-0-1_O-0-2_O-0-3_O-0-4_O-0-5_O-0-6_X-1-0_X-1-1_X-1-2_X-1-3",
			wantStatus: http.StatusOK,
			wantBody:   "Move:X-1-4", // Assuming AI blocks opponent from winning
		},
		{
			name:       "AI should win",
			query:      "?gid=1234&size=7&playing=X&moves=O-0-0_O-0-1_O-0-2_O-0-3_O-0-4_O-0-5_O-0-6_X-1-0_X-1-1_X-1-2_X-1-3_X-1-4",
			wantStatus: http.StatusOK,
			wantBody:   "Move:X-1-5",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "/move"+tt.query, nil)
			if err != nil {
				t.Fatalf("Could not create request: %v", err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(HandleGetMove)

			startTime := time.Now() // Start time measurement
			handler.ServeHTTP(rr, req)
			elapsedTime := time.Since(startTime) // Calculate elapsed time

			if rr.Code != tt.wantStatus {
				t.Errorf("Expected status %v, got %v", tt.wantStatus, rr.Code)
			}

			if rr.Body.String() != tt.wantBody {
				t.Errorf("Expected body %v, got %v", tt.wantBody, rr.Body.String())
			}

			if tt.name == "Response time should be less than 3s" && elapsedTime >= 3*time.Second {
				t.Errorf("Expected response time to be less than 3 seconds, got %v", elapsedTime)
			}
		})
	}
}
