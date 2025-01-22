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
			name:    "Create MockMatcher with Valid Controller",
			ctrl:    gomock.NewController(t),
			wantNil: false,
		},
		{
			name:    "Test with Nil Controller",
			ctrl:    nil,
			wantNil: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				r := recover()
				if (r != nil) != tt.wantPanic {
					t.Errorf("OldMockMatcher() panic = %v, wantPanic = %v", r, tt.wantPanic)
				}
			}()

			got := OldMockMatcher(tt.ctrl)

			if (got == nil) != tt.wantNil {
				t.Errorf("OldMockMatcher() returned nil: %v, want nil: %v", got == nil, tt.wantNil)
				return
			}

			if !tt.wantNil {
				if got.ctrl != tt.ctrl {
					t.Errorf("OldMockMatcher().ctrl = %v, want %v", got.ctrl, tt.ctrl)
				}

				if got.recorder == nil {
					t.Error("OldMockMatcher().recorder is nil")
				} else {
					if _, ok := got.recorder.(*MockMatcherMockRecorder); !ok {
						t.Errorf("OldMockMatcher().recorder is not of type *MockMatcherMockRecorder")
					}
					if got.recorder.(*MockMatcherMockRecorder).mock != got {
						t.Errorf("OldMockMatcher().recorder.mock does not point back to the MockMatcher")
					}
				}
			}
		})
	}

	t.Run("Test Multiple Invocations", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mock1 := OldMockMatcher(ctrl)
		mock2 := OldMockMatcher(ctrl)

		if mock1 == mock2 {
			t.Error("Multiple invocations returned the same instance")
		}
		if mock1.ctrl != mock2.ctrl {
			t.Error("Multiple invocations have different controllers")
		}
	})

	t.Run("Verify Recorder Initialization", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mock := OldMockMatcher(ctrl)

		if mock.recorder == nil {
			t.Error("Recorder is nil")
		} else {
			recorderType := reflect.TypeOf(mock.recorder)
			expectedType := reflect.TypeOf((*MockMatcherMockRecorder)(nil))
			if recorderType != expectedType {
				t.Errorf("Recorder type = %v, want %v", recorderType, expectedType)
			}
		}
	})
}

