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
    isCommon: boolean; // Flag for tasks derived from common tasks
    isDeleted?: boolean; // UI-only flag
}

export interface CommonTask {
    id: number;
    title: string;
    dateAdd: string;
    refUser: number;
}

export interface DashboardStats {
    tasksAdded: number;
    tasksDone: number;
    avgDuration: number;
    lastNote: Note | null;
}

export interface Note {
    id: number;
    title: string;
    content: string;
    dateAdd: string;
    dateUpdate?: string;
    refUser: number;
}

export const tasks: Writable<Task[]> = writable([])
export const commonTasks: Writable<CommonTask[]> = writable([])
export const notes: Writable<Note[]> = writable([])
export const dashboardStats: Writable<DashboardStats | null> = writable(null)
export const user: Writable<User | null> = writable(null)
export const error: Writable<string> = writable("")
export const url: Writable<string> = writable("")