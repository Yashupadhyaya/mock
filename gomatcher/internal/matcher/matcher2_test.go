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
			name:    "Multiple Calls to OldMockMatcher",
			ctrl:    gomock.NewController(t),
			wantNil: false,
		},
		{
			name:      "Nil Controller Handling",
			ctrl:      nil,
			wantNil:   true,
			wantPanic: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantPanic {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("OldMockMatcher() did not panic with nil controller")
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
					t.Errorf("OldMockMatcher().recorder is nil")
				} else if got.recorder.mock != got {
					t.Errorf("OldMockMatcher().recorder.mock does not point back to MockMatcher")
				}
			}
		})
	}

	t.Run("Multiple Calls to OldMockMatcher", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mock1 := OldMockMatcher(ctrl)
		mock2 := OldMockMatcher(ctrl)

		if mock1 == mock2 {
			t.Errorf("Multiple calls to OldMockMatcher() returned the same instance")
		}

		if mock1.ctrl != mock2.ctrl {
			t.Errorf("Multiple calls to OldMockMatcher() returned instances with different controllers")
		}
	})

	t.Run("Controller State Preservation", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		ctrl.RecordCall(ctrl, "SomeMethod", reflect.TypeOf(ctrl))

		mock := OldMockMatcher(ctrl)

		if len(ctrl.ExpectedCalls()) != 1 {
			t.Errorf("Controller state not preserved after OldMockMatcher() call")
		}

		if mock.ctrl != ctrl {
			t.Errorf("OldMockMatcher() did not use the provided controller")
		}
	})
}

