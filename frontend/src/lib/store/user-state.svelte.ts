import type { Organization } from '@/composables/useAuth';

type UserState = {
	name: string;
	email: string;
	orgs: Organization[];
	isAuth: boolean;
};

export const userState = $state<UserState>({
	name: '',
	email: '',
	orgs: [],
	isAuth: false
});
