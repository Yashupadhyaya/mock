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
			if tt.wantPanic {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("NewMockMatcher() did not panic as expected")
					}
				}()
			}

			got := NewMockMatcher(tt.ctrl)

			if tt.wantNil {
				if got != nil {
					t.Errorf("NewMockMatcher() = %v, want nil", got)
				}
				return
			}

			if got == nil {
				t.Fatalf("NewMockMatcher() returned nil, want non-nil")
			}

			if got.ctrl != tt.ctrl {
				t.Errorf("NewMockMatcher().ctrl = %v, want %v", got.ctrl, tt.ctrl)
			}

			if got.recorder == nil {
				t.Errorf("NewMockMatcher().recorder is nil")
			}

			if reflect.TypeOf(got.recorder) != reflect.TypeOf(&MockMatcherMockRecorder{}) {
				t.Errorf("NewMockMatcher().recorder type = %v, want *MockMatcherMockRecorder", reflect.TypeOf(got.recorder))
			}
		})
	}
}

func TestNewMockMatcherMultiple(t *testing.T) {
	ctrl := go_mock.NewController(t)

	mock1 := NewMockMatcher(ctrl)
	mock2 := NewMockMatcher(ctrl)

	if mock1 == nil || mock2 == nil {
		t.Fatalf("NewMockMatcher() returned nil")
	}

	if mock1 == mock2 {
		t.Errorf("NewMockMatcher() returned the same instance for different calls")
	}

	if mock1.ctrl != mock2.ctrl {
		t.Errorf("NewMockMatcher() set different controllers for mocks")
	}

	if mock1.recorder == mock2.recorder {
		t.Errorf("NewMockMatcher() set the same recorder for different mocks")
	}
}

