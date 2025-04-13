import { env } from "$env/dynamic/private";
import { clearSession } from "$lib/SessionStore";
import { redirect } from "@sveltejs/kit";
import type { PageServerLoad } from "../$types";
import type { SessionType } from "$lib/types/SessionType";

export const load: PageServerLoad = async ({cookies, parent}): Promise<{ session: SessionType, message: string }> => {

    const data = await parent();
    const session = data.session;

    const res = await fetch(`${env.BACKEND_URL}/hello`, {
        headers: [["Authorization", `Bearer ${session.sessionId}`]]
    });

    if (res.status == 401) {
        clearSession(cookies);
        redirect(302, "/auth");
    }

    const { message } = await res.json();

    return {session, message};
};
