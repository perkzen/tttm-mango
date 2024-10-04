package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
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

		//5x5 grid
		//{
		//	name:       "Valid move with 5x5 grid",
		//	query:      "?gid=1234&size=5&playing=O&moves=X-0-0",
		//	wantStatus: http.StatusOK,
		//	wantBody:   "Move:O-0-1", // Assuming AI takes the next available move
		//},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "/move"+tt.query, nil)
			if err != nil {
				t.Fatalf("Could not create request: %v", err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(HandleGetMove)
			handler.ServeHTTP(rr, req)

			if rr.Code != tt.wantStatus {
				t.Errorf("Expected status %v, got %v", tt.wantStatus, rr.Code)
			}

			if rr.Body.String() != tt.wantBody {
				t.Errorf("Expected body %v, got %v", tt.wantBody, rr.Body.String())
			}
		})
	}
}
