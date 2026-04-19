// Project contracts are shared between the project list page and service creation flow.
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

export interface ApiRes {
	message: string;
}
