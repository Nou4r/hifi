import { z } from 'zod/v4';

export const formSchema = z.object({
	username: z
		.string()
		.regex(/^[a-zA-Z0-9_]+$/, 'Username can only contain letters and numbers')
		.trim()
		.superRefine((val, ctx) => {
			if (!val) {
				ctx.addIssue({ code: 'custom', message: 'Username is required' });
				return;
			}
			if (val.length < 2) {
				ctx.addIssue({ code: 'custom', message: 'Username must be at least 2 characters long' });
				return;
			}
			if (val.length > 50) {
				ctx.addIssue({ code: 'custom', message: 'Username must not exceed 50 characters' });
				return;
			}
		}),
	password: z
		.string()
		.regex(/^[a-zA-Z0-9_@]+$/, 'Password can only contain letters and numbers')
		.trim()
		.superRefine((val, ctx) => {
			if (!val) {
				ctx.addIssue({ code: 'custom', message: 'Password is required' });
				return;
			}
			if (val.length < 8) {
				ctx.addIssue({ code: 'custom', message: 'Password must be at least 8 characters long' });
				return;
			}
			if (val.length > 50) {
				ctx.addIssue({ code: 'custom', message: 'Password must not exceed 50 characters' });
				return;
			}
		})
});

export const formSchema2 = z.object({
	username: z
		.string()
		.regex(/^[a-zA-Z0-9_]+$/, 'Username can only contain letters and numbers')
		.trim()
		.superRefine((val, ctx) => {
			if (!val) {
				ctx.addIssue({ code: 'custom', message: 'Username is required' });
				return;
			}
			if (val.length < 2) {
				ctx.addIssue({ code: 'custom', message: 'Username must be at least 2 characters long' });
				return;
			}
			if (val.length > 50) {
				ctx.addIssue({ code: 'custom', message: 'Username must not exceed 50 characters' });
				return;
			}
		})
});

export const updateSchema = z.object({
	username: z
		.string()
		.regex(/^[a-zA-Z0-9_]+$/, 'Username can only contain letters and numbers')
		.trim()
		.superRefine((val, ctx) => {
			if (!val) {
				ctx.addIssue({ code: 'custom', message: 'Username is required' });
				return;
			}
			if (val.length < 2) {
				ctx.addIssue({ code: 'custom', message: 'Username must be at least 2 characters long' });
				return;
			}
			if (val.length > 50) {
				ctx.addIssue({ code: 'custom', message: 'Username must not exceed 50 characters' });
				return;
			}
		}),

	oldpassword: z
		.string()
		.regex(/^[a-zA-Z0-9_@]+$/, 'Password can only contain letters and numbers')
		.trim()
		.superRefine((val, ctx) => {
			if (!val) {
				ctx.addIssue({ code: 'custom', message: 'Password is required' });
				return;
			}
			if (val.length < 8) {
				ctx.addIssue({ code: 'custom', message: 'Password must be at least 8 characters long' });
				return;
			}
			if (val.length > 50) {
				ctx.addIssue({ code: 'custom', message: 'Password must not exceed 50 characters' });
				return;
			}
		}),
	password: z
		.string()
		.regex(/^[a-zA-Z0-9_@]+$/, 'Password can only contain letters and numbers')
		.trim()
		.superRefine((val, ctx) => {
			if (!val) {
				ctx.addIssue({ code: 'custom', message: 'Password is required' });
				return;
			}
			if (val.length < 8) {
				ctx.addIssue({ code: 'custom', message: 'Password must be at least 8 characters long' });
				return;
			}
			if (val.length > 50) {
				ctx.addIssue({ code: 'custom', message: 'Password must not exceed 50 characters' });
				return;
			}
		})
});
