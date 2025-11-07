<script lang="ts" module>
	import { z } from 'zod/v4';

	const formSchema = z.object({
		logo: z
			.string()
			.trim()
			.superRefine((val, ctx) => {
				if (!val) {
					ctx.addIssue({ code: 'custom', message: 'Username is required' });
					return;
				}
				if (val.length < 2) {
					ctx.addIssue({ code: 'custom', message: 'Username must be at least 2 characters long' });
					return;
				}
				if (val.length > 50) {
					ctx.addIssue({ code: 'custom', message: 'Username must not exceed 50 characters' });
					return;
				}
			}),
		title: z
			.string()
			.trim()
			.superRefine((val, ctx) => {
				if (!val) {
					ctx.addIssue({ code: 'custom', message: 'Password is required' });
					return;
				}
				if (val.length < 8) {
					ctx.addIssue({ code: 'custom', message: 'Password must be at least 8 characters long' });
					return;
				}
				if (val.length > 50) {
					ctx.addIssue({ code: 'custom', message: 'Password must not exceed 50 characters' });
					return;
				}
			}),
		description: z
			.string()
			.trim()
			.superRefine((val, ctx) => {
				if (!val) {
					ctx.addIssue({ code: 'custom', message: 'New Password is required' });
					return;
				}
				if (val.length < 8) {
					ctx.addIssue({
						code: 'custom',
						message: 'New Password must be at least 8 characters long'
					});
					return;
				}
				if (val.length > 50) {
					ctx.addIssue({ code: 'custom', message: 'New Password must not exceed 50 characters' });
					return;
				}
			})
	});
</script>

<script lang="ts">
	import { defaults, superForm } from 'sveltekit-superforms';
	import { zod4 } from 'sveltekit-superforms/adapters';
	import { Toaster, toast } from 'svelte-sonner';

	import Button from '$lib/components/ui/button.svelte';
	import * as Dialog from '$lib/components/ui/dialog';
	import { cn } from '$lib/utils';

	import * as Empty from '$lib/components/ui/empty/index.js';
	import ArrowUpRightIcon from '@lucide/svelte/icons/arrow-up-right';

	import IconSettingsFilled from '@lucide/svelte/icons/settings';
	import Loader2 from '@lucide/svelte/icons/loader-2';

	let open = $state(false);

	const form = superForm(defaults(zod4(formSchema)), {
		validators: zod4(formSchema),
		SPA: true,
		onUpdate: async ({ form: f }) => {
			const logoValue = f.data.logo?.trim() ?? '';

			if (!logoValue) {
				return;
			}
			if (f.valid) {
				await new Promise((r) => setTimeout(r, 500));
				console.log('Form data:', f.data);
				open = false;
				toast.success(`You submitted ${JSON.stringify(f.data, null, 2)}`);
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
			<IconSettingsFilled />
		</Empty.Media>
		<Empty.Title class=" text-gray-200">Settings</Empty.Title>
		<Empty.Description class="text-gray-400"
			>Authenticate with your Tidal account to get tokens that you can use with Selfhosted HiFi.</Empty.Description
		>
	</Empty.Header>
	<Empty.Content>
		<div class="flex gap-2">
			<Button class="flex cursor-pointer items-center gap-2" variant="outline">
				<svg xmlns="http://www.w3.org/2000/svg" width="32" height="32" viewBox="0 0 256 171">
					<path
						fill="#0a0b09"
						d="m128.004 85.339l42.664 42.67l-42.664 42.667l-42.669-42.667zM42.667.002L85.335 42.67L42.667 85.34L0 42.67zm170.666 0L256 42.67l-42.667 42.67l-42.666-42.67l-42.663 42.669l-42.669-42.67L128.004 0l42.663 42.665z"
					/>
				</svg>

				<span>Connect Tidal Account</span>
			</Button>
		</div>
	</Empty.Content>
	<Button variant="link" class="text-gray-400" size="sm">
		<a href="https://github.com/sachinsenal0x64/hifi">
			Learn More <ArrowUpRightIcon class="inline" />
		</a>
	</Button>
</Empty.Root>
