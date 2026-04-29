// AI summary: Feature-scoped extraction of route query/mutation contracts to keep page runes and form logic focused on UI orchestration.
export interface GithubApp {
	name: string;
	app_id: number;
	created_at: string;
}

export interface GitProvider {
	name: string;
	icon: string;
	redirect: string;
}

export interface DeleteGithubAppPayload {
	app_id: number;
}
