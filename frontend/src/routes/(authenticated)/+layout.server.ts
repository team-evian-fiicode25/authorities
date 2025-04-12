import { redirect } from "@sveltejs/kit"
import type { LayoutServerLoad } from "./$types"
import { type SessionType } from "$lib/types/SessionType";
import { getSession } from "$lib/SessionStore";


export const load: LayoutServerLoad = ({cookies}): { session: SessionType } => {
    var session = getSession(cookies);

    if (!session) {
        redirect(302, "/auth");
    }

    return { session };
}
