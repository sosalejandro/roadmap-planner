package entities

import value_objects "roadmap-planner/pkg/planner/domain/value-objects"

var ideasList = []value_objects.BaseNode{
	NewIdea("Create repository"),
	NewIdea("Brainstorm use cases"),
	NewIdea("Create data models"),
	NewIdea("Define relationships"),
	NewIdea("Define scope of the problem"),
	NewIdea("Create sequence diagrams"),
	NewIdea("Create C4 Diagrams"),
	NewIdea("Create CI Pipeline"),
	NewIdea("Create CD Pipeline"),
	NewIdea("Setup actions to display tests"),
	NewIdea("Add PR template"),
	NewIdea("Create K8s infrastructure"),
	NewIdea("Define context"),
	NewIdea("Benchmark performance"),
	NewIdea("Create integration tests"),
	NewIdea("Create E2E tests"),
	NewIdea("Create Postman Documentation"),
	NewIdea("Create Postman Tests (E2E)"),
	NewIdea("Create Swagger Docs"),
	NewIdea("Add OAuth2"),
	NewIdea("Add Google Calendar integration"),
	NewIdea("Create software documentation"),
}

type args struct {
	name  string
	i     int
	idea  *Idea
	ideas []value_objects.BaseNode
	s, t  int
}

type fields struct {
	node  *value_objects.Node
	ideas []value_objects.BaseNode
}

type addIdeaResults struct {
	success bool
	state   []value_objects.BaseNode
}

type insertListResults struct {
	success bool
	idea    *Idea
}

var newIdeaTests = []struct {
	name string
	args args
	want *Idea
}{
	{
		name: "Creates an Idea named 'Create graph model'",
		args: args{name: "Create graph model"},
		want: func() *Idea {
			n := Idea{value_objects.NewNode("Create graph model")}
			return &n
		}(),
	},
}

var newListTests = []struct {
	name string
	args args
	want *List
}{
	{
		name: "Creates a List named 'Create an application to manage organization'",
		args: args{name: "Create an application to manage organization"},
		want: &List{
			Node: value_objects.NewNode(
				"Create an application to manage organization"),
			ideas: make([]value_objects.BaseNode, 0),
		},
	},
}

var listAddIdeaTests = []struct {
	name   string
	fields fields
	args   args
	want   addIdeaResults
}{
	{
		name: "Pushes 'Make benchmarks after logic'",
		fields: fields{
			node:  value_objects.NewNode("Makes benchmark after logic"),
			ideas: make([]value_objects.BaseNode, 0),
		},
		args: args{idea: NewIdea("Make benchmarks after logic")},
		want: addIdeaResults{
			success: true,
			state:   []value_objects.BaseNode{NewIdea("Make benchmarks after logic")},
		},
	},
}

var listInsertAtTests = []struct {
	name   string
	fields fields
	args   args
	want   insertListResults
}{
	{
		name: "Inserts at index 0 successfully",
		fields: fields{
			node: value_objects.NewNode("Create app"),
			ideas: []value_objects.BaseNode{
				NewIdea("Create Diagrams"),
				NewIdea("Create Business Logic"),
				NewIdea("Create Folder Structure"),
			},
		},
		args: args{
			i:    0,
			idea: NewIdea("Create GitHub Repo"),
		},
		want: insertListResults{
			success: true,
			idea:    NewIdea("Create GitHub Repo"),
		},
	},
	{
		name: "Inserts at index 2 successfully",
		fields: fields{
			node: value_objects.NewNode("Create app"),
			ideas: []value_objects.BaseNode{
				NewIdea("Create Diagrams"),
				NewIdea("Create Business Logic"),
				NewIdea("Create Folder Structure"),
			},
		},
		args: args{
			i:    2,
			idea: NewIdea("Create GitHub Repo"),
		},
		want: insertListResults{
			success: true,
			idea:    NewIdea("Create GitHub Repo"),
		},
	},
	{
		name: "Fails to push into ",
		fields: fields{
			node: value_objects.NewNode("Create app"),
			ideas: []value_objects.BaseNode{
				NewIdea("Create Diagrams"),
				NewIdea("Create Business Logic"),
				NewIdea("Create Folder Structure"),
			},
		},
		args: args{
			i:    15,
			idea: NewIdea("Create GitHub Repo"),
		},
		want: insertListResults{
			success: false,
			idea:    nil,
		},
	},
	{
		name: "Fails to push into ",
		fields: fields{
			node: value_objects.NewNode("Create app"),
			ideas: []value_objects.BaseNode{
				NewIdea("Create Diagrams"),
				NewIdea("Create Business Logic"),
				NewIdea("Create Folder Structure"),
			},
		},
		args: args{
			i:    -1,
			idea: NewIdea("Create GitHub Repo"),
		},
		want: insertListResults{
			success: false,
			idea:    nil,
		},
	},
	{
		name: "Fails to push into ",
		fields: fields{
			node: value_objects.NewNode("Create app"),
			ideas: []value_objects.BaseNode{
				NewIdea("Create Diagrams"),
				NewIdea("Create Business Logic"),
				NewIdea("Create Folder Structure"),
			},
		},
		args: args{
			i:    4,
			idea: NewIdea("Create GitHub Repo"),
		},
		want: insertListResults{
			success: false,
			idea:    nil,
		},
	},
}

