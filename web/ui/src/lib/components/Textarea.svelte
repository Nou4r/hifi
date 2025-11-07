<script lang="ts">
	import * as InputGroup from '$lib/components/ui/input-group/index.js';
	import CheckIcon from '@lucide/svelte/icons/check';
	import CopyIcon from '@lucide/svelte/icons/copy';
	import { UseClipboard } from '$lib/hooks/use-clipboard.svelte.js';
	import Label from './ui/label.svelte';

	const fields = [
		{ label: 'Hostname', value: 'example.com' },
		{ label: 'Username', value: 'John' },
		{ label: 'Password', value: 'password123' }
	];

	const clipboards = fields.map(() => new UseClipboard());
</script>

<div class="grid w-full max-w-sm gap-6">
	{#each fields as field, i}
		<div class="flex flex-col gap-2">
			<Label class="text-left text-sm text-zinc-400 ">{field.label}</Label>
			<InputGroup.Root class="border border-zinc-700 bg-zinc-900">
				<InputGroup.Input value={field.value} readonly class="text-zinc-300" />
				<InputGroup.Addon align="inline-end">
					<InputGroup.Button
						aria-label={`Copy ${field.label}`}
						title={`Copy ${field.label}`}
						size="icon-xs"
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
		</div>
	{/each}
</div>
