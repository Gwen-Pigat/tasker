import { url, user } from "$lib";
import type { LayoutLoad } from "./$types";

export const ssr: boolean = true;
export const prerender: boolean = false;

export const load: LayoutLoad = ({ data }) => {
    if (data) {
        if (data.apiURL) {
            url.set(data.apiURL);
        }
        user.set(data.user || null);
    }
    return data;
};