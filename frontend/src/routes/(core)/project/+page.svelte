<script lang="ts">
	import ServiceList from '@/components/services/service-list.svelte';
	import { getUserState } from '@/features/global/store.svelte';

	const { currentOrg } = getUserState();
	const currentOrgId = $derived(currentOrg.id);
</script>

<ServiceList scopeType="org" scopeId={currentOrgId} />
