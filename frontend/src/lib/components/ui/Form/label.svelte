<script lang="ts">
    import {twMerge} from 'tailwind-merge'
    let props = $props()
</script>

<label {...props} class={twMerge("input validator", props.class)}>
    {@render props.children?.()}
</label>