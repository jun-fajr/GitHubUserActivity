# GitHub User Activity CLI

GitHub User Activity CLI is a simple command-line application built with Go that fetches and displays recent activity for a specified GitHub user. It allows you to view recent events, such as commits, issues created, starred repositories, and more.
Task from https://roadmap.sh/projects/github-user-activity

## Features

- Fetch and display recent activity for any GitHub user
- View events such as pushes, issue creations, stars, and forks
- Simple, command-line interface for quick and easy usage

## Requirements

- Go (Golang) 1.16 or higher
- Internet connection (to access GitHub's API)

## Installation

1. Clone this repository:

   ```bash
   git clone https://github.com/jun-fajr/GitHubUserActivity.git
   cd GitHubUserActivity
   ```

2. Build the project:

   ```bash
   go build -o github-activity main.go
   ```

3. (Optional) Move `github-activity` to a directory in your `PATH` (e.g., `/usr/local/bin`) for easy access:

   ```bash
   sudo mv github-activity /usr/local/bin/
   ```

## Usage

To use the CLI, run the `github-activity` command followed by the GitHub username you want to fetch activity for.

### Example

Fetch recent activity for the GitHub user `jun-fajr`:

```bash
github-activity jun-fajr
```

### Sample Output

```plaintext
Fetching activity for user: jun-fajr
- Forked koprab/monalisa-font
- Pushed to jun-fajr/junizarfajri-betest
- Created a repository jun-fajr/junizarfajri-betest
- Created a repository jun-fajr/junizarfajri-betest
- Created a repository jun-fajr/test-be
- Created a repository jun-fajr/test-be
- Pushed to jun-fajr/Web-Scraping
...
```

### Command Reference

- **`github-activity <username>`**: Fetches and displays recent activity for the specified GitHub user.

## Error Handling

The CLI will handle common errors, such as:
- Invalid GitHub usernames
- Network issues
- Exceeding GitHub's API rate limits (for unauthenticated requests)

## GitHub API Rate Limiting

GitHub's API imposes rate limits on unauthenticated requests. If you make too many requests in a short period, you may encounter rate limits. To avoid this, consider using the tool sparingly or authenticating requests with a GitHub API token if needed (not implemented in this version).

## License

This project is licensed under the MIT License.

## Contribution

Feel free to open issues or submit pull requests if you'd like to contribute to this project. 

---

Enjoy tracking GitHub activity!