import { fetchAPI } from "$lib/_core";
import { redirect } from "@sveltejs/kit";
import type { PageServerLoad } from "./$types";
import type { Actions } from "./$types";


export const load:PageServerLoad = ({cookies}) => {
    const userCookie = cookies.get("user")
    if(!userCookie){
        return {
            user: undefined
        }
    }
    return {
        user: JSON.parse(userCookie)
    }
}

export const actions: Actions = {
    login: async ({cookies, request}) =>{
        const data = await request.formData()
        const result = await fetchAPI("/user/connect", "POST", data)
        if(result.error){
            return {
                error: result.error
            }
        }
        cookies.set("user", JSON.stringify(result), {path: "/"})
        return result
    },
    logout: async({cookies}) => {
        cookies.delete("user", {path: "/"})
        throw redirect(302, "/")
    }
} satisfies Actions