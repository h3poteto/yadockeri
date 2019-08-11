export type Commit = {
  sha: string
  url: string
}
export type GitHubBranch = {
  name: string
  commit: Commit
  protected: boolean
}
