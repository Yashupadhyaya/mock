package mock_gomock

import (
	"github.com/golang/mock/gomock"
)

func OldMockMatcher(ctrl *gomock.Controller) *MockMatcher {
	mock := &MockMatcher{ctrl: ctrl}
	mock.recorder = &MockMatcherMockRecorder{mock}
	return mock
}
g