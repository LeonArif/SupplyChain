import { useState } from 'react'
import type { Node } from '../../types'

type Props = {
  onAdd: (node: Node) => Promise<void>
}

export function NodeEditor({ onAdd }: Props) {
  const [id, setId] = useState('')
  const [name, setName] = useState('')
  const [stage, setStage] = useState(1)

  async function submit(event: React.FormEvent) {
    event.preventDefault()
    await onAdd({ id: id.trim().toUpperCase(), name: name.trim(), stage })
    setId('')
    setName('')
  }

  return (
    <form className="panel compact-form" onSubmit={submit}>
      <div className="panel-title"><span>Node</span><strong>Stage</strong></div>
      <input placeholder="ID" value={id} onChange={(event) => setId(event.target.value)} required />
      <input placeholder="Nama node" value={name} onChange={(event) => setName(event.target.value)} required />
      <select value={stage} onChange={(event) => setStage(Number(event.target.value))}>
        <option value={0}>Pabrik</option>
        <option value={1}>Gudang</option>
        <option value={2}>Hub</option>
        <option value={3}>Retailer</option>
      </select>
      <button type="submit">Tambah Node</button>
    </form>
  )
}
