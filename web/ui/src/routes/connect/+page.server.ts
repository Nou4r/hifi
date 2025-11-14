import type { PageServerLoad } from './$types';
import { redirect } from '@sveltejs/kit';

import { superValidate, message } from 'sveltekit-superforms';
import { zod4 } from 'sveltekit-superforms/adapters';
import { updateSchema } from '$lib/types/auth';

import { fail, type Actions } from '@sveltejs/kit';
import { API_URL } from '$env/static/private';

export const load: PageServerLoad = async (event) => {
	const { locals } = event;

	if (!locals.user) {
		const redirectTo = '/signin';
		throw redirect(303, redirectTo);
	}

	const form = await superValidate(event, zod4(updateSchema));

	return {
		form,
		user: locals.user
	};
};

export const actions: Actions = {
	default: async (e) => {
		const { locals } = e;

		const form = await superValidate(e, zod4(updateSchema));
		if (!form.valid) return fail(400, { form });

		if (locals.user?.password !== form.data.oldpassword) {
			form.valid = false;
			form.errors.oldpassword = ['Old password is required'];
			return fail(400, { form });
		}

		const token = e.cookies.get('hifi');

		const res = await e.fetch(`${API_URL}/v1/update`, {
			method: 'POST',
			headers: { 'Content-Type': 'application/json', Authorization: `Bearer ${token}` },
			body: JSON.stringify(form.data)
		});

		if (!res.ok) {
			form.valid = false;
			form.errors.username = ['Invalid deactivation'];
			return fail(400, { form });
		}

		return message(form, 'Username changed successfully');
	}
};
