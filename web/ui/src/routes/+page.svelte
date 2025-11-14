<script lang="ts">
	import Header from '$lib/components/Header.svelte';
	import { Tabs, TabsContent, TabsList, TabsTrigger } from '$lib/components/ui/tabs';
	import Page from '$lib/components/Signup.svelte';
	import Settings from '$lib/components/Settings.svelte';
	import { Marquee } from '@selemondev/svelte-marquee';
	import '@selemondev/svelte-marquee/dist/style.css';

	const { data } = $props();

	let currentTab = 'tab-0';
</script>

<div class="flex min-h-screen items-start justify-center overflow-hidden bg-black">
	<div class="fixed mx-auto w-full max-w-7xl py-4">
		<div class="px-4 py-4">
			<Header />
		</div>

		<Tabs
			value={currentTab}
			class="flex max-h-[calc(100dvh-3rem)]  w-full flex-col gap-8 md:flex-row"
		>
			<div class="grid w-full grid-cols-1 gap-4 overflow-y-auto px-10 pb-20 md:grid-cols-2">
				<div
					class="border-border relative col-span-1 min-w-[300px] rounded-lg border bg-zinc-900 p-8 md:col-span-2 md:min-h-[480px] md:overflow-y-hidden"
				>
					<TabsContent value="tab-0" class="h-[600px] md:min-h-[400px] xl:h-[600px]">
						<Tabs value="tab-1" class=" items-center">
							<TabsList
								class="border-border h-auto  gap-2 rounded-full border-b bg-zinc-800 px-10 py-2 text-zinc-400"
							>
								<TabsTrigger
									value="tab-1"
									class="relative cursor-pointer after:absolute  after:inset-x-0 after:bottom-0 after:-mb-1 after:h-0.5  hover:text-white hover:after:bg-white data-[state=active]:bg-transparent data-[state=active]:text-white data-[state=active]:shadow-none data-[state=active]:after:bg-white data-[state=active]:hover:text-white"
								>
									Home
								</TabsTrigger>

								<TabsTrigger
									value="tab-2"
									class="relative cursor-pointer after:absolute  after:inset-x-0 after:bottom-0 after:-mb-1 after:h-0.5  hover:text-white hover:after:bg-white data-[state=active]:bg-transparent data-[state=active]:text-white data-[state=active]:shadow-none data-[state=active]:after:bg-white data-[state=active]:hover:text-white"
								>
									Settings
								</TabsTrigger>
							</TabsList>
							<TabsContent value="tab-1">
								<Page {data} />
								<div class="hidden h-full items-center justify-center md:grid">
									<Marquee
										class="motion-reduce:overflow-auto "
										innerClassName="motion-reduce:animate-none motion-reduce:first:hidden"
										fade={true}
										pauseOnHover={true}
									>
										<img
											src="https://resources.tidal.com/images/ad522656/b4b6/4054/8b98/7ff39644cea6/640x640.jpg"
											alt="Album artwork for the track"
											class="mx-auto h-32 w-32 object-cover"
										/>

										<img
											src="https://resources.tidal.com/images/ad522656/b4b6/4054/8b98/7ff39644cea6/640x640.jpg"
											alt="album"
											class="mx-auto h-32 w-32 object-cover"
										/>
									</Marquee>
								</div>

								<div class="grid h-full items-center justify-center md:hidden">
									<Marquee direction="up" fade={true} pauseOnHover={true}>
										<img
											src="https://resources.tidal.com/images/ad522656/b4b6/4054/8b98/7ff39644cea6/640x640.jpg"
											alt=""
											class="mx-auto h-32 w-32 object-cover"
										/>
									</Marquee>
								</div>
							</TabsContent>
							<TabsContent value="tab-2">
								<Settings />
							</TabsContent>
						</Tabs>
					</TabsContent>
				</div>
			</div>
		</Tabs>
	</div>
</div>
