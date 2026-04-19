import type { GitProviderKey } from '../github/types';

// Service contracts are shared between the detail page and create-service dialog.
export type ServiceType = 'psql' | 'app';

type ServiceBase = {
	id: string;
	project_id: string;
	type: ServiceType;
	service_id: string;
	name: string;
	app_name: string;
	description: string;
	created_at: string;
};

type PsqlService = {
	type: 'psql';
	db_name: string;
	db_user: string;
	db_password: string;
	image: string;
	internal_url: string;
};

type AppService = {
	type: 'app';
};

export type ServiceDetails = ServiceBase & (PsqlService | AppService);

export interface CreateServiceResponse {
	id: string;
	type: ServiceType;
}

export interface CreateAppServiceBody {
	project_id: string;
	name: string;
	description: string;
	app_name: string;
	git_provider: GitProviderKey;
	git_repo_id: string;
	git_repo_name: string;
	git_branch: string;
	build_path: string;
}

export interface CreatePsqlServiceBody {
	project_id: string;
	name: string;
	description: string;
	app_name: string;
	db_name: string;
	db_user: string;
	db_password: string;
	image: string;
}

export type CreateServicePayload =
	| { type: 'app'; body: CreateAppServiceBody }
	| { type: 'psql'; body: CreatePsqlServiceBody };
