import { clearSession } from "$lib/SessionStore";
import { redirect } from "@sveltejs/kit";
import type { PageServerLoad } from "./$types";

export const load: PageServerLoad = ({cookies}) => {
    clearSession(cookies);

    redirect(303, "/auth");
}
