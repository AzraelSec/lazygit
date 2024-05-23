package branch

import (
	"github.com/jesseduffield/lazygit/pkg/config"
	. "github.com/jesseduffield/lazygit/pkg/integration/components"
	"github.com/jesseduffield/lazygit/pkg/integration/tests/shared"
)

var Squash = NewIntegrationTest(NewIntegrationTestArgs{
	Description:  "Squash a branch in the working tree",
	ExtraCmdArgs: []string{},
	Skip:         false,
	SetupConfig:  func(config *config.AppConfig) {},
	SetupRepo: func(shell *Shell) {
		shared.CreateMergeCommit(shell)
	},
	Run: func(t *TestDriver, keys config.KeybindingConfig) {
		t.Views().Commits().TopLines(
			Contains("first change"),
			Contains("original"),
		)

		t.Views().Branches().
			Focus().
			Lines(
				Contains("change-branch"),
				Contains("original-branch"),
			).
			SelectNextItem().
			Press(keys.Branches.SquashBranch)

		t.ExpectPopup().Menu().
			Title(Equals("Squash")).
			Select(Contains("Squash files in a new commit")).
			Confirm()

		t.Views().Commits().TopLines(
			Contains("Squash branch change-branch into original-branch"),
			Contains("first change"),
			Contains("original"),
		)
	},
})
