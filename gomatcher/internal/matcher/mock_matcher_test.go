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
		mock     *MockMatcher
		expected *MockMatcherMockRecorder
	}{
		{
			name: "Basic Functionality Test",
			mock: &MockMatcher{
				recorder: &MockMatcherMockRecorder{},
			},
			expected: &MockMatcherMockRecorder{},
		},
		{
			name: "Nil Recorder Test",
			mock: &MockMatcher{
				recorder: nil,
			},
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.mock.EXPECT()
			assert.Equal(t, tt.expected, result)
		})
	}

	t.Run("Multiple Calls Consistency Test", func(t *testing.T) {
		mock := &MockMatcher{
			recorder: &MockMatcherMockRecorder{},
		}
		first := mock.EXPECT()
		second := mock.EXPECT()
		third := mock.EXPECT()
		assert.Same(t, first, second)
		assert.Same(t, second, third)
	})

	t.Run("Concurrency Safety Test", func(t *testing.T) {
		mock := &MockMatcher{
			recorder: &MockMatcherMockRecorder{},
		}
		const goroutines = 10
		var wg sync.WaitGroup
		results := make([]*MockMatcherMockRecorder, goroutines)

		for i := 0; i < goroutines; i++ {
			wg.Add(1)
			go func(index int) {
				defer wg.Done()
				results[index] = mock.EXPECT()
			}(i)
		}

		wg.Wait()

		for i := 1; i < goroutines; i++ {
			assert.Same(t, results[0], results[i])
		}
	})

	t.Run("Integration with MockMatcherMockRecorder Test", func(t *testing.T) {
		recorder := &MockMatcherMockRecorder{}
		mock := &MockMatcher{
			recorder: recorder,
		}
		result := mock.EXPECT()
		assert.IsType(t, &MockMatcherMockRecorder{}, result)

	})
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
			name:      "Valid Controller",
			ctrl:      go_mock.NewController(t),
			wantNil:   false,
			wantPanic: false,
		},
		{
			name:      "Nil Controller",
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
			}

			if !tt.wantNil {
				if mock.ctrl != tt.ctrl {
					t.Errorf("NewMockMatcher().ctrl = %v, want %v", mock.ctrl, tt.ctrl)
				}

				if mock.recorder == nil {
					t.Error("NewMockMatcher().recorder is nil")
				} else {
					if _, ok := mock.recorder.(*MockMatcherMockRecorder); !ok {
						t.Errorf("NewMockMatcher().recorder type = %v, want *MockMatcherMockRecorder", reflect.TypeOf(mock.recorder))
					}
					if mock.recorder.(*MockMatcherMockRecorder).mock != mock {
						t.Error("NewMockMatcher().recorder.mock does not point back to the MockMatcher")
					}
				}
			}
		})
	}

	t.Run("Multiple calls return unique instances", func(t *testing.T) {
		ctrl := go_mock.NewController(t)
		mock1 := NewMockMatcher(ctrl)
		mock2 := NewMockMatcher(ctrl)

		if mock1 == mock2 {
			t.Error("NewMockMatcher() returned the same instance for multiple calls")
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
				Name string
				Age  int
			}{"John", 30},
			expected: true,
			setup: func(m *MockMatcher) {
				m.EXPECT().Matches(struct {
					Name string
					Age  int
				}{"John", 30}).Return(true)
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
		mockMatcher.EXPECT().Matches(struct{ Value int }{Value: 42}).Return(true)

		if !mockMatcher.Matches("string") {
			t.Errorf("Expected true for string argument")
		}
		if mockMatcher.Matches(123) {
			t.Errorf("Expected false for integer argument")
		}
		if !mockMatcher.Matches(struct{ Value int }{Value: 42}) {
			t.Errorf("Expected true for struct argument")
		}
	})
}


/*
ROOST_METHOD_HASH=String_e973565fd6
ROOST_METHOD_SIG_HASH=String_e75af914b2

FUNCTION_DEF=func (m *MockMatcher) String() string 

 */
func (m *MockMatcher) String() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "String")
	ret0, _ := ret[0].(string)
	return ret0
}

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
				assert.PanicsWithValue(t, tt.panicMessage, func() {
					mockMatcher.String()
				})
			} else {
				for i := 0; i < tt.callCount; i++ {
					result := mockMatcher.String()
					assert.Equal(t, tt.expectedResult, result)
				}
			}
		})
	}
}

