import { superValidate, message } from 'sveltekit-superforms';
import { zod4 } from 'sveltekit-superforms/adapters';
import { fail, type Actions } from '@sveltejs/kit';
import { formSchema } from '$lib/types/auth';
import { API_URL } from '$env/static/private';
import type { PageServerLoad } from './$types';
import { redirect } from '@sveltejs/kit';

export const load: PageServerLoad = async (event) => {
	// Redirect authenticated users away from the signin page
	const { locals, url } = event;

	if (locals.user) {
		const redirectTo = url.searchParams.get('redirect') || '/connect';
		throw redirect(303, redirectTo);
	}

	const form = await superValidate(event, zod4(formSchema));
	return { form };
};

export const actions: Actions = {
	default: async (e) => {
		const { request, cookies, url } = e;

		const form = await superValidate(e, zod4(formSchema));
		if (!form.valid) return fail(400, { form });
		const res = await e.fetch(`${API_URL}/v1/signin`, {
			method: 'POST',
			headers: { 'Content-Type': 'application/json' },
			body: JSON.stringify(form.data)
		});

		if (!res.ok) {
			return fail(res.status, { form, error: 'Signin failed' });
		}

		return message(form, `Signin successful!`);
	}
};
