<script lang="ts">
	import ServiceList from '@/components/services/service-list.svelte';
	import { userState } from '@/store/userState.svelte';

	const currentOrgId = $derived(userState.currentOrg.id);
</script>

<ServiceList scopeType="org" scopeId={currentOrgId} />
