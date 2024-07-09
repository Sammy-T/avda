<script>
    import Navbar from '$lib/components/Navbar.svelte';
    import Footer from '$lib/components/Footer.svelte';
    import EntryList from '$lib/components/EntryList.svelte';
    import { tweened } from 'svelte/motion';
    import { EventsOn } from '$wails/runtime/runtime';

    const countdown = tweened(30000);

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

    main {
        flex-grow: 1;
        display: flex;
        flex-direction: column;
        overflow: hidden;
    }
</style>
