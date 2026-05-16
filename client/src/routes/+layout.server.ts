import type { PageServerLoad } from "./$types";
import { API_URL } from "$env/static/private";

export const load: PageServerLoad = ({ cookies }: any) => {
    const userCookie = cookies.get("user");
    if (!userCookie) {
        return {
            user: undefined,
            apiURL: API_URL,
        }
    }
    return {
        user: JSON.parse(userCookie),
        apiURL: API_URL
    }
}   