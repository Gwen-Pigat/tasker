import { writable, type Writable } from "svelte/store";

// place files you want to import through the `$lib` alias in this folder.
export interface User {
    id: number;
    username: string;
    token?: string;
}

export interface Task {
    id: number;
    userId: number;
    title: string;
    dateAdd: string;
    dateTo: string | null;
    isDone: boolean;
    isDeleted?: boolean; // UI-only flag
}

export const tasks: Writable<Task[]> = writable([])
export const user: Writable<User | null> = writable(null)
export const error: Writable<string> = writable("")
export const url: Writable<string> = writable("")