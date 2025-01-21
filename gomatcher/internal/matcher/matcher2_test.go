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
		name      string
		ctrl      *gomock.Controller
		wantNil   bool
		wantPanic bool
	}{
		{
			name:    "Create MockMatcher with valid Controller",
			ctrl:    gomock.NewController(t),
			wantNil: false,
		},
		{
			name:    "Create MockMatcher with nil Controller",
			ctrl:    nil,
			wantNil: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result *MockMatcher
			defer func() {
				r := recover()
				if (r != nil) != tt.wantPanic {
					t.Errorf("OldMockMatcher() panic = %v, wantPanic = %v", r, tt.wantPanic)
				}
			}()

			result = OldMockMatcher(tt.ctrl)

			if (result == nil) != tt.wantNil {
				t.Errorf("OldMockMatcher() returned nil: %v, want nil: %v", result == nil, tt.wantNil)
			}

			if !tt.wantNil {
				if result.ctrl != tt.ctrl {
					t.Errorf("OldMockMatcher().ctrl = %v, want %v", result.ctrl, tt.ctrl)
				}

				if result.recorder == nil {
					t.Error("OldMockMatcher().recorder is nil")
				} else {
					if _, ok := result.recorder.(*MockMatcherMockRecorder); !ok {
						t.Errorf("OldMockMatcher().recorder type = %v, want *MockMatcherMockRecorder", reflect.TypeOf(result.recorder))
					}
					if result.recorder.(*MockMatcherMockRecorder).mock != result {
						t.Error("OldMockMatcher().recorder.mock does not reference the MockMatcher")
					}
				}
			}
		})
	}

	t.Run("Create multiple MockMatchers", func(t *testing.T) {
		ctrl1 := gomock.NewController(t)
		ctrl2 := gomock.NewController(t)

		mock1 := OldMockMatcher(ctrl1)
		mock2 := OldMockMatcher(ctrl2)

		if mock1 == mock2 {
			t.Error("OldMockMatcher() returned the same instance for different controllers")
		}

		if mock1.ctrl != ctrl1 || mock2.ctrl != ctrl2 {
			t.Error("OldMockMatcher() did not set the correct controller for each instance")
		}
	})
}

