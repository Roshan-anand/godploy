import { api } from '@/axios';
import { createQuery } from '@tanstack/svelte-query';
import type { AppServiceDetails, GetBranchDomainRes, GetEnvRes, ServiceListResponse } from './type';
import { GetUserData } from '../global/query';
import type { ApiRes } from '@/types';

export const getOrgServicesQueryKey = (orgId: string) => ['services-list', 'org', orgId] as const;

export function useGetAllServicesQuery() {
	const { org_id } = GetUserData();
	return createQuery(() => ({
		queryKey: getOrgServicesQueryKey(org_id),
		queryFn: async () => api.get<ApiRes<ServiceListResponse[]>>('/service').then((res) => res.data.data)
	}));
}

export function useGetServiceDetailsQuery(getID: () => string) {
	const serviceId = getID();
	return createQuery(() => ({
		queryKey: ['service-details', serviceId],
		queryFn: async () =>
			api.get<ApiRes<AppServiceDetails>>(`/service/app/${serviceId}`).then((res) => res.data.data),
		enabled: serviceId !== ''
	}));
}

export function useGetBranchDomainQuery(getServiceId: () => string) {
	const serviceId = getServiceId();
	return createQuery(() => ({
		queryKey: ['branch-domain', serviceId],
		queryFn: async () =>
			api
				.get<ApiRes<GetBranchDomainRes>>('/service/app/domain', {
					params: { service_id: serviceId }
				})
				.then((res) => res.data.data),
		enabled: serviceId !== ''
	}));
}

export function useGetServiceEnvQuery(getServiceId: () => string) {
	const serviceId = getServiceId();
	return createQuery(() => ({
		queryKey: ['service-env', serviceId],
		queryFn: async () =>
			api
				.get<ApiRes<GetEnvRes>>('/service/app/env', {
					params: { service_id: serviceId }
				})
				.then((res) => res.data.data),
		enabled: serviceId !== ''
	}));
}
