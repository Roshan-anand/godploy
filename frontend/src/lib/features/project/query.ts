import { api } from '@/axios';
import { userState } from '@/store/userState.svelte';
import { createQuery } from '@tanstack/svelte-query';
import type { Project } from './types';

// The project list query is reused by the project page and the service creation dialog.
export function getProjectsQueryKey(cacheScope: 'default' | 'service-create' = 'default') {
	return cacheScope === 'service-create'
		? ['projects', userState.currentOrg.id, 'service-create']
		: ['projects', userState.currentOrg.id];
}

export function createProjectsQuery(
	isProjectScoped: () => boolean = () => false,
	cacheScope: 'default' | 'service-create' = 'default'
) {
	return createQuery(() => ({
		queryKey: getProjectsQueryKey(cacheScope),
		queryFn: () => api.get<Project[]>('/project/all').then((res) => res.data),
		enabled: !isProjectScoped() && userState.currentOrg.id !== ''
	}));
}
