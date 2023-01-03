package entities

import (
	"fmt"
	value_objects "roadmap-planner/pkg/planner/domain/value-objects"
)

type ideaFields struct {
	list ListManager
}
type ideaArgs struct {
	index int
	idea  value_objects.BaseNode
}

var plannerTest = Planner(&List{
	Node:  value_objects.NewNode("To do list"),
	ideas: ideasList,
})

var plannerTest2 = func() Planner {
	pt2 := Planner(&Sublist{List{
		Node:  value_objects.NewNode("To do list"),
		ideas: append(make([]value_objects.BaseNode, 0), ideasList...),
	}})
	ip := &IdeaPlanner{&pt2}
	ip.PromoteToSublist(20)
	return pt2
}()

var ideaPlannerPromoteToSublistTests = []struct {
	name    string
	fields  ideaFields
	args    ideaArgs
	want    *Sublist
	wantErr bool
}{
	// TODO: Add test cases.
	{
		name: "IdeaPlanner.List creates a sublist out of an idea",
		fields: ideaFields{
			list: &IdeaPlanner{list: &plannerTest},
		},
		args: ideaArgs{
			index: 19,
			idea:  nil,
		},
		want:    NewSublist("Add OAuth2"),
		wantErr: false,
	},
	{
		name: fmt.Sprint("IdeaPlanner.Sublist: index 20 should be already a sublist thus ",
			"it doesn't create a new one returning the existing entry"),
		fields: ideaFields{
			list: &IdeaPlanner{list: &plannerTest2},
		},
		args: ideaArgs{
			index: 20,
			idea:  nil,
		},
		want:    NewSublist("Add Google Calendar integration"),
		wantErr: false,
	},
}
