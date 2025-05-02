# Supported Sources

- GitHub
- GitLab
- Homebrew (binaries only, if anything has a dependency, it will not work at this time)
- Hashicorp (special handling for their releases, pointing to GitHub repos will automatically pass through)
- Kubernetes (special handling for their releases, pointing to GitHub repos will automatically pass through)

## Authentication

Distillery supports authentication for GitHub and GitLab. There are CLI options to pass in a token, but the preferred
method is to set the `DISTILLERY_GITHUB_TOKEN` or `DISTILLERY_GITLAB_TOKEN` environment variables using a tool like
[direnv](https://direnv.net/).

This allows you to bypass any API rate limits that might be in place for unauthenticated requests, but more importantly
it allows you to install private repositories that you have access to!
