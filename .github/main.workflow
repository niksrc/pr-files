workflow "Github Release" {
  on       = "push"
  resolves = ["goreleaser"]
}

action "is-tag" {
  uses = "actions/bin/filter@master"
  args = "tag"
}

action "goreleaser" {
  uses = "docker://goreleaser/goreleaser"

  # at least GITHUB_TOKEN is required, you may need more though

  args    = "release"
  needs   = ["is-tag"]
  secrets = ["GITHUB_TOKEN"]

  # at least GITHUB_TOKEN is required, you may need more though
}
