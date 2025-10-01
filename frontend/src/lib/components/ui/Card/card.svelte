<script lang="ts">
	import { twMerge } from 'tailwind-merge'
	
	let { children, class: className, ...rest } = $props()
</script>

<div {...rest} class={twMerge('card bg-base-100 w-96 shadow-sm', className)}>
	{@render children?.()}
</div>