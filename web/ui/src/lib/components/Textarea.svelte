<script lang="ts">
	import * as InputGroup from '$lib/components/ui/input-group/index.js';
	import CheckIcon from '@lucide/svelte/icons/check';
	import CopyIcon from '@lucide/svelte/icons/copy';
	import { UseClipboard } from '$lib/hooks/use-clipboard.svelte.js';
	import Label from './ui/label.svelte';
	import * as Password from '$lib/components/ui/password';

	const slug = (s: string) =>
		s
			.toLowerCase()
			.trim()
			.replace(/\s+/g, '-')
			.replace(/[^a-z0-9_-]/g, '');

	const baseFields = [
		{ label: 'Hostname', value: 'example.com' },
		{ label: 'Username', value: 'John' },
		{ label: 'Password', value: 'password123' }
	];

	const fields = baseFields.map((f, i) => ({
		...f,
		id: `field-${i}-${slug(f.label)}`
	}));

	const clipboards = fields.map(() => new UseClipboard());
	const isPassword = (label: string) => label.toLowerCase() === 'password';
</script>

<svelte:head>
	<style></style>
</svelte:head>

<div class="grid w-full max-w-sm gap-6">
	{#each fields as field, i}
		<div class="flex flex-col gap-2">
			<Label for={field.id} class="text-left text-sm text-zinc-400">{field.label}</Label>

			{#if isPassword(field.label)}
				<Password.Root>
					<Password.Input
						id={field.id}
						name={field.label}
						value={field.value}
						autocomplete="off"
						readonly
						class="border border-zinc-700 bg-zinc-900 text-zinc-300"
					>
						<Password.ToggleVisibility class="cursor-pointer" />
						<Password.Copy
							class="cursor-pointer hover:bg-zinc-700 hover:text-white"
							aria-label="Copy Password"
							title="Copy Password"
						/>
					</Password.Input>
				</Password.Root>
			{:else}
				<InputGroup.Root class="border border-zinc-700 bg-zinc-900">
					<InputGroup.Input
						id={field.id}
						name={field.label}
						value={field.value}
						autocomplete="off"
						readonly
						class="text-zinc-300"
					/>
					<InputGroup.Addon align="inline-end">
						<InputGroup.Button
							aria-label={`Copy ${field.label}`}
							title={`Copy ${field.label}`}
							size="icon-xs"
							autocomplete="current-password"
							class="cursor-pointer hover:bg-zinc-700 hover:text-white"
							onclick={() => clipboards[i].copy(field.value)}
						>
							{#if clipboards[i].copied}
								<CheckIcon />
							{:else}
								<CopyIcon />
							{/if}
						</InputGroup.Button>
					</InputGroup.Addon>
				</InputGroup.Root>
			{/if}
		</div>
	{/each}
</div>
