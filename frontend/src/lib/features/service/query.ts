import { api } from '@/axios';
import { createQuery } from '@tanstack/svelte-query';
import type { ServiceDetails, ServiceType } from './types';

type ServiceParams = {
	getServiceType: () => ServiceType;
	getServiceId: () => string;
};

// Service details are fetched from a single shared query helper used by the service page.
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
