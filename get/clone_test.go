package get

import "testing"

func TestCloneShouldGetsRepos(t *testing.T) {

	testCases := []struct {
		desc string
		repo string
	}{
		{
			desc: "zahak",
			repo: "https://github.com/amanjpro/zahak.git",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {

			cloner, err := NewCloner(tC.repo)
			if err != nil {
				t.Fatal(err)
			}

			err = cloner.Clone()
			if err != nil {
				t.Fatal(err)
			}
		})
	}
}
