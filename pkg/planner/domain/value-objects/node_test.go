package value_objects

import (
	"reflect"
	"testing"
)

func TestNewNode(t *testing.T) {
	type args struct {
		name string
	}

	tests := []struct {
		name string
		args args
		want *Node
	}{
		{
			name: "Creates new node called Patricia",
			args: args{name: "Patricia"},
			want: &Node{name: "Patricia"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNode(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_String(t *testing.T) {
	type fields struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "Node is called Patricia",
			fields: fields{name: "Patricia"},
			want:   "Patricia",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &Node{
				name: tt.fields.name,
			}
			if got := n.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
