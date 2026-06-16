import type { Edge, GraphState, RouteResult } from '../../types'

type Props = {
  graph: GraphState
  greedy?: RouteResult
  dp?: RouteResult
}

const stageLabels = ['Pabrik', 'Gudang', 'Hub', 'Retailer']

export function GraphCanvas({ graph, greedy, dp }: Props) {
  const grouped = stageLabels.map((_, stage) => graph.nodes.filter((node) => node.stage === stage))
  const positions = new Map<string, { x: number; y: number }>()
  grouped.forEach((nodes, stage) => {
    nodes.forEach((node, index) => {
      const gap = 340 / Math.max(nodes.length, 1)
      positions.set(node.id, { x: 80 + stage * 210, y: 90 + gap * index })
    })
  })

  const inPath = (edge: Edge, result?: RouteResult) => {
    const path = result?.path ?? []
    return path.some((node, index) => node === edge.from && path[index + 1] === edge.to)
  }

  return (
    <section className="panel graph-panel">
      <div className="panel-title">
        <span>Graf Distribusi</span>
        <strong>{graph.nodes.length} node</strong>
      </div>
      <svg viewBox="0 0 760 460" role="img" aria-label="Graf distribusi makanan">
        {stageLabels.map((label, index) => (
          <g key={label}>
            <text className="stage-label" x={80 + index * 210} y="34">{label}</text>
            <line className="stage-line" x1={80 + index * 210} x2={80 + index * 210} y1="52" y2="420" />
          </g>
        ))}
        {graph.edges.map((edge) => {
          const from = positions.get(edge.from)
          const to = positions.get(edge.to)
          if (!from || !to) return null
          const highlighted = inPath(edge, dp) ? 'dp-edge' : inPath(edge, greedy) ? 'greedy-edge' : ''
          return (
            <g key={`${edge.from}-${edge.to}`}>
              <line className={`edge ${highlighted}`} x1={from.x + 34} y1={from.y} x2={to.x - 34} y2={to.y} />
              <text className="edge-label" x={(from.x + to.x) / 2} y={(from.y + to.y) / 2 - 6}>
                Rp{Math.round(edge.cost / 1000)}k / {edge.time}j
              </text>
            </g>
          )
        })}
        {graph.nodes.map((node) => {
          const point = positions.get(node.id)
          if (!point) return null
          return (
            <g key={node.id}>
              <circle className={`node stage-${node.stage}`} cx={point.x} cy={point.y} r="32" />
              <text className="node-id" x={point.x} y={point.y - 4}>{node.id}</text>
              <text className="node-name" x={point.x} y={point.y + 14}>{node.name}</text>
            </g>
          )
        })}
      </svg>
    </section>
  )
}
