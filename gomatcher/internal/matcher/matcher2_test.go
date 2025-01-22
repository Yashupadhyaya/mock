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
			name:     "Create MockMatcher with nil Controller",
			ctrl:     nil,
			wantNil:  true,
			checkPtr: false,
		},
		{
			name:     "Verify uniqueness of created MockMatchers",
			ctrl:     gomock.NewController(t),
			wantNil:  false,
			checkPtr: true,
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
				t.Errorf("OldMockMatcher() returned nil, want non-nil")
				return
			}

			if got.ctrl != tt.ctrl {
				t.Errorf("OldMockMatcher().ctrl = %v, want %v", got.ctrl, tt.ctrl)
			}

			if got.recorder == nil {
				t.Errorf("OldMockMatcher().recorder is nil")
			} else if got.recorder.mock != got {
				t.Errorf("OldMockMatcher().recorder.mock = %v, want %v", got.recorder.mock, got)
			}

			if tt.checkPtr {
				got2 := OldMockMatcher(tt.ctrl)
				if got == got2 {
					t.Errorf("OldMockMatcher() returned same instance for multiple calls")
				}
			}

			_ = got.EXPECT()
		})
	}
}

