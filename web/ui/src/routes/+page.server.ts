import { superValidate, message } from 'sveltekit-superforms';
import { zod4 } from 'sveltekit-superforms/adapters';
import { fail, type Actions } from '@sveltejs/kit';
import { formSchema } from '$lib/types/auth';
import { API_URL } from '$env/static/private';

export const load = async (event) => {
	const sessionUser = event.locals.user;
	const form = await superValidate(event, zod4(formSchema));

	return {
		form,
		user: sessionUser
	};
};

export const actions: Actions = {
	default: async (e) => {
		const form = await superValidate(e, zod4(formSchema));
		if (!form.valid) return fail(400, { form });
		const res = await e.fetch(`${API_URL}/v1/signup`, {
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
