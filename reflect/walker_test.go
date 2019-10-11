package reflect

import (
	"reflect"
	"testing"
)



type Profile struct {
	Age int
	City string
}

type Person struct {
	Name string
	Profile Profile
}

func TestWalker(t *testing.T)  {

	cases := []struct {
		Name string
		Input interface{}
		ExpectedCalls []string
	}{
		{
			"Struct with one string field",
			struct {
				Name string
			}{"chris"},
			[]string{"chris"},
		},
		{
			"Struct with two string fields",
			struct {
				Name string
				City string
			}{"chris", "london"},
			[]string{"chris", "london"},
		},
		{
			"Struct with non string field",
			struct {
				Name string
				Age  int
			}{"chris", 33},
			[]string{"chris"},
		},
		{
			"Nested fields",
			Person{
				"chris",
				Profile{
					Age:  33,
					City: "london",},
			},
			[]string{"chris", "london"},
		},

		{
			"pointer to things",
			&Person{
				Name:    "chris",
				Profile: Profile{
					33,
					"london",
				},
			},[]string{"chris", "london"},
		},
		{
			"slices",
			[]Profile {
				{33, "london"},
				{34, "shanghai"},
			},
			[]string{"london", "shanghai"},
		},
		{
			"arrays",
			[2]Profile{
				{33, "london"},
				{34,"shanghai"},
			},
			[]string{"london", "shanghai"},
		},
	}
	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if ! reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}
	//go 里的map 书无序的。 所以不能保证返回数组的顺序
	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Foo": "bar",
			"Baz": "boz",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "bar")
		assertContains(t, got, "boz")

	})
}

func assertContains(t *testing.T, strings []string, expected string) {
	contains := false
	for _, x := range strings {
		if x == expected {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected  %+v to contains '%s' but it didnt", strings, expected)
	}
	
}

