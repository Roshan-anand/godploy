import type { ServiceType } from '@/types.js';

export function load({ params, url }) {
	const queryString =url.hash.split('?')[1];
	const searchParams = new URLSearchParams(queryString);

	return {
		service: params.service as ServiceType,
		id: searchParams.get('id')
	};
}
