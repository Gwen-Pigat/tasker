import type { PageServerLoad } from "./$types";
import { env } from "$env/dynamic/private";

export const load: PageServerLoad = ({ cookies }: any) => {
    const userCookie = cookies.get("user");
    if (!userCookie) {
        return {
            user: undefined,
            apiURL: env.API_URL,
        }
    }
    return {
        user: JSON.parse(userCookie),
        apiURL: env.API_URL
    }
}   