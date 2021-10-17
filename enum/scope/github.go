package scope

type GithubScope struct {
	Scope       string
	Description string
	Default     bool
}

func (g GithubScope) GetScope() string {
	return g.Scope
}

func (g GithubScope) IsDefault() bool {
	return g.Default
}

var GithubScopes = []AuthScope{
	REPO_STATUS,
	REPO_DEPLOYMENT,
	PUBLIC_REPO,
	REPO_INVITE,
	SECURITY_EVENTS,
	WRITE_REPO_HOOK,
	READ_REPO_HOOK,
	ADMIN_ORG,
	WRITE_ORG,
	READ_ORG,
	ADMIN_PUBLIC_KEY,
	WRITE_PUBLIC_KEY,
	READ_PUBLIC_KEY,
	GIST,
	NOTIFICATIONS,
	USER,
	READ_USER,
	USER_EMAIL,
	USER_FOLLOW,
	DELETE_REPO,
	WRITE_DISCUSSION,
	READ_DISCUSSION,
	WRITE_PACKAGES,
	READ_PACKAGES,
	DELETE_PACKAGES,
	ADMIN_GPG_KEY,
	WRITE_GPG_KEY,
	READ_GPG_KEY,
	WORKFLOW,
}

/**
 * {@code scope} 含义，以{@code description} 为准
 */
var (
	REPO_STATUS      = GithubScope{"repo:status", "Grants read/write access to public and private repository commit statuses. This scope is only necessary to grant other users or services access to private repository commit statuses <em>without</em> granting access to the code.", false}
	REPO_DEPLOYMENT  = GithubScope{"repo_deployment", "Grants access to deployment statuses for public and private repositories. This scope is only necessary to grant other users or services access to deployment statuses, <em>without</em> granting access to the code.", false}
	PUBLIC_REPO      = GithubScope{"public_repo", "Limits access to public repositories. That includes read/write access to code, commit statuses, repository projects, collaborators, and deployment statuses for public repositories and organizations. Also required for starring public repositories.", false}
	REPO_INVITE      = GithubScope{"repo:invite", "Grants accept/decline abilities for invitations to collaborate on a repository. This scope is only necessary to grant other users or services access to invites <em>without</em> granting access to the code.", false}
	SECURITY_EVENTS  = GithubScope{"security_events", "Grants read and write access to security events in the code scanning API.", false}
	WRITE_REPO_HOOK  = GithubScope{"write:repo_hook", "Grants read, write, and ping access to hooks in public or private repositories.", false}
	READ_REPO_HOOK   = GithubScope{"read:repo_hook", "Grants read and ping access to hooks in public or private repositories.", false}
	ADMIN_ORG        = GithubScope{"admin:org", "Fully manage the organization and its teams, projects, and memberships.", false}
	WRITE_ORG        = GithubScope{"write:org", "Read and write access to organization membership, organization projects, and team membership.", false}
	READ_ORG         = GithubScope{"read:org", "Read-only access to organization membership, organization projects, and team membership.", false}
	ADMIN_PUBLIC_KEY = GithubScope{"admin:public_key", "Fully manage public keys.", false}
	WRITE_PUBLIC_KEY = GithubScope{"write:public_key", "Create, list, and view details for public keys.", false}
	READ_PUBLIC_KEY  = GithubScope{"read:public_key", "List and view details for public keys.", false}
	GIST             = GithubScope{"gist", "Grants write access to gists.", false}
	NOTIFICATIONS    = GithubScope{"notifications", "Grants: <br>* read access to a user's notifications <br>* mark as read access to threads <br>* watch and unwatch access to a repository, and <br>* read, write, and delete access to thread subscriptions.", false}
	USER             = GithubScope{"user", "Grants read/write access to profile info only.  Note that this scope includes <code>user:email</code> and <code>user:follow</code>.", false}
	READ_USER        = GithubScope{"read:user", "Grants access to read a user's profile data.", false}
	USER_EMAIL       = GithubScope{"user:email", "Grants read access to a user's email addresses.", false}
	USER_FOLLOW      = GithubScope{"user:follow", "Grants access to follow or unfollow other users.", false}
	DELETE_REPO      = GithubScope{"delete_repo", "Grants access to delete adminable repositories.", false}
	WRITE_DISCUSSION = GithubScope{"write:discussion", "Allows read and write access for team discussions.", false}
	READ_DISCUSSION  = GithubScope{"read:discussion", "Allows read access for team discussions.", false}
	WRITE_PACKAGES   = GithubScope{"write:packages", "Grants access to upload or publish a package in GitHub Packages. For more information, see \"<a href=\"https://help.github.com/github/managing-packages-with-github-packages/publishing-a-package\">Publishing a package</a>\" in the GitHub Help documentation.", false}
	READ_PACKAGES    = GithubScope{"read:packages", "Grants access to download or install packages from GitHub Packages. For more information, see \"<a href=\"https://help.github.com/github/managing-packages-with-github-packages/installing-a-package\">Installing a package</a>\" in the GitHub Help documentation.", false}
	DELETE_PACKAGES  = GithubScope{"delete:packages", "Grants access to delete packages from GitHub Packages. For more information, see \"<a href=\"https://help.github.com/github/managing-packages-with-github-packages/deleting-a-package\">Deleting packages</a>\" in the GitHub Help documentation.", false}
	ADMIN_GPG_KEY    = GithubScope{"admin:gpg_key", "Fully manage GPG keys.", false}
	WRITE_GPG_KEY    = GithubScope{"write:gpg_key", "Create, list, and view details for GPG keys.", false}
	READ_GPG_KEY     = GithubScope{"read:gpg_key", "List and view details for GPG keys.", false}
	WORKFLOW         = GithubScope{"workflow", "Grants the ability to add and update GitHub Actions workflow files. Workflow files can be committed without this scope if the same file (with both the same path and contents) exists on another branch in the same repository.", false}
)
