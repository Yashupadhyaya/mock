package mock_gomock

import (
	"reflect"
	"testing"
	"github.com/golang/mock/gomock"
)









/*
ROOST_METHOD_HASH=OldMockMatcher_b241463c7a
ROOST_METHOD_SIG_HASH=OldMockMatcher_4087741302

FUNCTION_DEF=func OldMockMatcher(ctrl *gomock.Controller) *MockMatcher 

 */
func (c *customTestHelper) Errorf(format string, args ...interface{}) {
	c.t.Errorf(format, args...)
}

func (c *customTestHelper) Fatalf(format string, args ...interface{}) {
	c.t.Fatalf(format, args...)
}

func (c *customTestHelper) Helper() {
	c.t.Helper()
}

func TestOldMockMatcher(t *testing.T) {
	type testCase struct {
		name    string
		ctrl    *gomock.Controller
		wantNil bool
		checkFn func(*testing.T, *MockMatcher)
	}

	tests := []testCase{
		{
			name: "Create MockMatcher with valid Controller",
			ctrl: gomock.NewController(t),
			checkFn: func(t *testing.T, m *MockMatcher) {
				if m == nil {
					t.Error("Expected non-nil MockMatcher, got nil")
				}
				if m.ctrl == nil {
					t.Error("Expected non-nil Controller in MockMatcher, got nil")
				}
				if m.recorder == nil {
					t.Error("Expected non-nil recorder in MockMatcher, got nil")
				}
			},
		},
		{
			name: "Verify MockMatcher's recorder initialization",
			ctrl: gomock.NewController(t),
			checkFn: func(t *testing.T, m *MockMatcher) {
				if m.recorder == nil {
					t.Fatal("Expected non-nil recorder, got nil")
				}
				if m.recorder.mock != m {
					t.Error("Recorder's mock field does not point back to MockMatcher")
				}
			},
		},
		{
			name:    "Test with nil Controller",
			ctrl:    nil,
			wantNil: true,
			checkFn: func(t *testing.T, m *MockMatcher) {
				if m != nil {
					t.Error("Expected nil MockMatcher, got non-nil")
				}
			},
		},
		{
			name: "Verify uniqueness of created MockMatchers",
			ctrl: gomock.NewController(t),
			checkFn: func(t *testing.T, m *MockMatcher) {
				m2 := OldMockMatcher(m.ctrl)
				if m == m2 {
					t.Error("Expected different MockMatcher instances, got the same")
				}
			},
		},
		{
			name: "Check Controller reference integrity",
			ctrl: gomock.NewController(t),
			checkFn: func(t *testing.T, m *MockMatcher) {
				if m.ctrl != m.ctrl {
					t.Error("Controller reference in MockMatcher does not match the one passed")
				}
			},
		},
		{
			name: "Verify behavior with a custom TestHelper",
			ctrl: gomock.NewController(&customTestHelper{t: t}),
			checkFn: func(t *testing.T, m *MockMatcher) {
				if _, ok := m.ctrl.T.(*customTestHelper); !ok {
					t.Error("Expected Controller to have custom TestHelper")
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := OldMockMatcher(tt.ctrl)
			if (got == nil) != tt.wantNil {
				t.Errorf("OldMockMatcher() returned nil: %v, want nil: %v", got == nil, tt.wantNil)
			}
			if tt.checkFn != nil {
				tt.checkFn(t, got)
			}
		})
	}
}

