export interface SessionType {
    sessionId: string
    login: LoginType
}

export function isSessionType(o: any): o is SessionType {
    return "sessionId" in o && typeof(o.sessionId) == "string" &&
        "login" in o && isLoginType(o.login);
}

export interface LoginType {
    username: string
}

function isLoginType(o: any): o is LoginType {
    return 'username' in o && typeof(o.username) == "string";
}
