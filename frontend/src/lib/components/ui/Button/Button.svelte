<script lang="ts">
	import { twMerge } from "tailwind-merge";


let props = $props()

</script>

<button type={props.type} {...props} class={twMerge('bg-base-100',props.class)}>
    {props.val}
</button>