<script>
    import Footer from '$lib/components/Footer.svelte';
import Navbar from '$lib/components/Navbar.svelte';
</script>

<header>
    <Navbar />
    <progress value="0.5" />
</header>

<main>
    <h1>Hello</h1>
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
    }
</style>
