// Shared auth contracts live here so route modules only wire UI to the feature layer.
export interface LoginPayload {
	email: string;
	password: string;
}

export interface RegisterPayload {
	name: string;
	email: string;
	password: string;
	org_name: string;
}

export interface Organization {
	id: string;
	name: string;
}

export interface AuthResponse {
	message: string;
	name: string;
	email: string;
	org_id: string;
	org_name: string;
}
