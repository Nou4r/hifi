<script lang="ts">
	import Delete from './Delete.svelte';
</script>

<Delete />

<!-- <div class="flex gap-2">
	<Delete />
	<Dialog.Root bind:open>
		<Dialog.Trigger class={cn('cursor-pointer', buttonVariants({ variant: 'destructive' }))}>
			Delete Account
		</Dialog.Trigger>
		<Dialog.Content class="bg-zinc-900">
			<div class="flex flex-col items-center gap-2">
				<div
					class="border-border flex size-9 shrink-0 items-center justify-center rounded-full border"
					aria-hidden="true"
				>
					<UserPen class="text-white opacity-80" size={32} />
				</div>
				<Dialog.Header>
					<Dialog.Title class=" text-gray-300 sm:text-center">HiFi</Dialog.Title>
					<Dialog.Description class="text-gray-400 sm:text-center">
						Update your HiFi account
					</Dialog.Description>
				</Dialog.Header>
			</div>

			<form method="POST" class="space-y-5" use:enhance>
				<div class="space-y-4">
					<div class="space-y-2">
						
					</div>
				</div>
				<Form.Button
					class="mt-2 w-full cursor-pointer "
					type="submit"
					variant="outline"
					disabled={$submitting}
					>{#if $submitting}
						<Loader2 class="size-4 animate-spin" />
					{:else}
						Update Profile
					{/if}
				</Form.Button>
			</form>
		</Dialog.Content>
	</Dialog.Root>
</div> -->
