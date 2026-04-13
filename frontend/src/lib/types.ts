export type ServiceType = 'psql' | 'app';

export type ServiceBase = {
	id: string;
	project_id: string;
	type: ServiceType;
	service_id: string;
	name: string;
	app_name: string;
	description: string;
	created_at: string;
};

export type PsqlService = {
	type: 'psql';
	db_name: string;
	db_user: string;
	db_password: string;
	image: string;
	internal_url: string;
};

export type AppService = {
	type: 'app';
};

export type ServiceDetails = ServiceBase & (PsqlService | AppService);
