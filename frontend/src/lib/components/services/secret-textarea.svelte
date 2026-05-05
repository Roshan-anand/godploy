<script lang="ts">
	import { Button } from '@/components/ui/button';
	import { Label } from '@/components/ui/label';
	import { Textarea } from '@/components/ui/textarea';
	import { cn } from '@/utils';
	import { Eye, EyeOff } from '@lucide/svelte';
	import type { FocusEventHandler, FormEventHandler } from 'svelte/elements';

	let {
		title,
		name,
		value,
		onblur,
		oninput,
		submitPending
	}: {
		title: string;
		name: string;
		value: string;
		onblur?: FocusEventHandler<HTMLTextAreaElement> | null;
		oninput?: FormEventHandler<HTMLTextAreaElement> | null;
		submitPending: boolean;
	} = $props();

	let isVisible = $state(false);
</script>

<div class="rounded-md border bg-card text-card-foreground shadow-sm overflow-hidden">
	<div class="flex items-center justify-between border-b px-3 py-2">
		<Label for={name} class="font-medium text-sm">{title}</Label>
		<!-- TODO : create an info btn to show what it is used for -->
		<Button
			variant="ghost"
			size="sm"
			class="h-6 w-6 p-0"
			onclick={(e) => {
				e.preventDefault();
				isVisible = !isVisible;
			}}
		>
			{#if isVisible}
				<EyeOff class="h-4 w-4" />
			{:else}
				<Eye class="h-4 w-4" />
			{/if}
			<span class="sr-only">Toggle visibility</span>
		</Button>
	</div>
	<Textarea
		id={name}
		spellcheck={false}
		placeholder="KEY=value"
		class={cn('rounded-none h-55', !isVisible ? 'style-text-security-disc' : '')}
		{value}
		{onblur}
		{oninput}
		disabled={!isVisible || submitPending}
	/>
</div>

<style>
	:global(.style-text-security-disc) {
		-webkit-text-security: disc;
	}
</style>
