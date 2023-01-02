package get

import (
	"os"
	"os/exec"
)

// Cloner represents "git clone repo" call.
// git clone git@github.com:amanjpro/zahak.git
// gh repo clone amanjpro/zahak
// git clone https://github.com/amanjpro/zahak.git
type Cloner struct {
	repo string
	cmd  *exec.Cmd
}

func NewCloner(repo string) (*Cloner, error) {

	cloner := new(Cloner)
	//ctx := context.Background()
	cmd := exec.Command("git", "clone", repo)
	cmd.Stdout = os.Stdout

	cloner.cmd = cmd
	cloner.repo = repo

	return cloner, nil
}

func (c *Cloner) Clone() error {
	return c.cmd.Run()

}