var listDeleteAtTests = []struct {
	name   string
	fields fields
	args   args
	want   bool
}{
	{
		name: "Deletes at index 2 successfully",
		fields: fields{
			node: value_objects.NewNode("Create app"),
			ideas: []value_objects.BaseNode{
				NewIdea("Create Diagrams"),
				NewIdea("Create Business Logic"),
				NewIdea("Create Folder Structure"),
			},
		},
		args: args{
			i: 2,
		},
		want: true,
	},
	{
		name: "Negative index returns false",
		fields: fields{
			node:  value_objects.NewNode("Create app"),
			ideas: []value_objects.BaseNode{},
		},
		args: args{
			i: -1,
		},
		want: false,
	},
	{
		name: "Positive index on empty List returns false",
		fields: fields{
			node:  value_objects.NewNode("Create app"),
			ideas: []value_objects.BaseNode{},
		},
		args: args{
			i: 2,
		},
		want: false,
	},
	{
		name: "Deletes at index 0 successfully",
		fields: fields{
			node: value_objects.NewNode("Create app"),
			ideas: []value_objects.BaseNode{
				NewIdea("Create Diagrams"),
				NewIdea("Create Business Logic"),
				NewIdea("Create Folder Structure"),
			},
		},
		args: args{
			i: 0,
		},
		want: true,
	},
	{
		name: "Deletes at index 1 successfully",
		fields: fields{
			node: value_objects.NewNode("Create app"),
			ideas: []value_objects.BaseNode{
				NewIdea("Create Diagrams"),
				NewIdea("Create Business Logic"),
				NewIdea("Create Folder Structure"),
			},
		},
		args: args{
			i: 1,
		},
		want: true,
	},
}

var ideaStringMethodTests = []struct {
	name string
	i    Idea
	want string
}{
	{
		name: "Idea name is 'Alejandro'",
		i:    *NewIdea("Alejandro"),
		want: "Alejandro",
	},
	{
		name: "Idea name is 'Patricia'",
		i:    *NewIdea("Patrica"),
		want: "Patrica",
	},
	{
		name: "Idea name is 'Jutta'",
		i:    *NewIdea("Jutta"),
		want: "Jutta",
	},
}

var listMergeListsTests = []struct {
	name   string
	fields fields
	args   args
	want   bool
}{
	{
		name:   "",
		fields: fields{},
		args:   args{},
		want:   false,
	},
}

var listSwapTests = []struct {
	name   string
	fields fields
	args   args
	want   bool
}{
	{
		name: "Moves idea at index 2 to index 5",
		fields: fields{
			node:  value_objects.NewNode("To do list"),
			ideas: ideasList,
		},
		args: args{
			s: 2,
			t: 5,
		},
		want: true,
	},
	{
		name: "Fails swapping idea at index 2 to index 20",
		fields: fields{
			node:  value_objects.NewNode("To do list"),
			ideas: ideasList,
		},
		args: args{
			s: 2,
			t: 22,
		},
		want: false,
	},
}
