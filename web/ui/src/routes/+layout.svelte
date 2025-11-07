<script lang="ts">
	import '../app.css';
	import favicon from '$lib/assets/favicon.svg';
	import '@fontsource-variable/inter';
	import { setContext } from 'svelte';

	let { children } = $props();

	let loggedIn = $state(false);

	function login() {
		loggedIn = true;
	}
	function logout() {
		loggedIn = false;
	}

	setContext('auth', {
		get loggedIn() {
			return loggedIn;
		},
		login,
		logout
	});
</script>

<svelte:head>
	<link rel="icon" href={favicon} />
</svelte:head>

{@render children?.()}
