<script>
    import Navbar from '$lib/components/Navbar.svelte';
    import Footer from '$lib/components/Footer.svelte';
    import EntryList from '$lib/components/EntryList.svelte';
</script>

<header>
    <Navbar />
    <progress value="0.5" />
</header>

<main>
    <EntryList />
</main>

<Footer />

<style>
    header {
        background-color: rgba(0, 0, 0, 0.264);
    }

    header > progress {
        margin-bottom: 0;
        border-radius: 0;
    }

    main {
        flex-grow: 1;
        display: flex;
        flex-direction: column;
        overflow: hidden;
    }
</style>
