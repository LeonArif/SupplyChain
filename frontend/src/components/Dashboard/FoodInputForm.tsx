import { useMemo, useState } from 'react'
import type { Food, Node } from '../../types'

type Props = {
  retailers: Node[]
  onSubmit: (food: Partial<Food>) => Promise<void>
}

const tomorrow = new Date(Date.now() + 86400000).toISOString().slice(0, 10)

export function FoodInputForm({ retailers, onSubmit }: Props) {
  const [name, setName] = useState('Ikan Tuna Dingin')
  const [expiryDate, setExpiryDate] = useState(tomorrow)
  const [destination, setDestination] = useState(retailers[0]?.id ?? '')
  const [weight, setWeight] = useState(8)
  const [urgency, setUrgency] = useState(6)
  const [busy, setBusy] = useState(false)

  const daysLeft = useMemo(() => {
    const today = new Date()
    today.setHours(0, 0, 0, 0)
    return Math.floor((new Date(expiryDate).getTime() - today.getTime()) / 86400000)
  }, [expiryDate])

  async function submit(event: React.FormEvent) {
    event.preventDefault()
    setBusy(true)
    await onSubmit({ name, expiryDate, destination, weight, urgency }).finally(() => setBusy(false))
  }

  return (
    <form className="panel form-grid" onSubmit={submit}>
      <div className="panel-title">
        <span>Input Makanan</span>
        <strong>{daysLeft} hari</strong>
      </div>
      <label>
        Nama
        <input value={name} onChange={(event) => setName(event.target.value)} required />
      </label>
      <label>
        Expiry
        <input type="date" value={expiryDate} onChange={(event) => setExpiryDate(event.target.value)} required />
      </label>
      <label>
        Tujuan
        <select value={destination} onChange={(event) => setDestination(event.target.value)} required>
          {retailers.map((node) => (
            <option key={node.id} value={node.id}>
              {node.name}
            </option>
          ))}
        </select>
      </label>
      <label>
        Berat kg
        <input type="number" min="0.1" step="0.1" value={weight} onChange={(event) => setWeight(Number(event.target.value))} />
      </label>
      <label>
        Urgensi
        <input type="range" min="1" max="10" value={urgency} onChange={(event) => setUrgency(Number(event.target.value))} />
      </label>
      <button className="cta" disabled={busy} type="submit">
        Tambah
      </button>
    </form>
  )
}
