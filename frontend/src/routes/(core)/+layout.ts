import { resolve } from '$app/paths';
import { fetchAuthUser, getAuthUserQueryKey } from '@/features/auth/query';
import { queryClient } from '@/query';
import { userState } from '@/store/userState.svelte';
import { redirect } from '@sveltejs/kit';
import axios from 'axios';

export async function load() {
	if (userState.isAuth) return;

	try {
		const authUser = await queryClient.fetchQuery({
			queryKey: getAuthUserQueryKey(),
			queryFn: fetchAuthUser
		});
		userState.setUser(authUser);
	} catch (err) {
		if (axios.isAxiosError(err) && err.response?.status === 403)
			redirect(302, resolve('/register'));
		redirect(302, resolve('/login'));
	}
}
