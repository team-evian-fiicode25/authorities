import type { Cookies } from "@sveltejs/kit";
import type { SessionType } from "./types/SessionType";

export const getSession = (cookies: Cookies): SessionType | null => {
    try {
        return JSON.parse(atob(cookies.get("session") ?? ""));   
    } catch {
        return null;
    }
}

export const setSession = (cookies: Cookies, session: SessionType): SessionType => {
    cookies.set("session", btoa(JSON.stringify(session)), {path: "/"});
    return session;
}

export const clearSession = (cookies: Cookies) => {
    cookies.delete("session", { path: "/" });
}
