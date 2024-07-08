<script>
    import closeFileIc from '$assets/images/file-minus.svg?raw';
</script>

<nav data-theme="dark">
    <ul>
        <li><strong>avdap</strong></li>
    </ul>

    <ul class="filename">
        <li><a href="##" class="contrast">filename</a></li>
    </ul>

    <ul>
        <li>
            <a href="##" class="contrast" data-tooltip="Close Vault" data-placement="bottom">
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
    }

    li strong {
        color: var(--pico-contrast);
    }

    a:hover {
        color: var(--pico-primary-background);
    }

    .filename a:hover {
        text-decoration: none;
    }
</style>
