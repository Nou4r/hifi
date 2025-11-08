<script lang="ts">
	import { superForm } from 'sveltekit-superforms';
	import { zod4 } from 'sveltekit-superforms/adapters';
	import { Toaster, toast } from 'svelte-sonner';
	import * as Form from '$lib/components/ui/form/index.js';
	import { formSchema } from '$lib/types/auth';
	import { goto } from '$app/navigation';

	import Button, { buttonVariants } from '$lib/components/ui/button.svelte';
	import Input from '$lib/components/ui/input.svelte';
	import * as Dialog from '$lib/components/ui/dialog';
	import { cn } from '$lib/utils';

	import * as Empty from '$lib/components/ui/empty/index.js';
	import IconUserCircle from '@tabler/icons-svelte/icons/user-circle';
	import ArrowUpRightIcon from '@lucide/svelte/icons/arrow-up-right';
	import Loader2 from '@lucide/svelte/icons/loader-2';

	let open = $state(false);

	const { data } = $props();

	const form = superForm(data.form, {
		resetForm: true,
		validators: zod4(formSchema),
		onResult: ({ result }) => {
			if (result.type === 'success') {
				open = false;
				toast.promise(
					new Promise((resolve) => {
						setTimeout(resolve, 800);
					}),
					{
						loading: 'Account created successfully!',
						success: () => {
							setTimeout(() => {
								goto('/signin');
							}, 800);
							return 'Redirecting to Sign In...';
						},

						error: 'Something went wrong. Please try again.'
					}
				);
			} else {
				open = false;
				toast.error('Something went wrong. Please try again.');
			}
		}
	});

	const { form: formData, submitting, enhance } = form;
</script>

<Toaster closeButton position="top-center" />

<Empty.Root>
	<Empty.Header>
		<Empty.Media variant="icon">
			<IconUserCircle />
		</Empty.Media>
		<Empty.Title class=" text-gray-200">Let’s Get Started</Empty.Title>
		<Empty.Description class="text-gray-400">
			Get started by creating your HiFi account, and you’ll be ready to listen to TIDAL music.
		</Empty.Description>
	</Empty.Header>
	<Empty.Content>
		<div class="flex gap-2">
			<Dialog.Root bind:open>
				<Dialog.Trigger class={cn('cursor-pointer', buttonVariants({ variant: 'outline' }))}
					>Get Started</Dialog.Trigger
				>
				<Dialog.Content class="bg-zinc-900">
					<div class="flex flex-col items-center gap-2">
						<Dialog.Header>
							<Dialog.Title class="mt-10 text-gray-300 sm:text-center">HiFi</Dialog.Title>
							<Dialog.Description class="text-gray-400 sm:text-center">
								Create your HiFi account
							</Dialog.Description>
						</Dialog.Header>
					</div>

					<form method="POST" class="space-y-5" use:enhance>
						<div class="space-y-4">
							<div class="space-y-2">
								<Form.Field {form} name="username">
									<Form.Control>
										{#snippet children({ props })}
											<Form.Label class="font-bold text-gray-300">Username</Form.Label>
											<Input
												class="border-zinc-700 text-white"
												placeholder="Joe Doe"
												type="text"
												{...props}
												bind:value={$formData.username}
											/>
										{/snippet}
									</Form.Control>
									<Form.FieldErrors />
								</Form.Field>
							</div>
							<div class="space-y-2">
								<Form.Field {form} name="password">
									<Form.Control>
										{#snippet children({ props })}
											<Form.Label class="font-bold text-gray-300">Password</Form.Label>
											<Input
												class="border-zinc-700 text-white"
												placeholder="Secure Password"
												type="password"
												{...props}
												bind:value={$formData.password}
											/>
										{/snippet}
									</Form.Control>
									<Form.FieldErrors />
								</Form.Field>
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
								Create Account
							{/if}
						</Form.Button>
					</form>
				</Dialog.Content>
			</Dialog.Root>
		</div>
	</Empty.Content>
	<Button variant="link" class="text-gray-400" size="sm">
		<a href="https://github.com/sachinsenal0x64/hifi">
			Learn More <ArrowUpRightIcon class="inline" />
		</a>
	</Button>
</Empty.Root>
