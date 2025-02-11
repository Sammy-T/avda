import { vaultPath, search, items } from './stores';
import { CloseVault } from '$wails/go/main/App';
import { BrowserOpenURL } from '$wails/runtime/runtime';

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
