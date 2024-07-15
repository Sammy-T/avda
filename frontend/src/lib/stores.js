import { writable } from 'svelte/store';

export const vaultPath = writable();
export const items = writable([]);
export const search = writable('');
