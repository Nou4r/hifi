<script lang="ts">
	import * as Empty from '$lib/components/ui/empty/index.js';
	import Textarea from './Textarea.svelte';

	const uid = $props.id();

	let apiUrl = $state('');

	function handleSubmit(event: Event) {
		event.preventDefault();
	}

	let status = $state('error');
</script>

<div class=" *:not-first:mt-4 mt-8">
	<Textarea />
</div>
