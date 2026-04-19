// GitHub integration contracts are shared by the provider page and service picker flow.
export type GitProviderKey = 'github' | 'gitlab' | 'bitbucket';

export interface GithubApp {
	name: string;
	app_id: number;
	created_at: string;
}

export interface GithubRepo {
	id: number;
	name: string;
	full_name: string;
	html_url: string;
	private: boolean;
	default_branch: string;
}

export interface GitProviderOption {
	key: GitProviderKey;
	name: string;
	icon: string;
	api: string;
}

export interface GetRepoResult {
	status: number;
	repos: GithubRepo[];
	message: string;
	provider: GitProviderKey;
}
