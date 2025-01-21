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
		wantType reflect.Type
	}{
		{
			name:     "Create MockMatcher with valid Controller",
			ctrl:     gomock.NewController(t),
			wantNil:  false,
			wantType: reflect.TypeOf(&MockMatcher{}),
		},
		{
			name:     "Create MockMatcher with nil Controller",
			ctrl:     nil,
			wantNil:  true,
			wantType: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := OldMockMatcher(tt.ctrl)

			if (got == nil) != tt.wantNil {
				t.Errorf("OldMockMatcher() returned nil: %v, want nil: %v", got == nil, tt.wantNil)
			}

			if !tt.wantNil {
				if got.ctrl != tt.ctrl {
					t.Errorf("OldMockMatcher() ctrl = %v, want %v", got.ctrl, tt.ctrl)
				}

				if got.recorder == nil {
					t.Error("OldMockMatcher() recorder is nil")
				}

				if reflect.TypeOf(got) != tt.wantType {
					t.Errorf("OldMockMatcher() returned type %v, want %v", reflect.TypeOf(got), tt.wantType)
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

	t.Run("Check Controller assignment consistency", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mock := OldMockMatcher(ctrl)

		if mock.ctrl != ctrl {
			t.Errorf("OldMockMatcher() assigned controller %v, want %v", mock.ctrl, ctrl)
		}
	})
}

