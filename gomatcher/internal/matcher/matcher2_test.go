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
func TestOldMockMatcher(t *testing.T) {
	tests := []struct {
		name    string
		ctrl    *gomock.Controller
		wantNil bool
		checkFn func(*testing.T, *MockMatcher)
	}{
		{
			name: "Create MockMatcher with valid Controller",
			ctrl: gomock.NewController(t),
			checkFn: func(t *testing.T, m *MockMatcher) {
				if m == nil {
					t.Error("Expected non-nil MockMatcher, got nil")
				}
				if m.ctrl == nil {
					t.Error("Expected non-nil Controller in MockMatcher")
				}
				if m.recorder == nil {
					t.Error("Expected non-nil recorder in MockMatcher")
				}
			},
		},
		{
			name: "Verify MockMatcher's recorder initialization",
			ctrl: gomock.NewController(t),
			checkFn: func(t *testing.T, m *MockMatcher) {
				if m.recorder == nil {
					t.Fatal("Expected non-nil recorder")
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
			name: "Check Controller assignment",
			ctrl: gomock.NewController(t),
			checkFn: func(t *testing.T, m *MockMatcher) {
				if m.ctrl != m.ctrl {
					t.Error("Controller not correctly assigned to MockMatcher")
				}
			},
		},
		{
			name: "Verify MockMatcher type conformance",
			ctrl: gomock.NewController(t),
			checkFn: func(t *testing.T, m *MockMatcher) {

			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil && !tt.wantNil {
					t.Errorf("OldMockMatcher panicked: %v", r)
				}
			}()

			got := OldMockMatcher(tt.ctrl)

			if tt.wantNil {
				if got != nil {
					t.Errorf("OldMockMatcher() = %v, want nil", got)
				}
				return
			}

			if got == nil {
				t.Fatal("OldMockMatcher() returned nil, want non-nil")
			}

			if !reflect.DeepEqual(got.ctrl, tt.ctrl) {
				t.Errorf("OldMockMatcher().ctrl = %v, want %v", got.ctrl, tt.ctrl)
			}

			if tt.checkFn != nil {
				tt.checkFn(t, got)
			}
		})
	}
}

