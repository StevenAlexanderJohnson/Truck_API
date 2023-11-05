package Authentication_Handler

import "testing"

func TestCheck_Roles(t *testing.T) {
	tests := []struct {
		testName       string
		roles          []string
		required_roles []string
		want           bool
		willFail       bool
	}{
		{
			testName:       "Valid Roles",
			roles:          []string{"admin", "user"},
			required_roles: []string{"admin", "user"},
			want:           true,
			willFail:       false,
		},
		{
			testName:       "Invalid Roles",
			roles:          []string{"admin", "user"},
			required_roles: []string{"admin", "user", "invalid"},
			want:           false,
			willFail:       true,
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			got := Check_Roles(test.roles, test.required_roles)
			if got != test.want {
				if test.willFail {
					return
				}
				t.Errorf("Check_Roles() = %v, want %v", got, test.want)
			}
		})
	}
}
