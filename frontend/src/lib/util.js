import { vaultPath, search, items } from './stores';
import { CloseVault } from '$wails/go/main/App';

/**
 * Attempts to format the code's characters into 
 * groupings of 3.
 * @param {string} code 
 * @returns The formatted code
 */
export function formatCode(code) {
    if(code.length < 6) return code;

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
 * Closes the current vault and clears the frontend's vault data. 
 * This triggers the selection modal to display.
 */
export async function closeFile() {
    await CloseVault();

    vaultPath.set(null);
    items.set([]);
    search.set('');
}
