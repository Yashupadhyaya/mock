package mock_gomock

import (
	"reflect"
	"testing"
	go_mock "github.com/golang/mock/gomock"
)









/*
ROOST_METHOD_HASH=NewMockMatcher_2a2968746f
ROOST_METHOD_SIG_HASH=NewMockMatcher_f53d5470ec

FUNCTION_DEF=func NewMockMatcher(ctrl *go_mock.Controller) *MockMatcher 

 */
func TestNewMockMatcher(t *testing.T) {
	tests := []struct {
		name      string
		ctrl      *go_mock.Controller
		wantNil   bool
		wantPanic bool
	}{
		{
			name:      "Create MockMatcher with valid Controller",
			ctrl:      go_mock.NewController(t),
			wantNil:   false,
			wantPanic: false,
		},
		{
			name:      "Create MockMatcher with nil Controller",
			ctrl:      nil,
			wantNil:   true,
			wantPanic: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var mock *MockMatcher
			defer func() {
				r := recover()
				if (r != nil) != tt.wantPanic {
					t.Errorf("NewMockMatcher() panic = %v, wantPanic %v", r, tt.wantPanic)
				}
			}()

			mock = NewMockMatcher(tt.ctrl)

			if (mock == nil) != tt.wantNil {
				t.Errorf("NewMockMatcher() returned nil: %v, want nil: %v", mock == nil, tt.wantNil)
				return
			}

			if !tt.wantNil {
				if mock.ctrl != tt.ctrl {
					t.Errorf("NewMockMatcher().ctrl = %v, want %v", mock.ctrl, tt.ctrl)
				}

				if mock.recorder == nil {
					t.Error("NewMockMatcher().recorder is nil")
				} else {
					if reflect.TypeOf(mock.recorder) != reflect.TypeOf(&MockMatcherMockRecorder{}) {
						t.Errorf("NewMockMatcher().recorder type = %v, want *MockMatcherMockRecorder", reflect.TypeOf(mock.recorder))
					}

					if mock.recorder.mock != mock {
						t.Error("NewMockMatcher().recorder.mock does not point back to MockMatcher")
					}
				}
			}
		})
	}

	t.Run("Create multiple MockMatchers", func(t *testing.T) {
		ctrl := go_mock.NewController(t)
		mock1 := NewMockMatcher(ctrl)
		mock2 := NewMockMatcher(ctrl)

		if mock1 == mock2 {
			t.Error("NewMockMatcher() returned the same instance for multiple calls")
		}

		if mock1.ctrl != mock2.ctrl {
			t.Error("NewMockMatcher() did not use the same Controller for multiple calls")
		}
	})
}

