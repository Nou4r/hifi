import type { PageServerLoad } from './$types';
import { redirect } from '@sveltejs/kit';

import { superValidate, message } from 'sveltekit-superforms';
import { zod4 } from 'sveltekit-superforms/adapters';
import { formSchema } from '$lib/types/auth';

import { fail, type Actions } from '@sveltejs/kit';
import { API_URL } from '$env/static/private';

export const load: PageServerLoad = async (event) => {
	const { locals } = event;

	if (!locals.user) {
		const redirectTo = '/signin';
		throw redirect(303, redirectTo);
	}

	const form = await superValidate(event, zod4(formSchema));

	return {
		form,
		user: locals.user
	};
};

export const actions: Actions = {
	default: async (e) => {
		const form = await superValidate(e, zod4(formSchema));
		console.log('form', form);
		if (!form.valid) return fail(400, { form });
		const res = await e.fetch(`${API_URL}/v1/delete`, {
			method: 'POST',
			headers: { 'Content-Type': 'application/json' },
			body: JSON.stringify(form.data)
		});

		if (!res.ok) {
			form.valid = false;
			form.errors.username = ['Invalid username'];
			return fail(400, { form });
		}

		return message(form, `Signup successful!`);
	}
};
