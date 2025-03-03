package cmd

import (
	"github.com/de1ux/gitstuff/audit"
	"github.com/de1ux/gitstuff/shell"
	"github.com/spf13/cobra"
)

var script = `
branch=$(git branch --no-color 2> /dev/null | sed -e '/^[^*]/d' -e 's/* \(.*\)/\1/')

# Retrieve destination branch
destination_branch=${branch%--*}

# Determine the default origin branch
if [[ $branch == $destination_branch ]]; then
  master_exists=$(git branch | grep main)

  if [[ -z $master_exists ]]; then
	destination_branch=develop
  else
	destination_branch=main
  fi
fi

# Retrieve repository URL
repository_url=$(git remote get-url origin | sed -e 's/git@//' -e 's/\.git//' -e 's/:/\//')

if [[ $repository_url == github* ]]; then
  pr_url=https://$repository_url/compare/$destination_branch...$branch
elif [[ $repository_url == gitlab* ]]; then
  pr_url="https://$repository_url/merge_requests/new?merge_request%5Bsource_branch%5D=$branch&merge_request%5Btarget_branch%5D=$destination_branch"
fi

# Open the new Pull Request URL
open $pr_url
`

var OpenCmd = &cobra.Command{
	Use:  "open",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		shell.BashStatusCode(script)
		return audit.Write(branch, "opening pull request")
	},
}
