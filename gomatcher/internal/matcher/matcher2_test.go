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
			name:      "Create MockMatcher with nil Controller",
			ctrl:      nil,
			wantPanic: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantPanic {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("OldMockMatcher() did not panic with nil Controller")
					}
				}()
			}

			got := OldMockMatcher(tt.ctrl)

			if (got == nil) != tt.wantNil {
				t.Errorf("OldMockMatcher() returned nil: %v, want nil: %v", got == nil, tt.wantNil)
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
						t.Errorf("OldMockMatcher().recorder.mock does not point back to the created MockMatcher")
					}
				}
			}
		})
	}

	t.Run("Verify uniqueness of created MockMatchers", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mock1 := OldMockMatcher(ctrl)
		mock2 := OldMockMatcher(ctrl)

		if mock1 == mock2 {
			t.Error("OldMockMatcher() returned the same instance for multiple calls")
		}

		if mock1.ctrl != mock2.ctrl {
			t.Error("OldMockMatcher() returned MockMatchers with different Controllers")
		}
	})

	t.Run("Verify MockMatcher type", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		got := OldMockMatcher(ctrl)

		if reflect.TypeOf(got) != reflect.TypeOf(&MockMatcher{}) {
			t.Errorf("OldMockMatcher() returned type %v, want *MockMatcher", reflect.TypeOf(got))
		}
	})
}

