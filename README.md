# gitstuff

## Install

### Quick Install (Latest Release)
```bash
# For Apple Silicon (M1/M2/M3)
curl -L https://github.com/de1ux/gitstuff/releases/latest/download/gitstuff-darwin-arm64 -o gitstuff && chmod +x gitstuff && sudo mv gitstuff /usr/local/bin/

# For Intel Mac
curl -L https://github.com/de1ux/gitstuff/releases/latest/download/gitstuff-darwin-amd64 -o gitstuff && chmod +x gitstuff && sudo mv gitstuff /usr/local/bin/

# For Linux ARM64
curl -L https://github.com/de1ux/gitstuff/releases/latest/download/gitstuff-linux-arm64 -o gitstuff && chmod +x gitstuff && sudo mv gitstuff /usr/local/bin/

# For Linux AMD64
curl -L https://github.com/de1ux/gitstuff/releases/latest/download/gitstuff-linux-amd64 -o gitstuff && chmod +x gitstuff && sudo mv gitstuff /usr/local/bin/
```

**Note: requires Bazel 8 or later to build**
```bash
bazel build //:gitstuff

# move the binary to /usr/local/bin, or wherever on $PATH 
cp $(bazel cquery //:gitstuff --output files) /usr/local/bin/gitstuff

# if that doesn't work, copy directly out of the bazel-out directory
# cp bazel-out/darwin_arm64-fastbuild/bin/gitstuff_/gitstuff <wherever>

# Update depencies, if go.mod changes at all
bazel mod tidy
bazel run //:gazelle -- update-repos -from_file=go.mod -to_macro=deps.bzl%go_dependencies
```

I also add these aliases to make it easier to reach the binary:
```bash
alias ggpull="gitstuff pull"
alias ggpush="gitstuff push"
alias gback="gitstuff back"
alias gforward="gitstuff forward"
alias gco="gitstuff checkout"
alias gsubmit="gitstuff submit"
alias ginit="gitstuff init"
alias gcommit="gitstuff commit"
alias gopen="gitstuff open"
alias gfetch="gitstuff fetch"
```

## Prerequisites

### GitHub CLI

The `gitstuff submit` command requires GitHub CLI (`gh`) to be installed and authenticated. To set this up:

1. Install GitHub CLI:
   - macOS: `brew install gh`
   - Windows: `winget install GitHub.cli`
   - Linux: See [GitHub CLI installation guide](https://github.com/cli/cli#installation)

2. Authenticate with GitHub:
   ```bash
   gh auth login
   ```
   Follow the prompts to complete authentication. If your organization requires SSO, the GitHub CLI will automatically handle the SSO flow.

## Commands

### `gcommit <message>`

Commits all changes with the given message, and runs  

Same as `git commit -am <message>` and `git push origin <current branch name>`.

### `ggpull`

Pulls the latest changes from the remote repository. 

Same as `git pull origin <current branch name>`

### `ggpush`

Pushes the latest changes to the remote repository.

Same as `git push origin <current branch name>`

### `gfetch [branch name]`

Fetches changes from the remote repository. If a branch name is provided, fetches only that branch. If no branch name is provided, fetches all branches from all remotes.

Same as `git fetch origin <branch name>:<branch name>` or `git fetch --all --prune`

## `gco <branch name>`

Checkout a branch. Passing `-b` will create a new branch. Passing `-u` will fetch the branch from origin before checking it out.

Same as `git checkout <branch name>`, `git checkout -b <branch name>`, or `git fetch origin <branch name>:<branch name> && git checkout <branch name>` (with `-u`)

### `gback` / `gforward`

`gback` checks out the last branch you were on. 

`gforward` checks out the next branch you were on.

### `gsubmit`

1. Pushes the current branch to the remote repository (`ggpush`)
2. Creates a pull request on GitHub
3. Opens the pull request in your browser

### `gopen`

Opens the current branch on Github, in the Pull Request view (for comparing changes).

### `ginit <branch name>`

Checks out `main`, runs `ggpull`, and creates a new branch with the given name (`gco -b`).

If there are working changes on the branch, asks if you want to `git reset --hard HEAD` before creating the new branch.