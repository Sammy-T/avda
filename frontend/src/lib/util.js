import { vaultPath } from './stores';
import { CloseVault } from '$wails/go/main/App';

/**
 * Closes the current vault and clears the vault path store. 
 * This triggers the selection modal to display.
 */
export async function closeFile() {
    await CloseVault();

    vaultPath.set(null);
}
