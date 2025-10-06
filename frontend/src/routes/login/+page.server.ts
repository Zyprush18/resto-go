import { API_URL } from '$env/static/private';
import { redirect } from '@sveltejs/kit';
import type { Actions } from './$types';

let response;
let content;

export const actions: Actions = {
	default: async (event) => {
		const formData = Object.fromEntries(await event.request.formData());

		const bodyReq = JSON.stringify({
			email: formData.email,
			password: formData.password
		});
		try {
			response = await fetch(API_URL + '/api/login', {
				method: 'POST',
				headers: {
					Accept: 'application/json',
					'Content-Type': 'application/json'
				},
				body: bodyReq
			});

			content = await response.json()
		} catch (error: any) {
			console.log(error);
			return {
				status: error.response.status,
				message: error.response.message
			}
		}
		

		if (response.status === 200) {
			console.log('berhasil login');

			// set cookies
			event.cookies.set('auth_token', content.token, {
				path: '/admin'
			})

			throw redirect(302,'/admin')
		}

	}
};
