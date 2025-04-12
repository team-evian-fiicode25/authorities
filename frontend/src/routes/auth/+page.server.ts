import { fail, redirect } from "@sveltejs/kit";
import type { Actions, PageServerLoad } from "./$types";
import { isSessionType } from "$lib/types/SessionType";
import { getSession, setSession } from "$lib/SessionStore";
import { env } from "$env/dynamic/private";

export const actions = {
    default: async ({cookies, request}) => {
        const data = await request.formData();

        const username = data.get("username");
        const password = data.get("password");

        const res = await fetch(`${env.BACKEND_URL}/auth/login`, {
            method: "POST",
            body: JSON.stringify({
                username: username,
                password: password,
            },)
        });

        if (res.status == 401) {
            return fail(401, {username, error: "credentials"});
        }

        if (res.status >= 500) {
            return fail(500, {username, error: "unexpected"})
        }

        const session = await res.json()

        if(!isSessionType(session))
            throw new TypeError("Received wrong SessionType from backend");

        setSession(cookies, session);
        
        redirect(303, "/home");
    }
} satisfies Actions;

export const load: PageServerLoad = ({ cookies }) => {
    const session = getSession(cookies);

    if (session) {
        redirect(302, "/home");
    }
}
