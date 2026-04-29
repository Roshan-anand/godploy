import { resolve } from '$app/paths';
import { api } from '@/axios';
import type { AuthResponse } from '@/features/auth/type';
import { queryClient } from '@/query';
import { getUserState } from '@/features/global/store.svelte';
import { redirect } from '@sveltejs/kit';
import axios from 'axios';

export async function load() {
	const { isAuth, setUser } = getUserState();
	if (isAuth) return;

	try {
		const res = await queryClient.fetchQuery({
			queryKey: ['auth', 'user'],
			queryFn: () => api.get<AuthResponse>('/auth/user')
		});
		setUser(res.data);
	} catch (err) {
		if (axios.isAxiosError(err) && err.response?.status === 403)
			redirect(302, resolve('/register'));
		redirect(302, resolve('/login'));
	}
}
