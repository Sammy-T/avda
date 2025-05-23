<script>
    import Navbar from '$lib/components/Navbar.svelte';
    import Footer from '$lib/components/Footer.svelte';
    import EntryList from '$lib/components/EntryList.svelte';
    import GroupFilter from '$lib/components/GroupFilter.svelte';
    import { search } from '$lib/stores.svelte';
    import { Tween } from 'svelte/motion';
    import { writable } from 'svelte/store';
    import { onDestroy, onMount, setContext } from 'svelte';
    import { EventsOn } from '$wails/runtime/runtime';
    import { GetTTN } from '$wails/go/main/App';
    import { closeFile, STORAGE_KEY_AUTO_CLOSE, STORAGE_KEY_AUTO_CLOSE_TIME } from '$lib/util.svelte';

    const countdown = new Tween(30000);
    let timeoutId;

    let searchInput = $state();

    const displaySearch = writable(false);
    setContext('displaySearch', displaySearch);

    const displayGroups = writable(false);
    setContext('displayGroups', displayGroups);

    $effect(() => {
        if($displaySearch) searchInput?.focus();
    });

    function onSearchSubmit() {
        $displaySearch = false;
    }

    /**
     * Responds to key events on the search input.
     * @param event {KeyboardEvent}
     */
    function onSearchInputKey(event) {
        event.preventDefault();

        if(event.key === 'Escape') $displaySearch = false;
    }

    /**
     * Updates the tweened store for the countdown.
     * @param ttn {number} - The time until the next OTP refresh
     */
    async function updateCountdown(ttn) {
        // Immediately update to the current time-til-next value
        countdown.set(ttn, { duration: 0 });

        // Update to 0 with a duration of the time-til-next value
        countdown.set(0, { duration: ttn });
    }

    EventsOn("onTimeUpdated", updateCountdown);

    onMount(async () => {
        const ttn = await GetTTN();
        updateCountdown(ttn);

        const autoClose = localStorage.getItem(STORAGE_KEY_AUTO_CLOSE) === 'true';
        const autoCloseTime = Number(localStorage.getItem(STORAGE_KEY_AUTO_CLOSE_TIME)) || 5;

        const timeout = autoCloseTime * 60 * 1000;

        if(autoClose) {
            timeoutId = setTimeout(closeFile, timeout);
        }
    });

    onDestroy(() => {
        clearTimeout(timeoutId);
    });
</script>

<header>
    <Navbar />
    <progress max="30000" value={countdown.current}></progress>

    {#if $displaySearch}
        <form spellcheck="false" onsubmit={onSearchSubmit}>
            <input type="search" name="search" placeholder="Search" autocomplete="off" 
                onkeyup={onSearchInputKey} bind:value={$search} bind:this={searchInput} />
        </form>
    {/if}
</header>

<main>
    {#if $displayGroups}
        <GroupFilter />
    {/if}

    <EntryList />
</main>

<Footer />

<style>
    header {
        background-color: rgba(56, 83, 122, 0.574);

        & > progress {
            margin-bottom: 0;
            border-radius: 0;
        }

        /* Remove Pico's transitions so they don't conflict with Svelte's while tweening. */
        & > progress[value]::-webkit-progress-value {
            transition: none;
            -webkit-transition: none;
        }

        & > form {
            position: absolute;
            width: 100%;
            padding: calc(var(--pico-spacing) * 0.5);
            display: flex;
            justify-content: center;
            pointer-events: none;

            & > * {
                width: auto;
                margin: 0;
                pointer-events: auto;
            }

            & > input {
                height: calc(1rem * var(--pico-line-height) + var(--pico-form-element-spacing-vertical) * 1 + var(--pico-border-width)* 2);
                z-index: 2;
            }
        }
    }

    main {
        flex-grow: 1;
        overflow-y: auto;
    }
</style>
