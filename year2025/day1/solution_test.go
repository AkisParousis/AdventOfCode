package main

import "testing"

func TestDial(t *testing.T) {
    tests := []struct {
        direction      string
        degrees        int
        start          int
        expectedDial   int
        expectedWraps  int
    }{
        // Right turn that does not cross zero
        {"R", 10, 50, 60, 0},      // 50 + 10 = 60, no wrap
        // Right turn that crosses zero
        {"R", 60, 50, 10, 1},      // 50 + 60 = 110 → wrap once → 10
        // Right turn that does not change position
        {"R", 0, 50, 50, 0},
        // Right turn that crosses zero multiple times
        {"R", 200, 50, 50, 2},     // 50 + 200 = 250 → wrap twice → 50
        // Right turn that lands exactly on zero
        {"R", 50, 50, 0, 1},       // 50 + 50 = 100 → wrap once → 0
        // Right turn that lands exactly on zero and crosses zero multiple times
        {"R", 350, 50, 0, 4},     // 50 + 350 = 400 → wrap four times → 0
        // Right turn that starts from zero and no crossing
        {"R", 30, 0, 30, 0},       // 0 + 30 = 30, no wrap
        // Right turn that starts from zero and crosses zero multiple times
        {"R", 340, 0, 40, 3},     // 0 + 340 = 340 → wrap three times → 40
        // Right turn that starts from zero and lands exactly on zero
        {"R", 100, 0, 0, 1},      // 0 + 300 = 300 → wrap three times → 0
        {"R", 400, 0, 0, 4},      // 0 + 400 = 400 → wrap four times → 0
        
        // Left turn that does not cross zero
        {"L", 10, 50, 40, 0},      // 50 - 10 = 40, no wrap
        // Left turn that crosses zero
        {"L", 20, 10, 90, 1},      // 10 - 20 = -10 → wrap once → 90
        // Left turn that does not change position
        {"L", 0, 50, 50, 0},
        // Left turn that crosses zero multiple times
        {"L", 200, 50, 50, 2},     // 50 - 200 = -150 → wrap twice → 50
        // Left turn that starts from zero and crosses zero
        {"L", 5, 0, 95, 0},        // 0 - 5 = -5 → wrap once → 95
        // Left turn that lands exactly on zero
        {"L", 50, 50, 0, 1},
        // Left turn that starts from zero and lands exactly on zero
        {"L", 400, 0, 0, 4},
        // Left turn that lands exactly on zero and crosses zero multiple times
        {"L", 250, 50, 0, 3},      // 50 - 250 = -200 → wrap three times → 0
        {"L", 390, 0, 10, 3},     // 0 - 390 = -390 → wrap three times → 10
    }

    for _, tt := range tests {
        gotDial, gotWraps := dial(tt.direction, tt.degrees, tt.start)
        if gotDial != tt.expectedDial || gotWraps != tt.expectedWraps {
            t.Errorf("dial(%s, %d, %d) = (%d, %d); want (%d, %d)",
                tt.direction, tt.degrees, tt.start,
                gotDial, gotWraps, tt.expectedDial, tt.expectedWraps)
        }
    }
}
