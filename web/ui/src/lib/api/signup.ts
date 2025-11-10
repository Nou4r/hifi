export async function signup(data: Record<string, string>) {
	const res = await fetch('http://localhost:5002/v1/signup', {
		method: 'POST',
		body: JSON.stringify(data)
	});

	const out = await res.json().catch(() => ({}));
	return { ok: res.ok, data: out };
}
