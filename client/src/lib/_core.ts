import { user, url, error } from '$lib'
import { get } from 'svelte/store'
import { env } from '$env/dynamic/public'
import { browser, dev } from '$app/environment'

export async function resetUser(){
    user.set(null)
}



export async function fetchAPI(
    path: string, 
    method: string,
    data?: any
): Promise<any> {
    const defaultAPI = dev ? "http://localhost:3000" : "https://tasker-api.orizenh.com"
    let publicAPI: string | undefined = env.PUBLIC_API_URL
    if (!dev && publicAPI && (publicAPI.includes("localhost") || publicAPI.includes("127.0.0.1"))) {
        publicAPI = undefined
    }
    let API_URL: string = (browser && localStorage.getItem("api_url")) || publicAPI || get(url) || defaultAPI
    if (!browser) {
        if (API_URL.startsWith("http://localhost:3000")) {
            API_URL = API_URL.replace("http://localhost:3000", "http://api:3000")
        } else if (API_URL.startsWith("http://127.0.0.1:3000")) {
            API_URL = API_URL.replace("http://127.0.0.1:3000", "http://api:3000")
        }
    }
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
        const errMsg = err instanceof Error ? err.message : String(err)
        if (browser) {
            error.set(errMsg)
        }
        return {
            "error": errMsg
        }
    }
}