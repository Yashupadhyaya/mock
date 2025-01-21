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
		wantType reflect.Type
		wantNil  bool
		wantErr  bool
	}{
		{
			name:     "Create MockMatcher with valid Controller",
			ctrl:     gomock.NewController(t),
			wantType: reflect.TypeOf(&MockMatcher{}),
			wantNil:  false,
			wantErr:  false,
		},
		{
			name:     "Test with nil Controller",
			ctrl:     nil,
			wantType: nil,
			wantNil:  true,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got *MockMatcher
			var panicked bool

			func() {
				defer func() {
					if r := recover(); r != nil {
						panicked = true
					}
				}()
				got = OldMockMatcher(tt.ctrl)
			}()

			if panicked != tt.wantErr {
				t.Errorf("OldMockMatcher() panic = %v, wantErr %v", panicked, tt.wantErr)
				return
			}

			if (got == nil) != tt.wantNil {
				t.Errorf("OldMockMatcher() returned nil: %v, want nil: %v", got == nil, tt.wantNil)
				return
			}

			if !tt.wantNil {
				if reflect.TypeOf(got) != tt.wantType {
					t.Errorf("OldMockMatcher() returned type = %v, want %v", reflect.TypeOf(got), tt.wantType)
				}

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

