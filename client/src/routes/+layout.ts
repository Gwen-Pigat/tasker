import { browser } from "$app/environment";
import { url, user } from "$lib";
import type { LayoutLoad } from "./$types";

export const ssr: boolean = false;
export const prerender: boolean = false;

export const load: LayoutLoad = () => {
    if (browser) {
        const storedUser = localStorage.getItem("user");
        if (storedUser) {
            try {
                user.set(JSON.parse(storedUser));
            } catch (e) {
                console.error("Failed to parse stored user:", e);
                localStorage.removeItem("user");
            }
        }
    }
};