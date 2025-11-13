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
						<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 16 16"
							><path
								fill="white"
								fill-rule="evenodd"
								d="M15.53 1.53A.75.75 0 0 0 14.47.47l-1.29 1.29a4.24 4.24 0 0 0-5.423.483l-.58.58a.96.96 0 0 0 0 1.354l4.646 4.646a.96.96 0 0 0 1.354 0l.58-.58a4.24 4.24 0 0 0 .484-5.423zm-8.5 4.94a.75.75 0 0 1 0 1.06L5.78 8.78l1.44 1.44l1.25-1.25a.75.75 0 0 1 1.06 1.06l-1.25 1.25l.543.543a.96.96 0 0 1 0 1.354l-.58.58a4.24 4.24 0 0 1-5.423.484l-1.29 1.29A.75.75 0 0 1 .47 14.47l1.29-1.29a4.24 4.24 0 0 1 .483-5.423l.58-.58a.96.96 0 0 1 1.354 0l.543.543l1.25-1.25a.75.75 0 0 1 1.06 0M3.5 8.62l-.197.197a2.743 2.743 0 0 0 3.879 3.879l.197-.197zm9.197-1.439l-.197.197L8.621 3.5l.197-.197a2.743 2.743 0 0 1 3.879 3.879"
								clip-rule="evenodd"
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
