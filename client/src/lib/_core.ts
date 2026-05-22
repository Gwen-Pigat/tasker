import { user,url } from '$lib'
import { get } from 'svelte/store'
import { env } from '$env/dynamic/public'

export async function resetUser(){
    user.set(null)
}



export async function fetchAPI(
    path: string, 
    method: string,
    data?: any
): Promise<any> {
    const API_URL: string = env.PUBLIC_API_URL || get(url) || "http://localhost:3000"
    const headers = new Headers()
    const currentUser: any = get(user)

    if (currentUser && currentUser.token) {
        headers.append("Authorization", "Bearer " + currentUser.token)
    }

    const options: any = {
        method: method,
        headers: headers
    }

    if (data) {
        headers.append("Content-Type", "application/json")
        // If data is already a string, use it, otherwise stringify it
        // Note: SvelteKit formData should be converted to an object
        if (data instanceof FormData) {
            const obj = Object.fromEntries(data.entries())
            options.body = JSON.stringify(obj)
        } else {
            options.body = typeof data === 'string' ? data : JSON.stringify(data)
        }
    }

    try {
        const response = await fetch(API_URL + path, options)
        const result = await response.json()

        if (!response.ok) {
            throw result.error || "Request failed"
        }
        return result
    } catch (err: any) {
        console.error("API Error:", err)
        return {
            "error": err
        }
    }
}