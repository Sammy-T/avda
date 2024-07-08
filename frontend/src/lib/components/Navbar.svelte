<script>
    import closeFileIc from '$assets/images/file-minus.svg?raw';
</script>

<nav>
    <ul>
        <li><strong>avdap</strong></li>
    </ul>

    <ul class="filename">
        <li><a href="##">filename</a></li>
    </ul>

    <ul>
        <li>
            <a href="##" data-tooltip="Close Vault" data-placement="bottom">
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

    .filename a:hover {
        text-decoration: none;
    }
</style>
