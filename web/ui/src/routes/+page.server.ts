import { superValidate, message } from 'sveltekit-superforms';
import { zod4 } from 'sveltekit-superforms/adapters';
import { fail, type Actions } from '@sveltejs/kit';
import { formSchema } from '$lib/types/auth';
import { redirect } from '@sveltejs/kit';

export const load = async (event) => {
	const form = await superValidate(event, zod4(formSchema));
	return { form };
};

export const actions: Actions = {
	default: async (e) => {
		const form = await superValidate(e, zod4(formSchema));
		if (!form.valid) return fail(400, { form });
		const res = await e.fetch('http://localhost:5002/v1/signup', {
			method: 'POST',
			body: JSON.stringify(form.data)
		});
		const out = await res.json().catch(() => ({}));
		return res.ok
			? (message(form, out.message), redirect(303, '/signin'))
			: fail(res.status, out.message);
	}
};
