<script>
    import Main from '$lib/pages/Main.svelte';
    import SelectVault from './modals/SelectVault.svelte';
    import { vaultPath } from './stores';
    import { WindowGetSize, WindowIsMaximised, WindowSetSize } from '$wails/runtime/runtime';

    /** @type {import('$wails/runtime/runtime').Size} */
    let lastSize;

    let timeoutId;

    /**
     * Responds to window resize events.
     */
    function onResize() {
        clearTimeout(timeoutId);
        timeoutId = setTimeout(onResizeEnd, 500);
    }

    /**
     * A helper to respond to when a resize finishes.
     */
    async function onResizeEnd() {
        try {
            const maximized = await WindowIsMaximised();
            if(maximized) return;

            const size = await WindowGetSize();
            const restored = JSON.stringify(size) === JSON.stringify(lastSize);

            console.log(`${window.innerWidth}x${window.innerHeight}`, lastSize, '=>', size, 
                `restored: ${restored}`);

            lastSize = size;
        } catch(error) {
            console.error('Unable to get window size.', error);
        }
    }
</script>

<svelte:window onresize={onResize} />

{#if !$vaultPath}
    <SelectVault />
{:else}
    <Main />
{/if}
