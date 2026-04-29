// AI summary: Feature-scoped extraction of route query/mutation contracts to keep page runes and form logic focused on UI orchestration.
export interface Project {
	id: string;
	name: string;
	description: string;
}

export interface CreateProjectPayload {
	project_name: string;
	description: string;
}

export interface DeleteProjectPayload {
	id: string;
}

export interface ApiMessageRes {
	message: string;
}
