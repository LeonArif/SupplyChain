import type { GraphState, RouteResult } from '../../types'

type Props = {
  graph: GraphState
  greedy?: RouteResult
  dp?: RouteResult
}

export function RouteVisualizer({ graph, greedy, dp }: Props) {
  const label = (id: string) => graph.nodes.find((node) => node.id === id)?.name ?? id

  return (
    <section className="panel">
      <div className="panel-title"><span>Rute Terpilih</span><strong>Highlight</strong></div>
      <div className="route-list">
        <RouteLine title="Greedy" className="greedy-dot" path={greedy?.path ?? []} label={label} />
        <RouteLine title="DP" className="dp-dot" path={dp?.path ?? []} label={label} />
      </div>
    </section>
  )
}

function RouteLine({ title, className, path, label }: { title: string; className: string; path: string[]; label: (id: string) => string }) {
  return (
    <div>
      <span className={`dot ${className}`}></span>
      <strong>{title}</strong>
      <p>{path.length ? path.map(label).join(' -> ') : 'Belum dijalankan'}</p>
    </div>
  )
}
