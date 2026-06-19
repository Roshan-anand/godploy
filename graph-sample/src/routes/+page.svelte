<script lang="ts">
	import { graphlib, layout } from '@dagrejs/dagre';
	import ForceGraph from '$lib/ForceGraph.svelte';

	type ServiceKind = 'frontend' | 'app' | 'postgres' | 'redis';
	type BackendGraphNode = {
		id: string;
		name: string;
		kind: ServiceKind;
		status: 'running' | 'internal' | 'public';
		description: string;
	};

	type BackendGraphEdge = {
		id: string;
		from: string;
		to: string;
		envKey: string;
		targetCol: string;
	};

	type LayoutGraphNode = BackendGraphNode & {
		x: number;
		y: number;
	};

	type LayoutGraphEdge = BackendGraphEdge & {
		points: { x: number; y: number }[];
	};

	type DependencyGraphResponse = {
		instance: {
			id: string;
			name: string;
			type: 'production' | 'preview';
		};
		nodes: BackendGraphNode[];
		edges: BackendGraphEdge[];
	};

	const NODE_WIDTH = 190;
	const NODE_HEIGHT = 76;
	const PADDING = 56;

	const sampleGraph: DependencyGraphResponse = {
		instance: {
			id: 'inst_prod_01',
			name: 'production',
			type: 'production'
		},
		nodes: [
			{
				id: 'svc_frontend',
				name: 'web-frontend',
				kind: 'frontend',
				status: 'public',
				description: 'SvelteKit SPA'
			},
			{
				id: 'svc_gateway',
				name: 'api-gateway',
				kind: 'app',
				status: 'public',
				description: 'Public backend API'
			},
			{
				id: 'svc_billing',
				name: 'billing-service',
				kind: 'app',
				status: 'internal',
				description: 'Invoices + payments'
			},
			{
				id: 'svc_worker',
				name: 'worker-service',
				kind: 'app',
				status: 'internal',
				description: 'Async jobs'
			},
			{
				id: 'svc_postgres',
				name: 'main-postgres',
				kind: 'postgres',
				status: 'internal',
				description: 'Predefined DB'
			},
			{
				id: 'svc_redis',
				name: 'cache-redis',
				kind: 'redis',
				status: 'internal',
				description: 'Cache + queue'
			},
			{
				id: 'svc_admin',
				name: 'admin-console',
				kind: 'app',
				status: 'internal',
				description: 'Standalone service'
			}
		],
		edges: [
			{
				id: 'dep_frontend_api',
				from: 'svc_frontend',
				to: 'svc_gateway',
				envKey: 'PUBLIC_API_URL',
				targetCol: 'domain'
			},
			{
				id: 'dep_gateway_billing',
				from: 'svc_gateway',
				to: 'svc_billing',
				envKey: 'BILLING_URL',
				targetCol: 'internal_url'
			},
			{
				id: 'dep_gateway_worker',
				from: 'svc_gateway',
				to: 'svc_worker',
				envKey: 'WORKER_URL',
				targetCol: 'internal_url'
			},
			{
				id: 'dep_gateway_db',
				from: 'svc_gateway',
				to: 'svc_postgres',
				envKey: 'DATABASE_URL',
				targetCol: 'internal_url'
			},
			{
				id: 'dep_billing_db',
				from: 'svc_billing',
				to: 'svc_postgres',
				envKey: 'DATABASE_URL',
				targetCol: 'internal_url'
			},
			{
				id: 'dep_worker_db',
				from: 'svc_worker',
				to: 'svc_postgres',
				envKey: 'DATABASE_URL',
				targetCol: 'internal_url'
			},
			{
				id: 'dep_worker_redis',
				from: 'svc_worker',
				to: 'svc_redis',
				envKey: 'REDIS_URL',
				targetCol: 'internal_url'
			},
			{
				id: 'dep_gateway_redis',
				from: 'svc_gateway',
				to: 'svc_redis',
				envKey: 'CACHE_URL',
				targetCol: 'internal_url'
			}
		]
	};

	const kindLabel: Record<ServiceKind, string> = {
		frontend: 'Frontend',
		app: 'Microservice',
		postgres: 'Postgres',
		redis: 'Redis'
	};

	function createLayout(data: DependencyGraphResponse) {
		const g = new graphlib.Graph()
			.setGraph({ rankdir: 'LR', nodesep: 52, ranksep: 110, marginx: PADDING, marginy: PADDING })
			.setDefaultEdgeLabel(() => ({}));

		for (const node of data.nodes) {
			g.setNode(node.id, { ...node, width: NODE_WIDTH, height: NODE_HEIGHT });
		}

		for (const edge of data.edges) {
			g.setEdge(edge.from, edge.to, edge);
		}

		layout(g);

		const nodes: LayoutGraphNode[] = data.nodes.map((node) => ({
			...node,
			x: g.node(node.id).x as number,
			y: g.node(node.id).y as number
		}));

		const edges: LayoutGraphEdge[] = data.edges.map((edge) => ({
			...edge,
			points: (g.edge(edge.from, edge.to).points ?? []) as { x: number; y: number }[]
		}));

		return {
			width: (g.graph().width as number) + PADDING,
			height: (g.graph().height as number) + PADDING,
			nodes,
			edges
		};
	}

	function edgePath(points: { x: number; y: number }[]) {
		if (points.length === 0) return '';
		if (points.length === 1) return `M ${points[0].x} ${points[0].y}`;

		return points
			.map((point, index) => {
				if (index === 0) return `M ${point.x} ${point.y}`;
				const previous = points[index - 1];
				const midX = (previous.x + point.x) / 2;
				return `C ${midX} ${previous.y}, ${midX} ${point.y}, ${point.x} ${point.y}`;
			})
			.join(' ');
	}

	function edgeLabelPosition(points: { x: number; y: number }[]) {
		const middle = points[Math.floor(points.length / 2)] ?? { x: 0, y: 0 };
		return { x: middle.x, y: middle.y - 10 };
	}

	let selectedNodeId = $state<string | null>('svc_gateway');
	let selectedEdgeId = $state<string | null>(null);

	function selectEdge(edgeId: string) {
		selectedEdgeId = edgeId;
		selectedNodeId = null;
	}

	function selectNode(nodeId: string) {
		selectedNodeId = nodeId;
		selectedEdgeId = null;
	}

	function handleGraphKeydown(event: KeyboardEvent, select: () => void) {
		if (event.key === 'Enter' || event.key === ' ') {
			event.preventDefault();
			select();
		}
	}

	let graph = $derived(createLayout(sampleGraph));
	let selectedNode = $derived(sampleGraph.nodes.find((node) => node.id === selectedNodeId));
	let selectedEdge = $derived(sampleGraph.edges.find((edge) => edge.id === selectedEdgeId));

	/* ── pan / zoom ── */
	let zoom = $state(1);
	let panX = $state(0);
	let panY = $state(0);
	let isPanning = $state(false);
	let panStart = $state({ x: 0, y: 0 });
	let panStartOffset = $state({ x: 0, y: 0 });

	const MIN_ZOOM = 0.15;
	const MAX_ZOOM = 4;

	function handleWheel(event: WheelEvent) {
		event.preventDefault();
		const delta = -event.deltaY * 0.001;
		const newZoom = Math.min(MAX_ZOOM, Math.max(MIN_ZOOM, zoom + delta * zoom));
		zoom = newZoom;
	}

	function handlePanStart(event: MouseEvent) {
		if (event.button !== 0) return;
		isPanning = true;
		panStart = { x: event.clientX, y: event.clientY };
		panStartOffset = { x: panX, y: panY };
	}

	function handlePanMove(event: MouseEvent) {
		if (!isPanning) return;
		panX = panStartOffset.x + (event.clientX - panStart.x);
		panY = panStartOffset.y + (event.clientY - panStart.y);
	}

	function handlePanEnd() {
		isPanning = false;
	}

	function resetView() {
		zoom = 1;
		panX = 0;
		panY = 0;
	}
