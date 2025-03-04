<script>
    import Main from '$lib/pages/Main.svelte';
    import SelectVault from './modals/SelectVault.svelte';
    import { vaultPath } from './stores';
    import { onMount } from 'svelte';
    import { WindowGetSize, WindowIsMaximised, WindowSetSize } from '$wails/runtime/runtime';
    import { STORAGE_KEY_WIN_SIZE } from './util';

    let windowInitFinished = false;

    /** @type {import('$wails/runtime/runtime').Size} */
    let lastSize;

    let timeoutId;

    /**
     * Responds to window resize events.
     */
    function onResize() {
        if(!windowInitFinished) {
            windowInitFinished = true;
            return;
        }

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

            if(!restored) localStorage.setItem(STORAGE_KEY_WIN_SIZE, JSON.stringify(size));

            lastSize = size;
        } catch(error) {
            console.error('Unable to get window size.', error);
        }
    }

    /**
     * Initializes the app window by checking for a stored window size
     * and resizing the window to it if one is found.
     */
    async function initWindow() {
        lastSize = await WindowGetSize();

        const storedSizeRaw = localStorage.getItem(STORAGE_KEY_WIN_SIZE);

        /** @type {import('$wails/runtime/runtime').Size} */
        const storedSize = JSON.parse(storedSizeRaw);

        console.log('Stored size', storedSize);

        if(!storedSize || JSON.stringify(lastSize) === storedSizeRaw) {
            windowInitFinished = true;
            return;
        }

        console.log('Resizing', lastSize, '=>', storedSize);

        WindowSetSize(storedSize.w, storedSize.h);

        lastSize = await WindowGetSize();
    }

    onMount(() => {
        initWindow();
    });
</script>

<svelte:window onresize={onResize} />

{#if !$vaultPath}
    <SelectVault />
{:else}
    <Main />
{/if}
