package cmd

import (
	"github.com/de1ux/gitstuff/audit"
	"github.com/de1ux/gitstuff/shell"
	"github.com/spf13/cobra"
)

var script = `
branch=$(git branch --no-color 2> /dev/null | sed -e '/^[^*]/d' -e 's/* \(.*\)/\1/')

# Try to get existing PR URL for this specific branch using gh CLI
if command -v gh &> /dev/null; then
  pr_url=$(gh pr list --head "$branch" --json url --jq '.[0].url' 2>/dev/null)

  # If PR exists for this branch, open it and exit
  if [[ -n $pr_url ]]; then
    open $pr_url
    exit 0
  fi
fi

# No existing PR found, open file explorer at current directory
# Retrieve repository URL
repository_url=$(git remote get-url origin | sed -e 's/git@//' -e 's/\.git//' -e 's/:/\//')

# Get the path relative to the git root
git_root=$(git rev-parse --show-toplevel)
current_dir=$(pwd)
relative_path=${current_dir#$git_root}

# Build the file explorer URL
if [[ $repository_url == github* ]]; then
  if [[ -z $relative_path || $relative_path == "/" ]]; then
    file_url="https://$repository_url/tree/$branch"
  else
    file_url="https://$repository_url/tree/$branch$relative_path"
  fi
elif [[ $repository_url == gitlab* ]]; then
  if [[ -z $relative_path || $relative_path == "/" ]]; then
    file_url="https://$repository_url/-/tree/$branch"
  else
    file_url="https://$repository_url/-/tree/$branch$relative_path"
  fi
fi

# Open the file explorer URL
open $file_url
`

var OpenCmd = &cobra.Command{
	Use:  "open",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		shell.BashStatusCode(script)
		return audit.Write(branch, "opening in browser")
	},
}
