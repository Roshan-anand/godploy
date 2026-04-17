import { api } from '@/axios';
import { createQuery } from '@tanstack/svelte-query';

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

type ServiceParams = {
	getServiceType: () => ServiceType;
	getServiceId: () => string;
};

export const createServiceDetailQuery = ({ getServiceType, getServiceId }: ServiceParams) => {
	return createQuery(() => ({
		queryKey: ['service-details', getServiceType(), getServiceId()],
		queryFn: async () => {
			const serviceType = getServiceType();
			const serviceId = getServiceId();
			const url =
				serviceType === 'app' ? `/service/app/${serviceId}` : `/service/psql/${serviceId}`;
			return api.get<ServiceDetails>(url).then((res) => res.data);
		},
		enabled: (() => {
			const serviceType = getServiceType();
			const serviceId = getServiceId();
			return serviceId !== '' && (serviceType === 'psql' || serviceType === 'app');
		})()
	}));
};
