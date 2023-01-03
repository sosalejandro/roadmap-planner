package entities

import (
	"reflect"
	"testing"
)

func TestIdeaPlanner_AddToList(t *testing.T) {

	for _, tt := range ideaPlannerPromoteToSublistTests {
		t.Run(tt.name, func(t *testing.T) {
			ip := tt.fields.list
			if err := ip.AddToList(tt.args.idea); (err != nil) != tt.wantErr {
				t.Errorf("AddToList() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestIdeaPlanner_InsertAtList(t *testing.T) {
	tests := []struct {
		name    string
		fields  ideaFields
		args    ideaArgs
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ip := tt.fields.list
			if err := ip.InsertAtList(tt.args.index, tt.args.idea); (err != nil) != tt.wantErr {
				t.Errorf("InsertAtList() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestListManager_PromoteToSubList(t *testing.T) {
	for _, tt := range ideaPlannerPromoteToSublistTests {
		t.Run(tt.name, func(t *testing.T) {
			ip := tt.fields.list
			got, err := ip.PromoteToSublist(tt.args.index)
			if (err != nil) != tt.wantErr {
				t.Errorf("promoteToSubList(%d) error = %v, wantErr %v", tt.args.index, err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("promoteToSubList(%d) got = %v, want %v", tt.args.index, got, tt.want)
			}
		})
	}
}
