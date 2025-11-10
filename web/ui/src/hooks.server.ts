import type { Handle } from '@sveltejs/kit';
import { API_URL } from '$env/static/private';

export const handle: Handle = async ({ event, resolve }) => {
	const response = await resolve(event);

	response.headers.set('Content-Security-Policy', 'script-src' + API_URL);
	response.headers.set('Access-Control-Allow-Origin', API_URL);

	return response;
};
