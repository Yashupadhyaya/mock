package mock_gomock

import (
	"reflect"
	"testing"
	"github.com/golang/mock/gomock"
)









/*
ROOST_METHOD_HASH=OldMockMatcher_b241463c7a
ROOST_METHOD_SIG_HASH=OldMockMatcher_4087741302

undefined
 */
func TestOldMockMatcher(t *testing.T) {
	tests := []struct {
		name      string
		ctrl      *gomock.Controller
		wantNil   bool
		checkFn   func(*testing.T, *MockMatcher)
		wantPanic bool
	}{
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
			},
		},
		{
			name: "Verify MockMatcher recorder initialization",
			ctrl: gomock.NewController(t),
			checkFn: func(t *testing.T, m *MockMatcher) {
				if m.recorder == nil {
					t.Error("Expected non-nil recorder, got nil")
				}
				if _, ok := m.recorder.(*MockMatcherMockRecorder); !ok {
					t.Errorf("Expected recorder of type *MockMatcherMockRecorder, got %T", m.recorder)
				}
			},
		},
		{
			name:    "Test with nil Controller",
			ctrl:    nil,
			wantNil: true,
		},
		{
			name: "Verify uniqueness of created MockMatcher instances",
			ctrl: gomock.NewController(t),
			checkFn: func(t *testing.T, m *MockMatcher) {
				m2 := OldMockMatcher(m.ctrl)
				if m == m2 {
					t.Error("Expected different MockMatcher instances, got the same instance")
				}
			},
		},
		{
			name: "Check Controller assignment consistency",
			ctrl: gomock.NewController(t),
			checkFn: func(t *testing.T, m *MockMatcher) {
				if !reflect.DeepEqual(m.ctrl, m.ctrl) {
					t.Error("Expected MockMatcher to have the same Controller that was passed in")
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil && !tt.wantPanic {
					t.Errorf("OldMockMatcher() panicked unexpectedly: %v", r)
				}
			}()

			got := OldMockMatcher(tt.ctrl)

			if tt.wantNil && got != nil {
				t.Errorf("OldMockMatcher() = %v, want nil", got)
			}

			if tt.checkFn != nil {
				tt.checkFn(t, got)
			}
		})
	}
}

