<script lang="ts">
	import CreateBtn from '@/components/CreateBtn.svelte';
	import { Input } from '@/components/ui/input';
	import { Label } from '@/components/ui/label';
	import { Search } from '@lucide/svelte';
	import { setServicePageUiState } from '@/components/services/context.svelte';
	import { userState } from '@/store/userState.svelte';
	import CreateServiceDialog from './CreateServiceDialog.svelte';

	let { children } = $props();

	// shared UI state for this page (search, dialog open/close)
	const pageUi = setServicePageUiState();
</script>

<nav class="flex gap-2">
	<div class="flex-1 flex relative">
		<Input
			id="service-search"
			placeholder="Search for services"
			class="p-2"
			bind:value={pageUi.searchQuery}
		/>
		<Label class="absolute top-0 right-0 m-1 opacity-75" for="service-search"><Search /></Label>
	</div>
	<CreateBtn onclick={pageUi.openCreateDialog} disabled={userState.currentOrg.id === ''} />
</nav>

<CreateServiceDialog {pageUi} />

{@render children()}
