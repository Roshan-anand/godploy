import { resolve } from '$app/paths';
import { fetchAuthUserQuery } from './query';
import { redirect } from '@sveltejs/kit';
import axios from 'axios';

export async function load() {
	try {
		await fetchAuthUserQuery();
	} catch (err) {
		if (axios.isAxiosError(err) && err.response?.status === 403)
			redirect(302, resolve('/register'));
		redirect(302, resolve('/login'));
	}
}
