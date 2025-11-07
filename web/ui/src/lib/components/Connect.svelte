<script lang="ts">
	import * as Empty from '$lib/components/ui/empty/index.js';
	import Textarea from './Textarea.svelte';

	const uid = $props.id();

	let apiUrl = $state('');

	function handleSubmit(event: Event) {
		event.preventDefault();
		console.log('API URL submitted:', apiUrl);
	}

	let status = $state('error');
</script>

<Empty.Root>
	<Empty.Header>
		<Empty.Description class="text-white">
			<div class="*:not-first:mt-4">
				<Textarea />
			</div>
		</Empty.Description>
	</Empty.Header>
</Empty.Root>
