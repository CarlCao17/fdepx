package util

import (
	"reflect"
	"testing"
)

func TestNewStringSet(t *testing.T) {
	type args struct {
		ms []string
	}
	tests := []struct {
		name string
		args args
		want StringSet
	}{
		{
			"empty",
			args{
				nil,
			},
			StringSet{},
		},
		{
			"3 case",
			args{
				[]string{"1", "2", "3", "2", "1"},
			},
			StringSet{
				"1": struct{}{}, "2": struct{}{}, "3": struct{}{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewStringSet(tt.args.ms...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStringSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringSet_Add(t *testing.T) {
	type args struct {
		ms []string
	}
	tests := []struct {
		name string
		s    StringSet
		args args
		want StringSet
	}{
		{
			"empty set add some one",
			StringSet{},
			args{
				[]string{"1"},
			},
			StringSet{
				"1": struct{}{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.s.Add(tt.args.ms...); !reflect.DeepEqual(tt.s, tt.want) {
				t.Errorf("StringSet{}")
			}
		})
	}
}

func TestStringSet_Del(t *testing.T) {
	type args struct {
		member string
	}
	tests := []struct {
		name string
		s    StringSet
		args args
		want bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Del(tt.args.member); got != tt.want {
				t.Errorf("Del() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringSet_Diff(t *testing.T) {
	type args struct {
		o StringSet
	}
	tests := []struct {
		name string
		s    StringSet
		args args
		want StringSet
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Diff(tt.args.o); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Diff() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringSet_DiffWith(t *testing.T) {
	type args struct {
		ms []string
	}
	tests := []struct {
		name string
		s    StringSet
		args args
		want StringSet
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.DiffWith(tt.args.ms...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DiffWith() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringSet_Union(t *testing.T) {
	type args struct {
		o StringSet
	}
	tests := []struct {
		name string
		s    StringSet
		args args
		want StringSet
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Union(tt.args.o); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Union() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringSet_UnionWith(t *testing.T) {
	type args struct {
		ms []string
	}
	tests := []struct {
		name string
		s    StringSet
		args args
		want StringSet
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.UnionWith(tt.args.ms...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnionWith() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringSet_String(t *testing.T) {
	tests := []struct {
		name string
		s    StringSet
		want string
	}{
		{
			"nil",
			nil,
			"StringSet{}",
		},
		{
			"empty",
			StringSet{},
			"StringSet{}",
		},
		{
			"common",
			NewStringSet("1", "2", "3"),
			"StringSet{ 1, 2, 3}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
