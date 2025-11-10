import type { Handle } from '@sveltejs/kit';
import { API_URL } from '$env/static/private';
import { dev } from '$app/environment';

export const handle: Handle = async ({ event, resolve }) => {
	if (dev && event.url.pathname === '/.well-known/appspecific/com.chrome.devtools.json') {
		return new Response(null, { status: 404 });
	}

	const response = await resolve(event);

	// Set security headers properly
	response.headers.set('Content-Security-Policy', `script-src   ${API_URL}`);
	response.headers.set('Access-Control-Allow-Origin', dev ? '*' : API_URL);

	return response;
};
