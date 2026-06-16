import { useState } from 'react'
import type { Edge, Node } from '../../types'

type Props = {
  nodes: Node[]
  onSave: (edge: Edge) => Promise<void>
  onRandomize: () => Promise<void>
  onReset: () => Promise<void>
}

export function EdgeWeightEditor({ nodes, onSave, onRandomize, onReset }: Props) {
  const [from, setFrom] = useState('SRC')
  const [to, setTo] = useState(nodes[1]?.id ?? '')
  const [cost, setCost] = useState(100000)
  const [time, setTime] = useState(4)

  async function submit(event: React.FormEvent) {
    event.preventDefault()
    await onSave({ from, to, cost, time })
  }

  return (
    <form className="panel compact-form" onSubmit={submit}>
      <div className="panel-title"><span>Edge</span><strong>Cost/Time</strong></div>
      <select value={from} onChange={(event) => setFrom(event.target.value)}>
        {nodes.map((node) => <option key={node.id} value={node.id}>{node.id}</option>)}
      </select>
      <select value={to} onChange={(event) => setTo(event.target.value)}>
        {nodes.map((node) => <option key={node.id} value={node.id}>{node.id}</option>)}
      </select>
      <input type="number" min="0" value={cost} onChange={(event) => setCost(Number(event.target.value))} />
      <input type="number" min="0" step="0.5" value={time} onChange={(event) => setTime(Number(event.target.value))} />
      <button type="submit">Simpan Edge</button>
      <div className="split-buttons">
        <button type="button" className="ghost" onClick={onRandomize}>Random</button>
        <button type="button" className="ghost" onClick={onReset}>Reset</button>
      </div>
    </form>
  )
}
