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
			name:      "Valid Controller",
			ctrl:      gomock.NewController(t),
			wantNil:   false,
			wantPanic: false,
		},
		{
			name:      "Nil Controller",
			ctrl:      nil,
			wantNil:   true,
			wantPanic: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result *MockMatcher
			defer func() {
				r := recover()
				if (r != nil) != tt.wantPanic {
					t.Errorf("OldMockMatcher() panic = %v, wantPanic %v", r, tt.wantPanic)
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
					if reflect.TypeOf(result.recorder) != reflect.TypeOf(&MockMatcherMockRecorder{}) {
						t.Errorf("OldMockMatcher().recorder type = %v, want *MockMatcherMockRecorder", reflect.TypeOf(result.recorder))
					}
					if result.recorder.mock != result {
						t.Error("OldMockMatcher().recorder.mock does not point back to the MockMatcher")
					}
				}
			}
		})
	}

	t.Run("Multiple invocations consistency", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		result1 := OldMockMatcher(ctrl)
		result2 := OldMockMatcher(ctrl)

		if result1 == nil || result2 == nil {
			t.Error("OldMockMatcher() returned nil")
		}

		if result1 == result2 {
			t.Error("OldMockMatcher() returned the same instance for multiple invocations")
		}

		if result1.ctrl != result2.ctrl {
			t.Error("OldMockMatcher() returned instances with different controllers")
		}
	})
}

