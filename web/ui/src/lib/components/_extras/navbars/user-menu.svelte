<script lang="ts">
	import Button from '$lib/components/ui/button.svelte';

	import LogOutIcon from '@lucide/svelte/icons/log-out';
	import { BookOpenText, LogInIcon } from 'lucide-svelte';
	import { Avatar, AvatarFallback, AvatarImage } from '$lib/components/ui/avatar';
	import {
		DropdownMenu,
		DropdownMenuContent,
		DropdownMenuGroup,
		DropdownMenuItem,
		DropdownMenuLabel,
		DropdownMenuSeparator,
		DropdownMenuTrigger
	} from '$lib/components/ui/dropdowns';
	import ProfileIcon from '$lib/components/ProfileIcon.svelte';

	let loggedIn = $state(false);

	let guest = $state('Guest');
</script>

<DropdownMenu>
	<DropdownMenuTrigger>
		{#snippet child({ props })}
			<Button
				variant="ghost"
				class="h-auto cursor-pointer p-0 hover:bg-transparent hover:opacity-80"
				{...props}
			>
				<Avatar>
					<!-- <AvatarImage src="" alt="Profile image" /> -->
					<AvatarFallback>
						<ProfileIcon username={guest} />
					</AvatarFallback>
				</Avatar>
			</Button>
		{/snippet}
	</DropdownMenuTrigger>
	<DropdownMenuContent class="max-w-64 bg-zinc-700" align="end">
		<DropdownMenuLabel class="flex min-w-0  flex-col">
			{#if loggedIn}
				<span class="truncate text-sm font-medium text-zinc-100">Keith Kennedy</span>
			{:else}
				<span class="truncate text-sm font-medium text-zinc-100">{guest}</span>
			{/if}
		</DropdownMenuLabel>

		{#if loggedIn}
			<DropdownMenuSeparator class="bg-zinc-600" />
			<DropdownMenuItem class="cursor-pointer text-zinc-100 focus:bg-zinc-600 focus:text-white">
				<LogOutIcon size={16} class="opacity-80" aria-hidden="true" />
				<span>Sign out</span>
			</DropdownMenuItem>
		{/if}

		<DropdownMenuSeparator class="bg-zinc-600" />
		<DropdownMenuItem class="cursor-pointer text-zinc-100 focus:bg-zinc-600 focus:text-white">
			<LogInIcon size={16} class="opacity-80" aria-hidden="true" />
			<span>Sign in</span>
		</DropdownMenuItem>
	</DropdownMenuContent>
</DropdownMenu>
