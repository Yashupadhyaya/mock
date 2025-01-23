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
			name:    "Create MockMatcher with valid Controller",
			ctrl:    gomock.NewController(t),
			wantNil: false,
		},
		{
			name:    "Create MockMatcher with nil Controller",
			ctrl:    nil,
			wantNil: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var mock *MockMatcher
			defer func() {
				if r := recover(); r != nil {
					if !tt.wantPanic {
						t.Errorf("OldMockMatcher() unexpected panic: %v", r)
					}
				}
			}()

			mock = OldMockMatcher(tt.ctrl)

			if (mock == nil) != tt.wantNil {
				t.Errorf("OldMockMatcher() returned nil: %v, want nil: %v", mock == nil, tt.wantNil)
			}

			if !tt.wantNil {
				if mock.ctrl != tt.ctrl {
					t.Errorf("OldMockMatcher() ctrl = %v, want %v", mock.ctrl, tt.ctrl)
				}

				if mock.recorder == nil {
					t.Error("OldMockMatcher() recorder is nil")
				} else {
					if mock.recorder.mock != mock {
						t.Error("OldMockMatcher() recorder.mock does not point back to mock")
					}
				}

				if _, ok := interface{}(mock.recorder).(*MockMatcherMockRecorder); !ok {
					t.Errorf("OldMockMatcher() recorder is not of type *MockMatcherMockRecorder")
				}
			}
		})
	}

	t.Run("Verify uniqueness of created MockMatchers", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mock1 := OldMockMatcher(ctrl)
		mock2 := OldMockMatcher(ctrl)

		if mock1 == mock2 {
			t.Error("OldMockMatcher() returned the same instance for multiple calls")
		}

		if mock1.ctrl != mock2.ctrl {
			t.Error("OldMockMatcher() created MockMatchers with different Controllers")
		}
	})

	t.Run("Test Controller association", func(t *testing.T) {
		ctrl1 := gomock.NewController(t)
		ctrl2 := gomock.NewController(t)

		mock1 := OldMockMatcher(ctrl1)
		mock2 := OldMockMatcher(ctrl2)

		if mock1.ctrl != ctrl1 || mock2.ctrl != ctrl2 {
			t.Error("OldMockMatcher() created MockMatchers with incorrect Controller association")
		}
	})

	t.Run("Verify recorder type", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mock := OldMockMatcher(ctrl)

		recorderType := reflect.TypeOf(mock.recorder)
		expectedType := reflect.TypeOf((*MockMatcherMockRecorder)(nil))

		if recorderType != expectedType {
			t.Errorf("OldMockMatcher() recorder type = %v, want %v", recorderType, expectedType)
		}
	})
}

