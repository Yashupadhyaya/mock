package mock_gomock

import (
	"reflect"
	"sync"
	"testing"
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
		expected *MockMatcherMockRecorder
	}{
		{
			name: "Basic Functionality Test",
			setup: func() *MockMatcher {
				ctrl := go_mock.NewController(t)
				recorder := &MockMatcherMockRecorder{}
				return &MockMatcher{ctrl: ctrl, recorder: recorder}
			},
			expected: &MockMatcherMockRecorder{},
		},
		{
			name: "Nil Recorder Test",
			setup: func() *MockMatcher {
				ctrl := go_mock.NewController(t)
				return &MockMatcher{ctrl: ctrl, recorder: nil}
			},
			expected: nil,
		},
		{
			name: "Multiple Calls Consistency Test",
			setup: func() *MockMatcher {
				ctrl := go_mock.NewController(t)
				recorder := &MockMatcherMockRecorder{}
				return &MockMatcher{ctrl: ctrl, recorder: recorder}
			},
			expected: &MockMatcherMockRecorder{},
		},
		{
			name: "Concurrency Safety Test",
			setup: func() *MockMatcher {
				ctrl := go_mock.NewController(t)
				recorder := &MockMatcherMockRecorder{}
				return &MockMatcher{ctrl: ctrl, recorder: recorder}
			},
			expected: &MockMatcherMockRecorder{},
		},
		{
			name: "Integration with MockMatcherMockRecorder Test",
			setup: func() *MockMatcher {
				ctrl := go_mock.NewController(t)
				recorder := &MockMatcherMockRecorder{}
				return &MockMatcher{ctrl: ctrl, recorder: recorder}
			},
			expected: &MockMatcherMockRecorder{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := tt.setup()

			switch tt.name {
			case "Multiple Calls Consistency Test":
				first := m.EXPECT()
				second := m.EXPECT()
				third := m.EXPECT()

				if !reflect.DeepEqual(first, second) || !reflect.DeepEqual(second, third) {
					t.Errorf("Multiple calls to EXPECT() returned different instances")
				}

			case "Concurrency Safety Test":
				var wg sync.WaitGroup
				results := make([]*MockMatcherMockRecorder, 100)

				for i := 0; i < 100; i++ {
					wg.Add(1)
					go func(index int) {
						defer wg.Done()
						results[index] = m.EXPECT()
					}(i)
				}

				wg.Wait()

				for i := 1; i < len(results); i++ {
					if !reflect.DeepEqual(results[0], results[i]) {
						t.Errorf("Concurrent calls to EXPECT() returned different instances")
						break
					}
				}

			default:
				result := m.EXPECT()

				if !reflect.DeepEqual(result, tt.expected) {
					t.Errorf("EXPECT() = %v, want %v", result, tt.expected)
				}
			}
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
			name: "Create a new MockMatcher with a valid Controller",
			ctrl: go_mock.NewController(t),
		},
		{
			name:      "Test with a nil Controller",
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

			if _, ok := got.recorder.(*MockMatcherMockRecorder); !ok {
				t.Errorf("NewMockMatcher().recorder is not of type *MockMatcherMockRecorder")
			}

			if got.recorder.(*MockMatcherMockRecorder).mock != got {
				t.Errorf("NewMockMatcher().recorder.mock does not point back to the MockMatcher")
			}
		})
	}

	t.Run("Verify multiple calls return unique instances", func(t *testing.T) {
		ctrl := go_mock.NewController(t)
		mock1 := NewMockMatcher(ctrl)
		mock2 := NewMockMatcher(ctrl)

		if mock1 == mock2 {
			t.Errorf("NewMockMatcher() returned the same instance for multiple calls")
		}
	})

	t.Run("Check consistency of Controller assignment", func(t *testing.T) {
		ctrl := go_mock.NewController(t)
		mock := NewMockMatcher(ctrl)

		if !reflect.DeepEqual(mock.ctrl, ctrl) {
			t.Errorf("NewMockMatcher() assigned Controller is not the same as the input")
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
			name:     "Basic Matching with a String Argument",
			arg:      "test string",
			expected: true,
			setup: func(m *MockMatcher) {
				m.EXPECT().Matches("test string").Return(true)
			},
		},
		{
			name:     "Matching with an Integer Argument",
			arg:      42,
			expected: true,
			setup: func(m *MockMatcher) {
				m.EXPECT().Matches(42).Return(true)
			},
		},
		{
			name:     "Matching with a Nil Argument",
			arg:      nil,
			expected: false,
			setup: func(m *MockMatcher) {
				m.EXPECT().Matches(nil).Return(false)
			},
		},
		{
			name: "Matching with a Complex Struct",
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
			name:     "Matching with an Empty Interface Argument",
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
func TestMockMatcherString(t *testing.T) {
	tests := []struct {
		name     string
		setup    func(*go_mock.Controller) *MockMatcher
		expected string
		wantErr  bool
	}{
		{
			name: "Basic String Representation",
			setup: func(ctrl *go_mock.Controller) *MockMatcher {
				m := NewMockMatcher(ctrl)
				m.EXPECT().String().Return("MockMatcher")
				return m
			},
			expected: "MockMatcher",
		},
		{
			name: "Consistent String Output",
			setup: func(ctrl *go_mock.Controller) *MockMatcher {
				m := NewMockMatcher(ctrl)
				m.EXPECT().String().Return("Consistent").Times(2)
				return m
			},
			expected: "Consistent",
		},
		{
			name: "Mocked String Return Value",
			setup: func(ctrl *go_mock.Controller) *MockMatcher {
				m := NewMockMatcher(ctrl)
				m.EXPECT().String().Return("Mocked Value")
				return m
			},
			expected: "Mocked Value",
		},
		{
			name: "Panic Handling",
			setup: func(ctrl *go_mock.Controller) *MockMatcher {
				m := NewMockMatcher(ctrl)
				m.EXPECT().String().Do(func() { panic("Simulated panic") })
				return m
			},
			wantErr: true,
		},
		{
			name: "Multiple Invocations Count",
			setup: func(ctrl *go_mock.Controller) *MockMatcher {
				m := NewMockMatcher(ctrl)
				m.EXPECT().String().Return("Multiple").Times(3)
				return m
			},
			expected: "Multiple",
		},
		{
			name: "Nil Controller Handling",
			setup: func(ctrl *go_mock.Controller) *MockMatcher {
				return &MockMatcher{ctrl: nil}
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := go_mock.NewController(t)
			defer ctrl.Finish()

			m := tt.setup(ctrl)

			if tt.wantErr {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("Expected panic, but got none")
					}
				}()
			}

			result := m.String()

			if !tt.wantErr {
				if result != tt.expected {
					t.Errorf("Expected %q, but got %q", tt.expected, result)
				}
			}

			if tt.name == "Multiple Invocations Count" {
				for i := 0; i < 2; i++ {
					result = m.String()
					if result != tt.expected {
						t.Errorf("Expected %q, but got %q on call %d", tt.expected, result, i+2)
					}
				}
			}
		})
	}
}

