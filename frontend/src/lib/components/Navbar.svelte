<script>
    import Dropdown from './navbar/Dropdown.svelte';
    import avdaIc from '$assets/images/avda.svg?raw';
    import searchIc from '$assets/images/search.svg?raw';
    import groupsIc from '$assets/images/category-2.svg?raw';
    import sortIc from '$assets/images/sort-descending.svg?raw';
    import closeFileIc from '$assets/images/file-minus.svg?raw';
    import { groupsMap, order, selectedGroupUuid, vaultPath } from '$lib/stores.svelte';
    import { closeFile, getGroupColor, STORAGE_KEY_ORDER, STORAGE_KEY_SHOW_GROUPS } from '$lib/util.svelte';
    import { getContext, onMount } from 'svelte';
    import { GetGroups } from '$wails/go/main/App';

    const displaySearch = getContext('displaySearch');
    const displayGroups = getContext('displayGroups');

    let defaultSort = $state('custom');

    const sortItems = [
        { label: 'Custom', value: 'custom' },
        { label: 'Account (A to Z)', value: 'account-asc' },
        { label: 'Account (Z to A)', value: 'account-desc' },
        { label: 'Issuer (A to Z)', value: 'issuer-asc' },
        { label: 'Issuer (Z to A)', value: 'issuer-desc' },
    ];

    /**
     * @param {Event} event
     */
    function toggleSearch(event) {
        event.preventDefault();

        // @ts-ignore
        event.currentTarget.blur(); // Remove focus from the anchor element so the tooltip doesn't remain

        $displaySearch = !$displaySearch;
    }

    /**
     * @param {Event} event
     */
    function toggleGroups(event) {
        event.preventDefault();

        // @ts-ignore
        event.currentTarget.blur(); // Remove focus from the anchor element so the tooltip doesn't remain

        $displayGroups = !$displayGroups;

        localStorage.setItem(STORAGE_KEY_SHOW_GROUPS, $displayGroups);
    }

    /**
     * @param {String} selected
     */
    function onSortSelect(selected) {
        $order = selected;

        localStorage.setItem(STORAGE_KEY_ORDER, $order);
    }

    async function initGroups() {
        const storedShowGroups = localStorage.getItem(STORAGE_KEY_SHOW_GROUPS) === 'true';

        const groups = await GetGroups();

        const mapping = new Map();
        mapping.set(null, { uuid: null, name: "All", color: null });
        
        groups?.forEach(group => {
            mapping.set(group.uuid, {
                uuid: group.uuid,
                name: group.name,
                color: getGroupColor(group.uuid)
            });
        });

        $groupsMap = mapping;

        // Reset the currently selected group when loading a file without a matching group
        if(!groups?.find((g) => g.uuid === $selectedGroupUuid)) $selectedGroupUuid = null;

        $displayGroups = storedShowGroups;
    }

    function initSort() {
        const storedOrder = localStorage.getItem(STORAGE_KEY_ORDER);

        if(storedOrder) {
            defaultSort = storedOrder;
            $order = storedOrder;
        }
    }

    onMount(() => {
        initSort();
        initGroups();
    });
</script>

<nav data-theme="dark">
    <ul>
        <li>
            {@html avdaIc}
            <strong>avda</strong>
        </li>
    </ul>

    <ul class="filename">
        <li>
            <a href="##" class="contrast" onclick={closeFile}>
                {$vaultPath?.split(/[\\/]/).at(-1) ?? 'filename'}
            </a>
        </li>
    </ul>

    <ul>
        <li>
            <a href="##" class="contrast" data-tooltip="Search" data-placement="bottom" onclick={toggleSearch}>
                {@html searchIc}
            </a>
        </li>
        <li>
            <a href="##" class="contrast" data-tooltip="Groups" data-placement="bottom" onclick={toggleGroups}>
                {@html groupsIc}
            </a>
        </li>
        <li>
            <Dropdown name="sort" items={sortItems} selected={defaultSort} onselect={onSortSelect}>
                {@html sortIc}
            </Dropdown>
        </li>
        <li>
            <a href="##" class="contrast" data-tooltip="Close Vault" data-placement="bottom" onclick={closeFile}>
                {@html closeFileIc}
            </a>
        </li>
    </ul>
</nav>

<style>
    nav {
        /* 
        Why I have to account for the intentional overflow is beyond me.
        see: https://picocss.com/docs/nav#overflow
        see: https://github.com/picocss/pico/issues/549
        */
        padding: 0 calc(var(--pico-nav-element-spacing-horizontal) * 4);
        overflow-x: clip;
        display: grid;
        grid-template-columns: 1fr max-content 1fr;
    }

    ul:last-child {
        justify-content: end;
    }

    li strong {
        color: var(--pico-contrast);
    }

    a:hover {
        color: var(--pico-primary-background);
    }

    .filename {
        justify-content: center;

        & a:hover {
            text-decoration: none;
        }
    }
</style>
