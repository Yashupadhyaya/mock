package mock_gomock

import (
	"reflect"
	"sync"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestMockMatcherExpect(t *testing.T) {
	tests := []struct {
		name     string
		setup    func() *MockMatcher
		expected *MockMatcherMockRecorder
	}{
		{
			name: "Basic Functionality Test",
			setup: func() *MockMatcher {
				ctrl := gomock.NewController(t)
				recorder := &MockMatcherMockRecorder{}
				return &MockMatcher{ctrl: ctrl, recorder: recorder}
			},
			expected: &MockMatcherMockRecorder{},
		},
		{
			name: "Nil Recorder Test",
			setup: func() *MockMatcher {
				ctrl := gomock.NewController(t)
				return &MockMatcher{ctrl: ctrl, recorder: nil}
			},
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := tt.setup()
			result := m.EXPECT()

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("EXPECT() = %v, want %v", result, tt.expected)
			}
		})
	}

	t.Run("Multiple Calls Consistency Test", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		recorder := &MockMatcherMockRecorder{}
		m := &MockMatcher{ctrl: ctrl, recorder: recorder}

		result1 := m.EXPECT()
		result2 := m.EXPECT()
		result3 := m.EXPECT()

		if result1 != result2 || result2 != result3 {
			t.Errorf("Multiple calls to EXPECT() returned different recorders")
		}
	})

	t.Run("Concurrency Safety Test", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		recorder := &MockMatcherMockRecorder{}
		m := &MockMatcher{ctrl: ctrl, recorder: recorder}

		const numGoroutines = 100
		var wg sync.WaitGroup
		results := make([]*MockMatcherMockRecorder, numGoroutines)

		for i := 0; i < numGoroutines; i++ {
			wg.Add(1)
			go func(index int) {
				defer wg.Done()
				results[index] = m.EXPECT()
			}(i)
		}

		wg.Wait()

		for i := 1; i < numGoroutines; i++ {
			if results[i] != results[0] {
				t.Errorf("Concurrent calls to EXPECT() returned different recorders")
				break
			}
		}
	})

	t.Run("Integration with Controller Test", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m := &MockMatcher{ctrl: ctrl, recorder: &MockMatcherMockRecorder{}}
		recorder := m.EXPECT()

		// Set up an expectation
		recorder.Matches(gomock.Any()).Return(true)

		// Verify that the expectation is registered with the controller
		if !m.Matches("test") {
			t.Errorf("Expectation set through EXPECT() was not registered correctly")
		}
	})
}
