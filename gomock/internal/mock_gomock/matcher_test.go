package mock_gomock

import (
	"sync"
	"testing"
	"github.com/stretchr/testify/assert"
	"reflect"
	go_mock "github.com/golang/mock/gomock"
)









/*
ROOST_METHOD_HASH=EXPECT_f21a025b11
ROOST_METHOD_SIG_HASH=EXPECT_62b86d5d58

FUNCTION_DEF=func (m *MockMatcher) EXPECT() *MockMatcherMockRecorder 

 */
func TestMockMatcherExpect(t *testing.T) {
	tests := []struct {
		name     string
		setup    func() *MockMatcher
		validate func(*testing.T, *MockMatcher)
	}{
		{
			name: "Basic Functionality Test",
			setup: func() *MockMatcher {
				recorder := &MockMatcherMockRecorder{}
				return &MockMatcher{recorder: recorder}
			},
			validate: func(t *testing.T, m *MockMatcher) {
				result := m.EXPECT()
				assert.Equal(t, m.recorder, result, "EXPECT() should return the recorder field")
			},
		},
		{
			name: "Nil Recorder Test",
			setup: func() *MockMatcher {
				return &MockMatcher{recorder: nil}
			},
			validate: func(t *testing.T, m *MockMatcher) {
				result := m.EXPECT()
				assert.Nil(t, result, "EXPECT() should return nil when recorder is nil")
			},
		},
		{
			name: "Multiple Calls Consistency Test",
			setup: func() *MockMatcher {
				recorder := &MockMatcherMockRecorder{}
				return &MockMatcher{recorder: recorder}
			},
			validate: func(t *testing.T, m *MockMatcher) {
				result1 := m.EXPECT()
				result2 := m.EXPECT()
				result3 := m.EXPECT()
				assert.Equal(t, result1, result2, "Multiple calls to EXPECT() should return the same recorder")
				assert.Equal(t, result2, result3, "Multiple calls to EXPECT() should return the same recorder")
			},
		},
		{
			name: "Concurrency Safety Test",
			setup: func() *MockMatcher {
				recorder := &MockMatcherMockRecorder{}
				return &MockMatcher{recorder: recorder}
			},
			validate: func(t *testing.T, m *MockMatcher) {
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
					assert.Equal(t, results[0], results[i], "Concurrent calls to EXPECT() should return the same recorder")
				}
			},
		},
		{
			name: "Integration with MockMatcherMockRecorder Test",
			setup: func() *MockMatcher {
				recorder := &MockMatcherMockRecorder{}
				return &MockMatcher{recorder: recorder}
			},
			validate: func(t *testing.T, m *MockMatcher) {
				result := m.EXPECT()
				assert.IsType(t, &MockMatcherMockRecorder{}, result, "EXPECT() should return a MockMatcherMockRecorder")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockMatcher := tt.setup()
			tt.validate(t, mockMatcher)
		})
	}
}


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
			name: "Create MockMatcher with valid Controller",
			ctrl: go_mock.NewController(t),
		},
		{
			name:      "Create MockMatcher with nil Controller",
			ctrl:      nil,
			wantNil:   true,
			wantPanic: true,
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

			if _, ok := got.recorder.(*MockMatcherMockRecorder); !ok {
				t.Errorf("NewMockMatcher().recorder is not of type *MockMatcherMockRecorder")
			}

			if got.recorder.(*MockMatcherMockRecorder).mock != got {
				t.Errorf("NewMockMatcher().recorder.mock does not point back to the MockMatcher")
			}
		})
	}

	t.Run("Multiple calls return unique instances", func(t *testing.T) {
		ctrl := go_mock.NewController(t)
		mock1 := NewMockMatcher(ctrl)
		mock2 := NewMockMatcher(ctrl)

		if mock1 == mock2 {
			t.Errorf("NewMockMatcher() returned the same instance for multiple calls")
		}
	})

	t.Run("Consistency of Controller assignment", func(t *testing.T) {
		ctrl := go_mock.NewController(t)
		mock := NewMockMatcher(ctrl)

		if !reflect.DeepEqual(mock.ctrl, ctrl) {
			t.Errorf("NewMockMatcher() assigned a different Controller than provided")
		}
	})
}


/*
ROOST_METHOD_HASH=Matches_221d68ae59
ROOST_METHOD_SIG_HASH=Matches_ce995a8670

FUNCTION_DEF=func (m *MockMatcher) Matches(arg0 interface) bool 

 */
