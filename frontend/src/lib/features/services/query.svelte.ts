import { api } from '@/axios';
import { createQuery } from '@tanstack/svelte-query';
import type { AppServiceDetails, GetBranchDomainRes, GetEnvRes, ServiceListResponse } from './type';
import { GetUserData } from '../global/query';

export const getOrgServicesQueryKey = (orgId: string) => ['services-list', 'org', orgId] as const;

export function useGetAllServicesQuery() {
	const { org_id } = GetUserData();
	return createQuery(() => ({
		queryKey: getOrgServicesQueryKey(org_id),
		queryFn: async () => api.get<ServiceListResponse[]>('/service').then((res) => res.data)
	}));
}

export function useGetServiceDetailsQuery(getID: () => string) {
	const serviceId = getID();
	return createQuery(() => ({
		queryKey: ['service-details', serviceId],
		queryFn: async () =>
			api.get<AppServiceDetails>(`/service/app/${serviceId}`).then((res) => res.data),
		enabled: serviceId !== ''
	}));
}

export function useGetBranchDomainQuery(getServiceId: () => string) {
	const serviceId = getServiceId();
	return createQuery(() => ({
		queryKey: ['branch-domain', serviceId],
		queryFn: async () =>
			api
				.get<GetBranchDomainRes>('/service/app/domain', {
					params: { service_id: serviceId }
				})
				.then((res) => res.data),
		enabled: serviceId !== ''
	}));
}

export function useGetServiceEnvQuery(getServiceId: () => string) {
	const serviceId = getServiceId();
	return createQuery(() => ({
		queryKey: ['service-env', serviceId],
		queryFn: async () =>
			api
				.get<GetEnvRes>('/service/app/env', {
					params: { service_id: serviceId }
				})
				.then((res) => res.data),
		enabled: serviceId !== ''
	}));
}
