package mock_gomock

import (
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestNewMockMatcher(t *testing.T) {
	tests := []struct {
		name     string
		ctrl     *gomock.Controller
		wantNil  bool
		wantPanic bool
	}{
		{
			name:     "Create MockMatcher with valid Controller",
			ctrl:     gomock.NewController(t),
			wantNil:  false,
			wantPanic: false,
		},
		{
			name:     "Create MockMatcher with nil Controller",
			ctrl:     nil,
			wantNil:  true,
			wantPanic: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantPanic {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("NewMockMatcher() did not panic with nil Controller")
					}
				}()
			}

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
			}
		})
	}

	// Test multiple calls to NewMockMatcher
	t.Run("Multiple calls to NewMockMatcher", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mock1 := NewMockMatcher(ctrl)
		mock2 := NewMockMatcher(ctrl)

		if mock1 == mock2 {
			t.Error("Multiple calls to NewMockMatcher returned the same instance")
		}

		if mock1.ctrl != mock2.ctrl {
			t.Error("MockMatchers created with the same Controller have different ctrl fields")
		}
	})
}
