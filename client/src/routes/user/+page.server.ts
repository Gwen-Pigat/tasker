import { fetchAPI } from "$lib/_core";
import type { Actions } from "./$types";

export const actions: Actions = {
    login: async ({cookies, request}) =>{
        const data = await request.formData()
        const result = await fetchAPI("/user/connect", "POST", data)
        if(result.error){
            return result
        }
        cookies.set("user", result, {path: "/"})
        return result
    }
} satisfies Actions