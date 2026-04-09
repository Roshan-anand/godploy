import { resolve } from '$app/paths';
import { api } from '@/axios';
import type { AuthResponse } from '@/composables/useAuth';
import { queryClient } from '@/query';
import { userState } from '@/store/user-state.svelte';
import { redirect } from '@sveltejs/kit';
import axios from 'axios';

export async function load() {
	if (userState.isAuth) return;

	try {
		const res = await queryClient.fetchQuery({
			queryKey: ['auth', 'user'],
			queryFn: () => api.get<AuthResponse>('/auth/user')
		});
		const { email, name, org_id, org_name } = res.data;
		userState.email = email;
		userState.name = name;
		userState.currentOrg = {
			id: org_id,
			name: org_name
		};
		userState.isAuth = true;
	} catch (err) {
		if (axios.isAxiosError(err) && err.response?.status === 403)
			redirect(302, resolve('/register'));
		redirect(302, resolve('/login'));
	}
}
