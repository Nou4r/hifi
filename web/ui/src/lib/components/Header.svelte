<script lang="ts">
	import Button from '$lib/components/ui/button.svelte';
	import { NotificationMenu, UserMenu } from '$lib/components/_extras/navbars';
	import { Popover, PopoverTrigger } from '$lib/components/ui/popover';
	import { page } from '$app/state';

	const title = 'Hifi';
	const { user } = $props();
</script>

<svelte:head>
	<title>{title}</title>
</svelte:head>

<header class="border-b px-4 md:px-6">
	<div class="flex h-16 items-center justify-between gap-4">
		<div class="flex items-center gap-2">
			<Popover>
				<PopoverTrigger>
					{#snippet child({ props })}
						<Button class="group size-8 md:hidden" variant="ghost" size="icon" {...props}>
							<svg
								class="pointer-events-none"
								width={16}
								height={16}
								viewBox="0 0 24 24"
								fill="none"
								stroke="currentColor"
								stroke-width="2"
								stroke-linecap="round"
								stroke-linejoin="round"
								xmlns="http://www.w3.org/2000/svg"
							>
								<path
									d="M4 12L20 12"
									class="group-aria-expanded:rotate-315 origin-center -translate-y-[7px] transition-all duration-300 [transition-timing-function:cubic-bezier(.5,.85,.25,1.1)] group-aria-expanded:translate-x-0 group-aria-expanded:translate-y-0"
								/>
								<path
									d="M4 12H20"
									class="origin-center transition-all duration-300 [transition-timing-function:cubic-bezier(.5,.85,0.25,1.8)] group-aria-expanded:rotate-45"
								/>
								<path
									d="M4 12H20"
									class="group-aria-expanded:rotate-135 origin-center translate-y-[7px] transition-all duration-300 [transition-timing-function:cubic-bezier(.5,.85,.25,1.1)] group-aria-expanded:translate-y-0"
								/>
							</svg>
						</Button>
					{/snippet}
				</PopoverTrigger>
			</Popover>
			<div class="flex items-center">
				<a
					href="/"
					aria-label="Home"
					title="Home"
					class=" text-primary hover:text-primary/90 rounded-full p-3 hover:bg-zinc-600 hover:opacity-80"
				>
					<svg
						xmlns="http://www.w3.org/2000/svg"
						aria-label="home"
						width="24"
						height="24"
						viewBox="0 0 24 24"
						><g
							fill="none"
							stroke="white"
							stroke-linecap="round"
							stroke-linejoin="round"
							stroke-width="2"
							><path d="M15 21v-8a1 1 0 0 0-1-1h-4a1 1 0 0 0-1 1v8" /><path
								d="M3 10a2 2 0 0 1 .709-1.528l7-6a2 2 0 0 1 2.582 0l7 6A2 2 0 0 1 21 10v9a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z"
							/></g
						></svg
					>
				</a>

				{#if user?.username && page.url.pathname !== '/connect'}
					<a
						href="/connect"
						aria-label="Connect account"
						data-sveltekit-reload
						title="Home"
						class=" text-primary hover:text-primary/90 rounded-full p-3 hover:bg-zinc-600 hover:opacity-80 md:hidden"
					>
						<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"
							><path
								fill="white"
								d="M19.49 5.57a6 6 0 0 1-1.893 8.962c-.649.351-1.43.135-1.952-.386l-5.79-5.791c-.522-.522-.738-1.303-.388-1.952A6 6 0 0 1 18.43 4.51l2.29-2.29a.75.75 0 1 1 1.061 1.06zm-2.017 7.263a4.5 4.5 0 1 0-6.306-6.306c-.266.349-.186.833.125 1.143l5.038 5.038c.31.31.794.391 1.143.125M3.28 21.78l2.29-2.29a6 6 0 0 0 8.962-1.893c.351-.649.135-1.43-.387-1.952l-5.79-5.79c-.522-.522-1.303-.738-1.952-.388A6 6 0 0 0 4.51 18.43l-2.29 2.29a.75.75 0 1 0 1.06 1.061m4.39-10.488l5.038 5.038c.31.31.39.794.125 1.143a4.5 4.5 0 1 1-6.306-6.306c.349-.266.833-.186 1.143.125"
							/></svg
						>
					</a>
				{/if}

				{#if user?.username && page.url.pathname !== '/deactivate'}
					<a
						href="/deactivate"
						aria-label="Deactivate account"
						data-sveltekit-reload
						title="Home"
						class=" text-primary hover:text-primary/90 rounded-full p-3 hover:bg-zinc-600 hover:opacity-80 md:hidden"
					>
						<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"
							><g fill="white" fill-rule="evenodd" clip-rule="evenodd"
								><path
									d="M16 9a4 4 0 1 1-8 0a4 4 0 0 1 8 0m-2 0a2 2 0 1 1-4 0a2 2 0 0 1 4 0"
								/><path
									d="M12 1C5.925 1 1 5.925 1 12s4.925 11 11 11s11-4.925 11-11S18.075 1 12 1M3 12c0 2.09.713 4.014 1.908 5.542A8.99 8.99 0 0 1 12.065 14a8.98 8.98 0 0 1 7.092 3.458A9 9 0 1 0 3 12m9 9a8.96 8.96 0 0 1-5.672-2.012A6.99 6.99 0 0 1 12.065 16a6.99 6.99 0 0 1 5.689 2.92A8.96 8.96 0 0 1 12 21"
								/></g
							></svg
						>
					</a>
				{/if}
			</div>
		</div>
		<!-- Right side -->
		<div class="flex items-center gap-4">
			<div class="flex items-center gap-2">
				<!-- Notification -->
				<NotificationMenu />
			</div>
			<!-- User menu -->
			<UserMenu {user} />
		</div>
	</div>
</header>
