import { fetchAPI } from "$lib/_core";
import { redirect } from "@sveltejs/kit";
import type { Actions } from "./$types";


export const actions: Actions = {
    login: async ({ cookies, request }) => {
        const data = await request.formData()
        const result = await fetchAPI("/user/connect", "POST", data)
        if (result.error) {
            return {
                "error": result.error,
                "view": "login"
            }
        }
        cookies.set("user", JSON.stringify(result), { path: "/" })
        throw redirect(302, "/")
    },
    register: async ({ cookies, request }) => {
        const data = await request.formData()
        const result = await fetchAPI("/user", "POST", data)
        if (result.error) {
            return {
                "error": result.error,
                "view": "register"
            }
        }
        cookies.set("user", JSON.stringify(result), { path: "/" })
        throw redirect(302, "/")
    },
    logout: async ({ cookies }) => {
        cookies.delete("user", { path: "/" })
        throw redirect(302, "/")
    }
} satisfies Actions