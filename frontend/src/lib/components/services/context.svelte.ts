import { getContext, setContext } from 'svelte';

class ServicePageUiState {
	searchQuery = $state('');
	createDialogOpen = $state(false);

	openCreateDialog = () => {
		this.createDialogOpen = true;
	};

	closeCreateDialog = () => {
		this.createDialogOpen = false;
	};
}

const KEY = 'service-page-ui-state';

export function setServicePageUiState() {
	return setContext(Symbol.for(KEY), new ServicePageUiState());
}

export function useServicePageUiState() {
	return getContext<ServicePageUiState>(Symbol.for(KEY));
}
