import { superValidate, message } from 'sveltekit-superforms';
import { zod4 } from 'sveltekit-superforms/adapters';
import { fail, type Actions } from '@sveltejs/kit';
import { formSchema } from '$lib/types/auth';
import { API_URL } from '$env/static/private';

export const load = async (event) => {
	const sessionUser = event.locals.user;
	const form = await superValidate(event, zod4(formSchema));

	const albums = [
		{
			src: 'https://resources.tidal.com/images/ad522656/b4b6/4054/8b98/7ff39644cea6/640x640.jpg',
			alt: 'Album artwork for the track'
		},
		{
			src: 'https://resources.tidal.com/images/ad522656/b4b6/4054/8b98/7ff39644cea6/640x640.jpg',
			alt: 'Second album artwork'
		}
	];

	const titles = ['Album artwork for the track', 'Second album artwork'];

	return {
		form,
		user: sessionUser,
		albums,
		titles
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
