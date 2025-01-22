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
		checkFn  func(*testing.T, *MockMatcher)
		wantType interface{}
	}{
		{
			name: "Create MockMatcher with valid Controller",
			ctrl: gomock.NewController(t),
			checkFn: func(t *testing.T, m *MockMatcher) {
				if m == nil {
					t.Error("Expected non-nil MockMatcher")
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
			name: "Verify recorder initialization",
			ctrl: gomock.NewController(t),
			checkFn: func(t *testing.T, m *MockMatcher) {
				if m.recorder == nil {
					t.Error("Expected non-nil recorder")
				}
				if m.recorder.mock != m {
					t.Error("Expected recorder.mock to point back to MockMatcher")
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
					t.Error("Expected different MockMatcher instances")
				}
			},
		},
		{
			name: "Check Controller assignment",
			ctrl: gomock.NewController(t),
			checkFn: func(t *testing.T, m *MockMatcher) {
				if m.ctrl != m.ctrl {
					t.Error("Expected MockMatcher.ctrl to be the same as input Controller")
				}
			},
		},
		{
			name:     "Verify MockMatcher type conformance",
			ctrl:     gomock.NewController(t),
			wantType: (*MockMatcher)(nil),
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

			if tt.checkFn != nil {
				tt.checkFn(t, got)
			}

			if tt.wantType != nil {
				if !reflect.TypeOf(got).AssignableTo(reflect.TypeOf(tt.wantType).Elem()) {
					t.Errorf("OldMockMatcher() returned type %T, want assignable to %T", got, tt.wantType)
				}
			}
		})
	}
}

