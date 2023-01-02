package generate_docs

import "testing"

func TestDataResourceLinkBuilder(t *testing.T) {
	type args struct {
		s string
		r string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DataResourceLinkBuilder(tt.args.s, tt.args.r); got != tt.want {
				t.Errorf("DataResourceLinkBuilder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProviderLinkBuilder(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ProviderLinkBuilder(tt.args.s); got != tt.want {
				t.Errorf("ProviderLinkBuilder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResourceLinkBuilder(t *testing.T) {
	type args struct {
		s string
		r string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ResourceLinkBuilder(tt.args.s, tt.args.r); got != tt.want {
				t.Errorf("ResourceLinkBuilder() = %v, want %v", got, tt.want)
			}
		})
	}
}
