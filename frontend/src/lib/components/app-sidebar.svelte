<script lang="ts">
	import {
		Sidebar,
		SidebarContent,
		SidebarFooter,
		SidebarGroup,
		SidebarGroupContent,
		SidebarGroupLabel,
		SidebarHeader,
		SidebarMenu,
		SidebarMenuButton,
		SidebarMenuItem,
		SidebarRail
	} from '@/components/ui/sidebar';
</script>

<Sidebar>
	<SidebarHeader>
		<SidebarMenu>
			<SidebarMenuItem>
				<SidebarMenuButton size="lg">
					<div
						class="flex aspect-square size-8 items-center justify-center rounded-lg bg-sidebar-primary text-sidebar-primary-foreground"
					>
						<!-- Replace with your icon component -->
						<span class="size-4">◆</span>
					</div>
					<div class="grid flex-1 text-left text-sm leading-tight">
						<span class="truncate font-semibold">Acme Inc</span>
						<span class="truncate text-xs">Enterprise</span>
					</div>
				</SidebarMenuButton>
			</SidebarMenuItem>
		</SidebarMenu>
	</SidebarHeader>
	<SidebarContent>
		<SidebarGroup>
			<SidebarGroupLabel>Platform</SidebarGroupLabel>
			<SidebarGroupContent>
				<SidebarMenu>
					<SidebarMenuItem>
						<SidebarMenuButton>
							<!-- Replace with your icon component -->
							<span>⌂</span>
							<span>Home</span>
						</SidebarMenuButton>
					</SidebarMenuItem>
				</SidebarMenu>
			</SidebarGroupContent>
		</SidebarGroup>
	</SidebarContent>
	<SidebarFooter />
	<SidebarRail />
</Sidebar>
