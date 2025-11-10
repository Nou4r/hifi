import { superValidate } from 'sveltekit-superforms';
import { zod4 } from 'sveltekit-superforms/adapters';
import { fail, type Actions } from '@sveltejs/kit';
import { formSchema } from '$lib/types/auth';
import { signup } from '$lib/api/signup';
import type { PageServerLoad } from './$types';
import { onMount } from 'svelte';

export const load: PageServerLoad = async (event) => {
	const form = await superValidate(event, zod4(formSchema));
	return { form };
};

export const actions: Actions = {
	default: async (e) => {
		const form = await superValidate(e, zod4(formSchema));
		if (!form.valid) return fail(400, { form });

		onMount(async () => {
			const ok = await signup(form.data);

			if (!ok) {
				return fail(400, 'Signup failed. Please try again.');
			}

			return 'Signup successful!';
		});
	}
};
