import type { CompareResult } from '../../types'

type Props = {
  result?: CompareResult
  onRun: () => Promise<void>
}

export function AlgoComparison({ result, onRun }: Props) {
  return (
    <section className="panel">
      <div className="panel-title">
        <span>Algorithm Lab</span>
        <button className="cta small" onClick={onRun}>Run</button>
      </div>
      {!result ? (
        <p className="muted">Pilih food, lalu jalankan perbandingan Greedy vs DP.</p>
      ) : (
        <div className="comparison-grid">
          <ResultCard title="Greedy" tone="greedy" result={result.greedy} />
          <ResultCard title="Dynamic Programming" tone="dp" result={result.dp} />
        </div>
      )}
    </section>
  )
}

function ResultCard({ title, tone, result }: { title: string; tone: string; result: CompareResult['greedy'] }) {
  return (
    <div className={`result-card ${tone}`}>
      <h3>{title}</h3>
      <p>{result.path?.join(' -> ') || 'Tidak ada rute'}</p>
      <dl>
        <div><dt>Mode</dt><dd>{result.mode}</dd></div>
        <div><dt>Cost</dt><dd>Rp {Math.round(result.totalCost).toLocaleString('id-ID')}</dd></div>
        <div><dt>Time</dt><dd>{result.totalTime.toFixed(1)} jam</dd></div>
        {/* <div><dt>Exec</dt><dd>{result.execTimeMs.toFixed(6)} ms</dd></div> */}
      </dl>
    </div>
  )
}
