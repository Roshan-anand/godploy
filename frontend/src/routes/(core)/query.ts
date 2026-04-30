import { api } from '@/axios';
import type { AuthResponse } from '@/features/auth/type';
import { queryClient } from '@/query';

const authUserQueryKey = () => ['auth', 'user'];

export const GetAuthUserData = () => queryClient.getQueryData<AuthResponse>(authUserQueryKey());

export const fetchAuthUserQuery = () =>
	queryClient.fetchQuery({
		queryKey: authUserQueryKey(),
		queryFn: () => api.get<AuthResponse>('/auth/user').then((res) => res.data)
	});
