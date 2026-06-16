import type { BenchmarkPoint } from '../../types'

type Props = {
  data: BenchmarkPoint[]
  onRun: () => Promise<void>
}

export function PerformanceChart({ data, onRun }: Props) {
  const values = data.flatMap((point) => [safeNumber(point.greedyMs), safeNumber(point.dpMs)])
  const maxY = Math.max(0.001, ...values)
  const xFor = (index: number) => 40 + index * (420 / Math.max(data.length - 1, 1))
  const yFor = (value: number) => 200 - Math.sqrt(safeNumber(value) / maxY) * 160
  const points = (key: 'greedyMs' | 'dpMs') =>
    data.map((point, index) => {
      return `${xFor(index)},${yFor(point[key])}`
    }).join(' ')

  return (
    <section className="panel">
      <div className="panel-title">
        <span>Benchmark</span>
        <button className="ghost" onClick={onRun}>Run</button>
      </div>
      <svg className="chart" viewBox="0 0 520 240" role="img" aria-label="Benchmark Greedy dan DP">
        <line className="axis" x1="40" y1="200" x2="480" y2="200" />
        <line className="axis" x1="40" y1="30" x2="40" y2="200" />
        <text className="chart-label y-label" x="40" y="24">ms</text>
        {data.length > 0 && <polyline className="chart-line greedy-line" points={points('greedyMs')} />}
        {data.length > 0 && <polyline className="chart-line dp-line" points={points('dpMs')} />}
        {data.map((point, index) => (
          <g key={point.nodeCount}>
            <circle className="chart-point greedy-point" cx={xFor(index)} cy={yFor(point.greedyMs)} r="4" />
            <circle className="chart-point dp-point" cx={xFor(index)} cy={yFor(point.dpMs)} r="4" />
            <text className="chart-label" x={xFor(index)} y="222">{point.nodeCount}</text>
          </g>
        ))}
      </svg>
      {data.length === 0 && <p className="muted">Klik Run untuk menghasilkan data kompleksitas waktu.</p>}
      {data.length > 0 && (
        <div className="benchmark-table">
          {data.map((point) => (
            <div key={point.nodeCount}>
              <span>{point.nodeCount} node</span>
              <strong>G {safeNumber(point.greedyMs).toFixed(4)} ms</strong>
              <strong>DP {safeNumber(point.dpMs).toFixed(4)} ms</strong>
            </div>
          ))}
        </div>
      )}
      <div className="legend"><span className="dot greedy-dot"></span>Greedy <span className="dot dp-dot"></span>DP</div>
    </section>
  )
}

function safeNumber(value: number | undefined) {
  return Number.isFinite(value) ? Number(value) : 0
}
