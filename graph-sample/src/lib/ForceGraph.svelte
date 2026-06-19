<script lang="ts">
	import { forceSimulation, forceLink, forceManyBody, forceCenter, forceCollide, type SimulationNodeDatum } from 'd3-force';

	type ServiceKind = 'frontend' | 'app' | 'postgres' | 'redis';

	interface ServiceNode {
		id: string;
		name: string;
		kind: ServiceKind;
		status: string;
		description: string;
	}

	interface ServiceEdge {
		id: string;
		from: string;
		to: string;
	}

	interface PositionedNode extends ServiceNode {
		x: number;
		y: number;
	}

	type SimNode = ServiceNode & SimulationNodeDatum;

	let {
		nodes = [] as ServiceNode[],
		edges = [] as ServiceEdge[]
	}: {
		nodes: ServiceNode[];
		edges: ServiceEdge[];
	} = $props();

	const NODE_WIDTH = 190;
	const NODE_HEIGHT = 76;
	const PADDING = 80;
	const EDGE_INSET = 40;

	const kindLabel: Record<ServiceKind, string> = {
		frontend: 'Frontend',
		app: 'Microservice',
		postgres: 'Postgres',
		redis: 'Redis'
	};

	let layoutNodes = $state<PositionedNode[]>([]);
	let graphWidth = $state(600);
	let graphHeight = $state(400);

	$effect(() => {
		if (nodes.length === 0) return;

		const simNodes: SimNode[] = nodes.map((n) => ({ ...n }));
		const simLinks = edges.map((e) => ({ source: e.from, target: e.to }));

		const simulation = forceSimulation(simNodes)
			.force('link', forceLink(simLinks).id((d) => (d as SimNode).id).distance(240).strength(0.25))
			.force('charge', forceManyBody().strength(-800))
			.force('center', forceCenter(0, 0))
			.force('collide', forceCollide(110))
			.alpha(0.7)
			.stop();

		for (let i = 0; i < 300; i++) simulation.tick();

		let minX = Infinity,
			maxX = -Infinity,
			minY = Infinity,
			maxY = -Infinity;
		for (const n of simNodes) {
			if (n.x === undefined || n.y === undefined) continue;
			if (n.x < minX) minX = n.x;
			if (n.x > maxX) maxX = n.x;
			if (n.y < minY) minY = n.y;
			if (n.y > maxY) maxY = n.y;
		}

		const w = maxX - minX + NODE_WIDTH + PADDING * 2;
		const h = maxY - minY + NODE_HEIGHT + PADDING * 2;
		const ox = -minX + PADDING;
		const oy = -minY + PADDING;

		layoutNodes = simNodes.map((n) => ({
			id: n.id,
			name: n.name,
			kind: n.kind,
			status: n.status,
			description: n.description,
			x: (n.x ?? 0) + ox,
			y: (n.y ?? 0) + oy
		}));

		graphWidth = Math.max(w, 600);
		graphHeight = Math.max(h, 400);
	});

	function shorten(
		sx: number,
		sy: number,
		tx: number,
		ty: number,
		inset: number
	): { x: number; y: number } {
		const dx = tx - sx;
		const dy = ty - sy;
		const len = Math.sqrt(dx * dx + dy * dy);
		if (len < 1) return { x: tx, y: ty };
		return {
			x: tx - (dx / len) * inset,
			y: ty - (dy / len) * inset
		};
	}
</script>

<svg
	viewBox={`0 0 ${graphWidth} ${graphHeight}`}
	role="img"
	aria-label="Force-directed service dependency graph"
>
	<defs>
		<marker
			id="fa-arrowhead"
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
		{#each edges as edge (edge.id)}
			{@const src = layoutNodes.find((n) => n.id === edge.from)}
			{@const tgt = layoutNodes.find((n) => n.id === edge.to)}
			{#if src && tgt}
				{@const end = shorten(src.x, src.y, tgt.x, tgt.y, EDGE_INSET)}
				<line
					x1={src.x}
					y1={src.y}
					x2={end.x}
					y2={end.y}
					marker-end="url(#fa-arrowhead)"
				/>
			{/if}
		{/each}
	</g>

	<g class="nodes">
		{#each layoutNodes as node (node.id)}
			<g
				class={`node ${node.kind}`}
				transform={`translate(${node.x - NODE_WIDTH / 2}, ${node.y - NODE_HEIGHT / 2})`}
			>
				<rect width={NODE_WIDTH} height={NODE_HEIGHT} rx="18" />
				<text x="18" y="27" class="node-name">{node.name}</text>
				<text x="18" y="49" class="node-meta">{kindLabel[node.kind]} · {node.status}</text>
				<text x="18" y="66" class="node-desc">{node.description}</text>
			</g>
		{/each}
	</g>
</svg>

<style>
	svg {
		display: block;
		min-width: 100%;
		width: 100%;
		height: auto;
	}

	.edges line {
		stroke: #64748b;
		stroke-width: 2;
		marker-end: url(#fa-arrowhead);
	}

	marker path {
		fill: #64748b;
	}

	.node {
		cursor: default;
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
</style>
