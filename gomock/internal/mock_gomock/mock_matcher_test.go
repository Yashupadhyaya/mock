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
		name     string
		ctrl     *go_mock.Controller
		wantNil  bool
		checkPtr bool
	}{
		{
			name:     "Create MockMatcher with valid Controller",
			ctrl:     go_mock.NewController(t),
			wantNil:  false,
			checkPtr: false,
		},
		{
			name:     "Create MockMatcher with nil Controller",
			ctrl:     nil,
			wantNil:  false,
			checkPtr: false,
		},
		{
			name:     "Create multiple MockMatchers with same Controller",
			ctrl:     go_mock.NewController(t),
			wantNil:  false,
			checkPtr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewMockMatcher(tt.ctrl)

			if (got == nil) != tt.wantNil {
				t.Errorf("NewMockMatcher() returned nil: %v, want nil: %v", got == nil, tt.wantNil)
			}

			if !tt.wantNil {
				if got.ctrl != tt.ctrl {
					t.Errorf("NewMockMatcher().ctrl = %v, want %v", got.ctrl, tt.ctrl)
				}

				if got.recorder == nil {
					t.Error("NewMockMatcher().recorder is nil")
				} else {
					if reflect.TypeOf(got.recorder) != reflect.TypeOf(&MockMatcherMockRecorder{}) {
						t.Errorf("NewMockMatcher().recorder type = %v, want *MockMatcherMockRecorder", reflect.TypeOf(got.recorder))
					}

					if got.recorder.mock != got {
						t.Error("NewMockMatcher().recorder.mock does not point back to MockMatcher")
					}
				}

				if tt.checkPtr {
					got2 := NewMockMatcher(tt.ctrl)
					if got == got2 {
						t.Error("Multiple calls to NewMockMatcher() returned the same instance")
					}
					if got.ctrl != got2.ctrl {
						t.Error("Multiple MockMatchers have different Controllers")
					}
				}
			}
		})
	}
}

