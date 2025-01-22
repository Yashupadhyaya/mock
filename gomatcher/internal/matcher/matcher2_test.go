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
		checkCtl bool
	}{
		{
			name:     "Create MockMatcher with valid Controller",
			ctrl:     gomock.NewController(t),
			wantNil:  false,
			checkCtl: true,
		},
		{
			name:     "Test with nil Controller",
			ctrl:     nil,
			wantNil:  true,
			checkCtl: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := OldMockMatcher(tt.ctrl)

			if (result == nil) != tt.wantNil {
				t.Errorf("OldMockMatcher() returned nil: %v, want nil: %v", result == nil, tt.wantNil)
			}

			if !tt.wantNil {
				if result.recorder == nil {
					t.Error("OldMockMatcher() returned MockMatcher with nil recorder")
				}

				if _, ok := result.recorder.(*MockMatcherMockRecorder); !ok {
					t.Error("OldMockMatcher() returned MockMatcher with incorrect recorder type")
				}

				if tt.checkCtl && result.ctrl != tt.ctrl {
					t.Error("OldMockMatcher() returned MockMatcher with incorrect Controller")
				}

				if result.recorder.(*MockMatcherMockRecorder).mock != result {
					t.Error("OldMockMatcher() returned MockMatcher with incorrect recorder.mock reference")
				}
			}
		})
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
		t.Error("OldMockMatcher() returned MockMatchers with different Controllers")
	}
}

func TestOldMockMatcherRecorderInitialization(t *testing.T) {
	ctrl := gomock.NewController(t)
	result := OldMockMatcher(ctrl)

	if result.recorder == nil {
		t.Fatal("OldMockMatcher() returned MockMatcher with nil recorder")
	}

	recorderType := reflect.TypeOf(result.recorder)
	expectedType := reflect.TypeOf((*MockMatcherMockRecorder)(nil))

	if recorderType != expectedType {
		t.Errorf("OldMockMatcher() returned MockMatcher with incorrect recorder type. Got %v, want %v", recorderType, expectedType)
	}
}

