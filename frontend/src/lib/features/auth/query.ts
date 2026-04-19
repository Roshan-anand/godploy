import { api } from '@/axios';
import type { AuthResponse } from './types';

// The current authenticated user is fetched through this shared query key and fetcher.
export function getAuthUserQueryKey() {
	return ['auth', 'user'] as const;
}

export function fetchAuthUser() {
	return api.get<AuthResponse>('/auth/user').then((res) => res.data);
}
