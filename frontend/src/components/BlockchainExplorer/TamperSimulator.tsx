import { useMemo, useState } from 'react'
import type { Block, TransactionData } from '../../types'

type Props = {
  blocks: Block[]
  onTamper: (index: number, data: TransactionData) => Promise<void>
  onRestore: () => Promise<void>
}

export function TamperSimulator({ blocks, onTamper, onRestore }: Props) {
  const [index, setIndex] = useState(1)
  const selected = useMemo(() => blocks.find((block) => block.index === index) ?? blocks[0], [blocks, index])
  const [temperature, setTemperature] = useState(3)

  async function tamper() {
    if (!selected) return
    await onTamper(selected.index, { ...selected.data, temperature })
  }

  return (
    <section className="panel compact-form">
      <div className="panel-title"><span>Tamper Test</span><strong>Attack</strong></div>
      <select value={index} onChange={(event) => setIndex(Number(event.target.value))}>
        {blocks.map((block) => <option key={block.index} value={block.index}>Block #{block.index}</option>)}
      </select>
      <input type="number" step="0.1" value={temperature} onChange={(event) => setTemperature(Number(event.target.value))} />
      <button className="danger-button" type="button" onClick={tamper}>Manipulasi Suhu</button>
      <button className="ghost" type="button" onClick={onRestore}>Restore Chain</button>
    </section>
  )
}
