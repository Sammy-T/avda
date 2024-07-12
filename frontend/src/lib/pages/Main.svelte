<script>
    import Navbar from '$lib/components/Navbar.svelte';
    import Footer from '$lib/components/Footer.svelte';
    import EntryList from '$lib/components/EntryList.svelte';
    import { search } from '$lib/stores';
    import { tweened } from 'svelte/motion';
    import { writable } from 'svelte/store';
    import { setContext } from 'svelte';
    import { EventsOn } from '$wails/runtime/runtime';

    const countdown = tweened(30000);

    let searchInput;

    const displaySearch = writable(false);
    setContext('displaySearch', displaySearch);

    $: if($displaySearch) searchInput?.focus();

    function onSearchSubmit() {
        $displaySearch = false;
    }

    /**
     * Responds to key events on the search input.
     * @param event {KeyboardEvent}
     */
    function onSearchInputKey(event) {
        if(event.key === 'Escape') $displaySearch = false;
    }

    /**
     * Updates the tweened store for the countdown.
     * @param ttn {number} - The time until the next OTP refresh
     */
    async function updateCountdown(ttn) {
        // Immediately update to the current time-til-next value
        await countdown.update(() => ttn, { duration: 0 });

        // Update to 0 with a duration of the time-til-next value
        countdown.update(() => 0, { duration: ttn });
    }

    EventsOn("onTimeUpdated", updateCountdown);
</script>

<header>
    <Navbar />
    <progress max="30000" value={$countdown} />

    {#if $displaySearch}
        <form spellcheck="false" on:submit|preventDefault={onSearchSubmit}>
            <input type="search" name="search" placeholder="Search" autocomplete="off" 
                on:keyup={onSearchInputKey} bind:value={$search} bind:this={searchInput} />
        </form>
    {/if}
</header>

<main>
    <EntryList />
</main>

<Footer />

<style>
    header {
        background-color: rgba(56, 83, 122, 0.574);
    }

    header > progress {
        margin-bottom: 0;
        border-radius: 0;
    }

    /* Remove Pico's transitions so they don't conflict with Svelte's while tweening. */
    header > progress[value]::-webkit-progress-value {
        transition: none;
        -webkit-transition: none;
    }

    header > form {
        position: absolute;
        width: 100%;
        padding: calc(var(--pico-spacing) * 0.5);
        display: flex;
        justify-content: center;
        pointer-events: none;
    }

    header > form > * {
        width: auto;
        margin: 0;
        pointer-events: auto;
    }

    header > form > input {
        height: calc(1rem * var(--pico-line-height) + var(--pico-form-element-spacing-vertical) * 1 + var(--pico-border-width)* 2);
    }

    main {
        flex-grow: 1;
        display: flex;
        flex-direction: column;
        overflow: hidden;
    }
</style>
