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
					t.Errorf("OldMockMatcher() ctrl = %v, want %v", result.ctrl, tt.ctrl)
				}

				if result.recorder == nil {
					t.Error("OldMockMatcher() recorder is nil")
				} else {
					if result.recorder.mock != result {
						t.Error("OldMockMatcher() recorder.mock does not point back to MockMatcher")
					}
				}
			}
		})
	}
}

func TestOldMockMatcherControllerAssignment(t *testing.T) {
	type customTestHelper struct {
		gomock.TestHelper
	}
	customHelper := &customTestHelper{}
	ctrl := gomock.NewController(customHelper)

	result := OldMockMatcher(ctrl)

	if !reflect.DeepEqual(result.ctrl, ctrl) {
		t.Error("OldMockMatcher() did not correctly assign the provided controller")
	}
}

func TestOldMockMatcherMultipleCalls(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock1 := OldMockMatcher(ctrl)
	mock2 := OldMockMatcher(ctrl)

	if mock1 == mock2 {
		t.Error("OldMockMatcher() returned the same instance for multiple calls")
	}

	if mock1.ctrl != mock2.ctrl {
		t.Error("OldMockMatcher() returned mocks with different controllers")
	}
}

