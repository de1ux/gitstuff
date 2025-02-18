# gitstuff

## Install

**Note: requires Bazel 8 or later to build**
```bash
bazel build //:gitstuff

# move the binary to /usr/local/bin, or wherever on $PATH 
cp $(bazel cquery //:gitstuff --output files) /usr/local/bin/gitstuff

# if that doesn't work, copy directly out of the bazel-out directory
# cp bazel-out/darwin_arm64-fastbuild/bin/gitstuff_/gitstuff <wherever>
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
```

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

## `gco <branch name>`

Checkout a branch. Passing `-b` will create a new branch.

Same as `git checkout <branch name>` or `git checkout -b <branch name>`

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
