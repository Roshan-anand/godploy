import type { Organization, AuthResponse } from '@/composables/useAuth';
type UserState = {
	name: string;
	email: string;
	currentOrg: Organization;
	orgs: Organization[];
	isAuth: boolean;
	setUser: (data: AuthResponse) => void;
};

export const userState = $state<UserState>({
	name: '',
	email: '',
	currentOrg: { id: '', name: '' },
	orgs: [],
	isAuth: false,

	setUser(data) {
		this.name = data.name;
		this.email = data.email;
		this.currentOrg = {
			id: data.org_id,
			name: data.org_name
		};
		this.isAuth = true;
	}
});
