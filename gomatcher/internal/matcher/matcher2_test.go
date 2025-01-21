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
		wantType  reflect.Type
		wantNil   bool
		wantPanic bool
	}{
		{
			name:     "Create MockMatcher with valid Controller",
			ctrl:     gomock.NewController(t),
			wantType: reflect.TypeOf(&MockMatcher{}),
			wantNil:  false,
		},
		{
			name:     "Test with nil Controller",
			ctrl:     nil,
			wantType: nil,
			wantNil:  true,
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
			}

			if tt.wantType != nil && reflect.TypeOf(got) != tt.wantType {
				t.Errorf("OldMockMatcher() returned type = %v, want %v", reflect.TypeOf(got), tt.wantType)
			}

			if got != nil {
				if got.ctrl != tt.ctrl {
					t.Errorf("OldMockMatcher() ctrl = %v, want %v", got.ctrl, tt.ctrl)
				}

				if got.recorder == nil {
					t.Error("OldMockMatcher() recorder is nil")
				}

				if _, ok := got.recorder.(*MockMatcherMockRecorder); !ok {
					t.Errorf("OldMockMatcher() recorder type = %T, want *MockMatcherMockRecorder", got.recorder)
				}
			}
		})
	}

	t.Run("Verify uniqueness of created MockMatcher instances", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mock1 := OldMockMatcher(ctrl)
		mock2 := OldMockMatcher(ctrl)

		if mock1 == mock2 {
			t.Error("OldMockMatcher() returned the same instance for multiple calls")
		}
	})
}