</script>



<main class="shell">


	<section class="canvas-card">
		<div class="canvas-toolbar">
			<div>
				<h2>Instance dependency graph</h2>
				<p>Drag to pan · Scroll to zoom</p>
			</div>
			<div class="legend">
				<span><i class="dot frontend"></i>Frontend</span>
				<span><i class="dot app"></i>Microservice</span>
				<span><i class="dot postgres"></i>Postgres</span>
				<span><i class="dot redis"></i>Redis</span>
			</div>
		</div>

		<!-- svelte-ignore a11y_no_noninteractive_element_interactions -->
		<div
			class="svg-wrap"
			class:panning={isPanning}
			role="application"
			aria-label="Zoomable service dependency graph. Drag to pan, scroll to zoom."
			onwheel={handleWheel}
			onmousedown={handlePanStart}
			onmousemove={handlePanMove}
			onmouseup={handlePanEnd}
			onmouseleave={handlePanEnd}
		>
			<svg
				viewBox={`0 0 ${graph.width} ${graph.height}`}
				role="img"
				aria-label="Service dependency graph"
				style="transform: scale({zoom}) translate({panX / zoom}px, {panY / zoom}px); transform-origin: 0 0;"
			>
				<defs>
					<marker
						id="arrowhead"
						markerWidth="12"
						markerHeight="12"
						refX="10"
						refY="6"
						orient="auto"
						markerUnits="strokeWidth"
					>
						<path d="M 1 1 L 11 6 L 1 11 z" />
					</marker>
				</defs>

				<g class="edges">
					{#each graph.edges as edge (edge.id)}
						{@const labelPosition = edgeLabelPosition(edge.points)}
						<g
							class:selected={selectedEdgeId === edge.id}
							class="edge"
							role="button"
							tabindex="0"
							onclick={() => selectEdge(edge.id)}
							onkeydown={(event) => handleGraphKeydown(event, () => selectEdge(edge.id))}
						>
							<path d={edgePath(edge.points)} />
							<!-- <foreignObject x={labelPosition.x - 76} y={labelPosition.y - 14} width="152" height="30">
								<div class="edge-label">{edge.envKey} ← {edge.targetCol}</div>
							</foreignObject> -->
						</g>
					{/each}
				</g>

				<g class="nodes">
					{#each graph.nodes as node (node.id)}
						<g
							class:selected={selectedNodeId === node.id}
							class={`node ${node.kind}`}
							transform={`translate(${node.x - NODE_WIDTH / 2}, ${node.y - NODE_HEIGHT / 2})`}
							role="button"
							tabindex="0"
							onclick={() => selectNode(node.id)}
							onkeydown={(event) => handleGraphKeydown(event, () => selectNode(node.id))}
						>
							<rect width={NODE_WIDTH} height={NODE_HEIGHT} rx="18" />
							<text x="18" y="27" class="node-name">{node.name}</text>
							<text x="18" y="49" class="node-meta">{kindLabel[node.kind]} · {node.status}</text>
							<text x="18" y="66" class="node-desc">{node.description}</text>
						</g>
					{/each}
				</g>
			</svg>
		</div>
	</section>

	<div class="zoom-bar">
		<span class="zoom-label">Zoom</span>
		<span class="zoom-value">{Math.round(zoom * 100)}%</span>
		<input
			class="zoom-slider"
			type="range"
			min={MIN_ZOOM}
			max={MAX_ZOOM}
			step="0.01"
			bind:value={zoom}
			title="Zoom level"
		/>
		<button class="zoom-reset" onclick={resetView}>Reset</button>
	</div>

	<section class="details-grid">
		<article class="panel">
			<h2>Backend JSON shape</h2>
			<p>
				This is the minimal response shape backend can send for graph rendering. Backend sends no
				coordinates; UI computes node x/y and edge points client-side with dagre. Includes one
				standalone service with no edges.
			</p>
			<pre>{JSON.stringify(sampleGraph, null, 2)}</pre>
		</article>

	</section>

	<section class="canvas-card force-graph">
		<div class="canvas-toolbar">
			<div>
				<h2>Force-directed layout (d3-force)</h2>
				<p>Same graph data, but nodes spread in all directions organically</p>
			</div>
		</div>
		<div class="svg-wrap svg-wrap--static">
			<ForceGraph nodes={sampleGraph.nodes} edges={sampleGraph.edges} />
		</div>
	</section>
</main>

<style>
	:global(body) {
		margin: 0;
		background: #080b12;
		color: #ecf3ff;
		font-family:
			Inter, ui-sans-serif, system-ui, -apple-system, BlinkMacSystemFont, 'Segoe UI', sans-serif;
	}

	.shell {
		min-height: 100vh;
		padding: 40px;
	}

	.hero,
	.canvas-toolbar,
	.details-grid {
		display: grid;
		gap: 24px;
		grid-template-columns: minmax(0, 1fr) auto;
		align-items: end;
	}

	.hero {
		margin: 0 auto 24px;
		max-width: 1280px;
	}

	.eyebrow {
		margin: 0 0 10px;
		color: #8fb3ff;
		font-size: 0.74rem;
		font-weight: 800;
		letter-spacing: 0.14em;
		text-transform: uppercase;
	}

	h2,
	p {
		margin-top: 0;
	}

	h2 {
		margin-bottom: 6px;
		font-size: 1rem;
	}

	p {
		color: #94a3b8;
		line-height: 1.6;
	}

	.badge,
	.canvas-card,
	.panel {
		border: 1px solid rgba(148, 163, 184, 0.18);
		background: rgba(15, 23, 42, 0.8);
		box-shadow: 0 24px 80px rgba(0, 0, 0, 0.32);
		backdrop-filter: blur(18px);
	}

	.badge {
		min-width: 180px;
		border-radius: 22px;
		padding: 18px;
	}

	.canvas-card {
		max-width: 1280px;
		margin: 0 auto;
		overflow: hidden;
		border-radius: 32px;
	}

	.canvas-toolbar {
		padding: 22px 24px;
		border-bottom: 1px solid rgba(148, 163, 184, 0.14);
	}

	/* ── zoom bar outside canvas ── */
	.zoom-bar {
		display: flex;
		align-items: center;
		gap: 12px;
		max-width: 1280px;
		margin: 14px auto 0;
		padding: 10px 18px;
		border: 1px solid rgba(148, 163, 184, 0.14);
		border-radius: 999px;
		background: rgba(8, 11, 18, 0.7);
		backdrop-filter: blur(10px);
		width: fit-content;
	}

	.zoom-label {
		color: #64748b;
		font-size: 0.75rem;
		font-weight: 700;
		text-transform: uppercase;
		letter-spacing: 0.06em;
	}

	.zoom-value {
		color: #e2e8f0;
		font-size: 0.95rem;
		font-weight: 700;
		font-variant-numeric: tabular-nums;
		min-width: 3.2ch;
		text-align: right;
	}

	.zoom-slider {
		-webkit-appearance: none;
		appearance: none;
		width: 100px;
		height: 4px;
		border-radius: 999px;
		background: rgba(148, 163, 184, 0.18);
		outline: none;
		cursor: pointer;
	}

	.zoom-slider::-webkit-slider-thumb {
		-webkit-appearance: none;
		appearance: none;
		width: 14px;
		height: 14px;
		border-radius: 999px;
		background: #38bdf8;
		border: 2px solid #0b1020;
		cursor: pointer;
		transition: transform 0.1s;
	}

	.zoom-slider::-webkit-slider-thumb:hover {
		transform: scale(1.25);
	}

	.zoom-slider::-moz-range-thumb {
		width: 14px;
		height: 14px;
		border-radius: 999px;
		background: #38bdf8;
		border: 2px solid #0b1020;
		cursor: pointer;
	}

	.zoom-reset {
		padding: 4px 14px;
		border: 1px solid rgba(148, 163, 184, 0.2);
		border-radius: 999px;
		background: transparent;
		color: #94a3b8;
		font-size: 0.78rem;
		font-weight: 600;
		letter-spacing: 0.02em;
		cursor: pointer;
		transition:
			background 0.15s,
			color 0.15s,
			border-color 0.15s;
	}

	.zoom-reset:hover {
		background: rgba(56, 189, 248, 0.12);
		color: #38bdf8;
		border-color: rgba(56, 189, 248, 0.35);
	}

	.legend {
		display: flex;
		flex-wrap: wrap;
		gap: 12px;
		color: #cbd5e1;
		font-size: 0.85rem;
	}

	.legend span {
		display: inline-flex;
		gap: 7px;
		align-items: center;
	}

	.dot {
		width: 10px;
		height: 10px;
		border-radius: 999px;
	}

	.dot.frontend {
		background: #38bdf8;
	}

	.dot.app {
		background: #a78bfa;
	}

	.dot.postgres {
		background: #22c55e;
	}

	.dot.redis {
		background: #f97316;
	}

	.svg-wrap {
		overflow: hidden;
		cursor: grab;
		user-select: none;
		background:
			radial-gradient(circle at 20% 10%, rgba(56, 189, 248, 0.16), transparent 28%),
			linear-gradient(rgba(148, 163, 184, 0.05) 1px, transparent 1px),
			linear-gradient(90deg, rgba(148, 163, 184, 0.05) 1px, transparent 1px),
			#0b1020;
		background-size: auto, 32px 32px, 32px 32px, auto;
	}

	.svg-wrap.panning {
		cursor: grabbing;
	}

	.svg-wrap--static {
		cursor: default;
	}

	svg {
		display: block;
		min-width: 960px;
		width: 100%;
		height: auto;
		will-change: transform;
	}

	.edge {
		cursor: pointer;
	}

	.edge path {
		fill: none;
		stroke: #64748b;
		stroke-width: 2.2;
		marker-end: url(#arrowhead);
	}

	.edge:hover path,
	.edge.selected path {
		stroke: #38bdf8;
		stroke-width: 3.2;
	}

	marker path {
		fill: #64748b;
	}

	.edge-label {
		display: inline-flex;
		max-width: 146px;
		align-items: center;
		justify-content: center;
		border: 1px solid rgba(148, 163, 184, 0.24);
		border-radius: 999px;
		background: rgba(8, 11, 18, 0.86);
		color: #dbeafe;
		font-size: 10px;
		font-weight: 700;
		line-height: 1;
		padding: 7px 9px;
		overflow: hidden;
		text-overflow: ellipsis;
		white-space: nowrap;
	}

	.node {
		cursor: pointer;
	}

	.node rect {
		fill: rgba(15, 23, 42, 0.96);
		stroke: rgba(148, 163, 184, 0.26);
		stroke-width: 1.4;
	}

	.node.frontend rect {
		stroke: rgba(56, 189, 248, 0.75);
	}

	.node.app rect {
		stroke: rgba(167, 139, 250, 0.75);
	}

	.node.postgres rect {
		stroke: rgba(34, 197, 94, 0.75);
	}

	.node.redis rect {
		stroke: rgba(249, 115, 22, 0.75);
	}

	.node:hover rect,
	.node.selected rect {
		filter: drop-shadow(0 0 18px rgba(56, 189, 248, 0.28));
		stroke-width: 2.5;
	}

	.node-name {
		fill: #f8fafc;
		font-size: 15px;
		font-weight: 800;
	}

	.node-meta,
	.node-desc {
		fill: #94a3b8;
		font-size: 11px;
	}

	.node-desc {
		fill: #64748b;
	}

	.details-grid {
		max-width: 1280px;
		margin: 24px auto 0;
		grid-template-columns: minmax(0, 1.4fr) minmax(300px, 0.6fr);
		align-items: start;
	}

	.panel {
		border-radius: 28px;
		padding: 22px;
	}

	pre {
		max-height: 460px;
		margin: 0;
		overflow: auto;
		border-radius: 18px;
		background: #020617;
		color: #bfdbfe;
		font-size: 0.78rem;
		line-height: 1.5;
		padding: 18px;
	}

	@media (max-width: 900px) {
		.shell {
			padding: 20px;
		}

		.hero,
		.canvas-toolbar,
		.details-grid {
			grid-template-columns: 1fr;
		}
	}
</style>
