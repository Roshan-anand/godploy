import { getContext, setContext } from 'svelte';

interface OrgState {
	name: string;
	id: string;
	setOrg: (id: string, name: string) => void;
}

class OrgStateClass implements OrgState {
	constructor(id: string, name: string) {
		this.id = id;
		this.name = name;
	}

	id = $state('');
	name = $state('');

	setOrg = (id: string, name: string) => {
		this.id = id;
		this.name = name;
	};
}

const DEFAULT_KEY = 'org:state';

export const getCurrentOrgState = (key: string = DEFAULT_KEY) =>
	getContext<OrgState>(Symbol.for(key));

export const setCurrentOrgState = (id: string, name: string, key: string = DEFAULT_KEY) => {
	setContext(Symbol.for(key), new OrgStateClass(id, name));
};
