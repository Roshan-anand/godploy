import { api } from '@/axios';
import { createQuery } from '@tanstack/svelte-query';
import type { GithubApp } from './type';
import { GetUserData } from '../global/query';
import type { ApiRes } from '@/types';

export const getGithubAppsQueryKey = (orgId: string) => ['github-apps', orgId] as const;
export function useGithubAppsQuery() {
	const { org_id } = GetUserData();
	return createQuery(() => ({
		queryKey: getGithubAppsQueryKey(org_id),
		queryFn: () =>
			api.get<ApiRes<GithubApp[] | null>>('/provider/github/app/list').then((res) => res.data.data),
		enabled: org_id != ''
	}));
}
