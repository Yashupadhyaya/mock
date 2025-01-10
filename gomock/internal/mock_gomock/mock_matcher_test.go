package mock_gomock

import (
	"reflect"
	"testing"
	"github.com/golang/mock/gomock"
)








/*
ROOST_METHOD_HASH=NewMockMatcher_2a2968746f
ROOST_METHOD_SIG_HASH=NewMockMatcher_f53d5470ec

FUNCTION_DEF=func NewMockMatcher(ctrl *go_mock.Controller) *MockMatcher 

 */
func TestNewMockMatcher(t *testing.T) {
	tests := []struct {
		name     string
		ctrl     *go_mock.Controller
		wantType reflect.Type
		wantNil  bool
	}{
		{
			name:     "Create MockMatcher with valid Controller",
			ctrl:     go_mock.NewController(t),
			wantType: reflect.TypeOf(&MockMatcher{}),
			wantNil:  false,
		},
		{
			name:     "Create MockMatcher with nil Controller",
			ctrl:     nil,
			wantType: nil,
			wantNil:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewMockMatcher(tt.ctrl)

			if (got == nil) != tt.wantNil {
				t.Errorf("NewMockMatcher() returned nil: %v, want nil: %v", got == nil, tt.wantNil)
			}

			if !tt.wantNil {
				if reflect.TypeOf(got) != tt.wantType {
					t.Errorf("NewMockMatcher() returned type %v, want %v", reflect.TypeOf(got), tt.wantType)
				}

				if got.ctrl != tt.ctrl {
					t.Errorf("NewMockMatcher().ctrl = %v, want %v", got.ctrl, tt.ctrl)
				}

				if got.recorder == nil {
					t.Error("NewMockMatcher().recorder is nil")
				}

				if reflect.TypeOf(got.recorder) != reflect.TypeOf(&MockMatcherMockRecorder{}) {
					t.Errorf("NewMockMatcher().recorder type = %v, want %v", reflect.TypeOf(got.recorder), reflect.TypeOf(&MockMatcherMockRecorder{}))
				}

				if got.recorder.mock != got {
					t.Errorf("NewMockMatcher().recorder.mock = %v, want %v", got.recorder.mock, got)
				}
			}
		})
	}

	t.Run("Create multiple MockMatchers with same Controller", func(t *testing.T) {
		ctrl := go_mock.NewController(t)
		mock1 := NewMockMatcher(ctrl)
		mock2 := NewMockMatcher(ctrl)

		if mock1 == mock2 {
			t.Error("NewMockMatcher() returned the same instance for different calls")
		}

		if mock1.ctrl != mock2.ctrl {
			t.Error("NewMockMatcher() created MockMatchers with different Controllers")
		}
	})
}

