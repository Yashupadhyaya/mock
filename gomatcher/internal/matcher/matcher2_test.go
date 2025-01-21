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
		checkPtr bool
	}{
		{
			name:     "Create MockMatcher with valid Controller",
			ctrl:     gomock.NewController(t),
			wantNil:  false,
			checkPtr: false,
		},
		{
			name:     "Test with nil Controller",
			ctrl:     nil,
			wantNil:  true,
			checkPtr: false,
		},
		{
			name:     "Verify consistency across multiple calls",
			ctrl:     gomock.NewController(t),
			wantNil:  false,
			checkPtr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := OldMockMatcher(tt.ctrl)

			if tt.wantNil && result != nil {
				t.Errorf("OldMockMatcher() = %v, want nil", result)
			}

			if !tt.wantNil && result == nil {
				t.Errorf("OldMockMatcher() returned nil, want non-nil")
			}

			if result != nil {
				if result.ctrl != tt.ctrl {
					t.Errorf("OldMockMatcher().ctrl = %v, want %v", result.ctrl, tt.ctrl)
				}

				if result.recorder == nil {
					t.Errorf("OldMockMatcher().recorder is nil")
				}

				if _, ok := result.recorder.(*MockMatcherMockRecorder); !ok {
					t.Errorf("OldMockMatcher().recorder is not of type *MockMatcherMockRecorder")
				}

				if result.recorder.(*MockMatcherMockRecorder).mock != result {
					t.Errorf("OldMockMatcher().recorder.mock does not point back to the MockMatcher")
				}
			}

			if tt.checkPtr {
				result2 := OldMockMatcher(tt.ctrl)
				if result == result2 {
					t.Errorf("Multiple calls to OldMockMatcher() returned the same instance")
				}
				if result.ctrl != result2.ctrl {
					t.Errorf("Multiple calls to OldMockMatcher() did not share the same Controller")
				}
			}
		})
	}
}

