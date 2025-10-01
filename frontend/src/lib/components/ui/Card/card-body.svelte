<script lang="ts">
    import { twMerge } from 'tailwind-merge'
    let props = $props()
</script>

<div {...props} class={twMerge("card-body",props.class)}>
    {@render props.children?.()}
</div>