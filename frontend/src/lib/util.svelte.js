import { vaultPath, search, items } from './stores.svelte';
import { CloseVault } from '$wails/go/main/App';
import { BrowserOpenURL } from '$wails/runtime/runtime';

export const STORAGE_KEY_WIN_SIZE = 'avda:window-size';
export const STORAGE_KEY_RECENT_FILES = 'avda:recent-files';
export const STORAGE_KEY_ORDER = 'avda:sort-order';
export const STORAGE_KEY_SHOW_GROUPS = 'avda:display-groups';
export const STORAGE_KEY_AUTO_CLOSE = 'avda:auto-close';
export const STORAGE_KEY_AUTO_CLOSE_TIME = 'avda:auto-close-time';

/**
 * Attempts to format the code's characters into 
 * groupings of 3.
 * @param {string} code 
 * @returns The formatted code
 */
export function formatCode(code) {
    // Don't format codes that are under the minimum format length
    // or responses for code types that aren't implemented.
    if(code.length < 6 || code.includes('<!')) return code;

    const temp = code.split('');
    const formatted = [];

    for(let i=0; i < code.length; i++) {
        if(i !== 0 && i % 3 === 0 && temp.length > 1) {
            formatted.push(' ');
        }

        const char = temp.shift();
        formatted.push(char);
    }

    return formatted.join('');
}

/**
 * Captures a link click event and opens the external url
 * in the default browser.
 * @param event {Event}
 */
export function openExtUrl(event) {
    event.preventDefault();

    // @ts-ignore
    const url = event.currentTarget.href;
    BrowserOpenURL(url);
}

/**
 * Queries the GitHub API for info about the latest release.
 * @returns The release info
 */
export async function getLatestReleaseInfo() {
    try {
        /** @type {RequestInit} */
        const opts = {
            headers: {
                'Accept': 'application/vnd.github+json',
                'X-GitHub-Api-Version': '2022-11-28'
            }
        };

        const resp = await fetch('https://api.github.com/repos/sammy-t/avda/releases/latest', opts);
        const respJson = await resp.json();

        console.log(new Date, respJson);

        return respJson;
    } catch(error) {
        console.error('Unable to check for update.', error);
        return null;
    }
}

/**
 * Closes the current vault and clears the frontend's vault data. 
 * This triggers the selection modal to display.
 * @param {Event?} event
 */
export async function closeFile(event = null) {
    if(event) event.preventDefault();

    await CloseVault();

    vaultPath.set(null);
    items.set([]);
    search.set('');
}

/**
 * Generates a consistent color based on a group UUID
 * @param {string} uuid - The group UUID
 * @returns {string} - A CSS color string
 */
export function getGroupColor(uuid) {
    let hash = 0;
    for (let i = 0; i < uuid.length; i++) {
        hash = ((hash << 5) - hash) + uuid.charCodeAt(i);
        hash |= 0;
    }
    
    // Generate HSL color with fixed saturation and lightness for readability
    // Use the hash to determine the hue (0-360)
    const hue = Math.abs(hash % 360);
    return `hsl(${hue}, 70%, 60%)`;
}
