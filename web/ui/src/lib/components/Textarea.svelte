<script lang="ts">
	import * as InputGroup from '$lib/components/ui/input-group/index.js';
	import * as Popover from '$lib/components/ui/popover/index.js';
	import CheckIcon from '@lucide/svelte/icons/check';
	import CopyIcon from '@lucide/svelte/icons/copy';
	import InfoIcon from '@lucide/svelte/icons/info';
	import StarIcon from '@lucide/svelte/icons/star';
	import { UseClipboard } from '$lib/hooks/use-clipboard.svelte.js';

	let isFavorite = $state(false);

	const clipboard = new UseClipboard();
</script>

<div class="grid w-full max-w-sm gap-6">
	<InputGroup.Root class="border border-zinc-700 bg-zinc-900">
		<InputGroup.Input
			class="placeholder:text-green-300"
			placeholder="https://x.com/shadcn"
			readonly
		/>
		<InputGroup.Addon align="inline-end">
			<InputGroup.Button
				aria-label="Copy"
				title="Copy"
				size="icon-xs"
				class="cursor-pointer hover:bg-zinc-300"
				onclick={() => clipboard.copy('https://x.com/shadcn')}
			>
				{#if clipboard.copied}
					<CheckIcon />
				{:else}
					<CopyIcon />
				{/if}
			</InputGroup.Button>
		</InputGroup.Addon>
	</InputGroup.Root>
</div>
