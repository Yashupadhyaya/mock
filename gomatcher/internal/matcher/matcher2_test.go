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
		name     string
		ctrl     *gomock.Controller
		wantNil  bool
		wantFunc func(*testing.T, *MockMatcher)
	}{
		{
			name: "Create MockMatcher with Valid Controller",
			ctrl: gomock.NewController(t),
			wantFunc: func(t *testing.T, m *MockMatcher) {
				if m == nil {
					t.Error("Expected non-nil MockMatcher, got nil")
				}
				if m.ctrl == nil {
					t.Error("Expected non-nil controller in MockMatcher")
				}
				if m.recorder == nil {
					t.Error("Expected non-nil recorder in MockMatcher")
				}
			},
		},
		{
			name: "Verify Recorder Initialization",
			ctrl: gomock.NewController(t),
			wantFunc: func(t *testing.T, m *MockMatcher) {
				if m.recorder == nil {
					t.Fatal("Expected non-nil recorder")
				}
				if m.recorder.mock != m {
					t.Error("Recorder's mock field does not point back to MockMatcher")
				}
			},
		},
		{
			name: "Multiple Calls to OldMockMatcher",
			ctrl: gomock.NewController(t),
			wantFunc: func(t *testing.T, m *MockMatcher) {
				m2 := OldMockMatcher(m.ctrl)
				if m == m2 {
					t.Error("Expected different MockMatcher instances")
				}
				if m.ctrl != m2.ctrl {
					t.Error("Expected same controller for both MockMatcher instances")
				}
			},
		},
		{
			name:    "Nil Controller Handling",
			ctrl:    nil,
			wantNil: true,
		},
		{
			name: "Controller State After MockMatcher Creation",
			ctrl: gomock.NewController(t),
			wantFunc: func(t *testing.T, m *MockMatcher) {

				if m.ctrl == nil {
					t.Error("Controller should not be nil")
				}

			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
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

			if tt.wantFunc != nil {
				tt.wantFunc(t, got)
			}
		})
	}
}

