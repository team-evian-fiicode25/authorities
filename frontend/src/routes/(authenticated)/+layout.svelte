<script lang="ts">
    let { children } = $props();
</script>

<div class="navbar bg-base-200 flex items-center justify-between p-2 shadow-sm">
    <div>
        <span class="text-xl font-bold">RideMe Authorities</span>
    </div>
    <ul class="menu menu-horizontal px-1">
        <li><a href="/logout">Logout</a></li>
    </ul>
</div>
{@render children()}