func TestMockMatcherMatches(t *testing.T) {
	tests := []struct {
		name     string
		arg      interface{}
		expected bool
		setup    func(*MockMatcher)
	}{
		{
			name:     "Basic Matching with String Argument",
			arg:      "test string",
			expected: true,
			setup: func(m *MockMatcher) {
				m.EXPECT().Matches("test string").Return(true)
			},
		},
		{
			name:     "Matching with Integer Argument",
			arg:      42,
			expected: true,
			setup: func(m *MockMatcher) {
				m.EXPECT().Matches(42).Return(true)
			},
		},
		{
			name:     "Matching with Nil Argument",
			arg:      nil,
			expected: false,
			setup: func(m *MockMatcher) {
				m.EXPECT().Matches(nil).Return(false)
			},
		},
		{
			name: "Matching with Complex Struct",
			arg: struct {
				Name  string
				Value int
			}{"test", 123},
			expected: true,
			setup: func(m *MockMatcher) {
				m.EXPECT().Matches(struct {
					Name  string
					Value int
				}{"test", 123}).Return(true)
			},
		},
		{
			name:     "Matching with Empty Interface Argument",
			arg:      interface{}(nil),
			expected: true,
			setup: func(m *MockMatcher) {
				m.EXPECT().Matches(interface{}(nil)).Return(true)
			},
		},
		{
			name:     "Error Handling for Unexpected Argument Type",
			arg:      []int{1, 2, 3},
			expected: false,
			setup: func(m *MockMatcher) {
				m.EXPECT().Matches(go_mock.Any()).Return(false)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := go_mock.NewController(t)
			defer ctrl.Finish()

			mockMatcher := NewMockMatcher(ctrl)
			tt.setup(mockMatcher)

			result := mockMatcher.Matches(tt.arg)
			if result != tt.expected {
				t.Errorf("Matches() = %v, want %v", result, tt.expected)
			}
		})
	}

	t.Run("Multiple Consecutive Calls with Different Arguments", func(t *testing.T) {
		ctrl := go_mock.NewController(t)
		defer ctrl.Finish()

		mockMatcher := NewMockMatcher(ctrl)
		mockMatcher.EXPECT().Matches("string").Return(true)
		mockMatcher.EXPECT().Matches(123).Return(false)
		mockMatcher.EXPECT().Matches(struct{ Name string }{"test"}).Return(true)

		if !mockMatcher.Matches("string") {
			t.Errorf("Expected true for string argument")
		}
		if mockMatcher.Matches(123) {
			t.Errorf("Expected false for integer argument")
		}
		if !mockMatcher.Matches(struct{ Name string }{"test"}) {
			t.Errorf("Expected true for struct argument")
		}
	})
}


/*
ROOST_METHOD_HASH=String_e973565fd6
ROOST_METHOD_SIG_HASH=String_e75af914b2

FUNCTION_DEF=func (m *MockMatcher) String() string 

 */
func TestMockMatcherString(t *testing.T) {
	tests := []struct {
		name           string
		setupMock      func(*MockMatcher)
		expectedResult string
		expectPanic    bool
		panicMessage   string
		callCount      int
	}{
		{
			name: "Basic String Representation",
			setupMock: func(m *MockMatcher) {
				m.EXPECT().String().Return("MockMatcher")
			},
			expectedResult: "MockMatcher",
			callCount:      1,
		},
		{
			name: "Consistent String Output",
			setupMock: func(m *MockMatcher) {
				m.EXPECT().String().Return("Consistent").Times(2)
			},
			expectedResult: "Consistent",
			callCount:      2,
		},
		{
			name: "Mocked String Return Value",
			setupMock: func(m *MockMatcher) {
				m.EXPECT().String().Return("Mocked Value")
			},
			expectedResult: "Mocked Value",
			callCount:      1,
		},
		{
			name: "Panic Handling",
			setupMock: func(m *MockMatcher) {
				m.EXPECT().String().Do(func() {
					panic("Simulated panic")
				})
			},
			expectPanic:  true,
			panicMessage: "Simulated panic",
			callCount:    1,
		},
		{
			name: "Multiple Invocations Count",
			setupMock: func(m *MockMatcher) {
				m.EXPECT().String().Return("Multiple").Times(3)
			},
			expectedResult: "Multiple",
			callCount:      3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := go_mock.NewController(t)
			defer ctrl.Finish()

			mockMatcher := NewMockMatcher(ctrl)
			tt.setupMock(mockMatcher)

			if tt.expectPanic {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("Expected panic, but no panic occurred")
					} else if r.(string) != tt.panicMessage {
						t.Errorf("Expected panic message %q, but got %q", tt.panicMessage, r)
					}
				}()
			}

			for i := 0; i < tt.callCount; i++ {
				result := mockMatcher.String()
				if !tt.expectPanic && result != tt.expectedResult {
					t.Errorf("Expected %q, but got %q", tt.expectedResult, result)
				}
			}
		})
	}
}

