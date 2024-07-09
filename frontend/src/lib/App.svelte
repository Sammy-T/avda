<script>
    import Main from '$lib/pages/Main.svelte';
    import SelectVault from './modals/SelectVault.svelte';
    import { vaultPath } from './stores';
</script>

<Main />

{#if !$vaultPath}
    <SelectVault />
{/if}
