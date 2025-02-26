<script>
    import Main from '$lib/pages/Main.svelte';
    import SelectVault from './modals/SelectVault.svelte';
    import { vaultPath } from './stores';
</script>

{#if !$vaultPath}
    <SelectVault />
{:else}
    <Main />
{/if}
