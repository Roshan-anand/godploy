import { api } from '@/axios';
import type { AuthResponse } from '@/features/auth';
import { queryClient } from '@/query';
import type { ApiRes } from '@/types';

const authUserQueryKey = () => ['auth', 'user'];

// query to fetch auth user data
export const fetchUserQuery = () =>
	queryClient.fetchQuery({
		queryKey: authUserQueryKey(),
		queryFn: () => api.get<ApiRes<AuthResponse>>('/auth/user').then((res) => res.data.data)
	});

// helper function to get auth user data from cache
export const GetUserData = (): AuthResponse =>
	queryClient.getQueryData<AuthResponse>(authUserQueryKey()) || {
		name: '',
		email: '',
		org_id: '',
		org_name: ''
	};

export const setUserData = (userData: AuthResponse | null) =>
	queryClient.setQueryData<AuthResponse | null>(authUserQueryKey(), userData);
