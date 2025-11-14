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
