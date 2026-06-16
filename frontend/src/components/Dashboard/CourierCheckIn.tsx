import { useState } from 'react'
import type { Food, Node, TransactionData } from '../../types'

type Props = {
  foods: Food[]
  nodes: Node[]
  onSubmit: (data: TransactionData) => Promise<void>
}

export function CourierCheckIn({ foods, nodes, onSubmit }: Props) {
  const [foodId, setFoodId] = useState(foods[0]?.id ?? '')
  const [location, setLocation] = useState(nodes[0]?.id ?? '')
  const [temperature, setTemperature] = useState(4)
  const [humidity, setHumidity] = useState(65)
  const [busy, setBusy] = useState(false)

  async function submit(event: React.FormEvent) {
    event.preventDefault()
    const food = foods.find((item) => item.id === foodId)
    if (!food) return

    setBusy(true)
    await onSubmit({
      foodId,
      foodName: food.name,
      location,
      temperature,
      humidity,
      expiryDate: food.expiryDate,
      courierId: 'COURIER-01',
      eventType: 'CHECK_IN',
    }).finally(() => setBusy(false))
  }

  return (
    <form className="panel form-grid" onSubmit={submit}>
      <div className="panel-title">
        <span>Check-In Kurir</span>
        <strong>ECDSA</strong>
      </div>
      <label>
        Food
        <select value={foodId} onChange={(event) => setFoodId(event.target.value)}>
          {foods.map((food) => (
            <option key={food.id} value={food.id}>{food.id} - {food.name}</option>
          ))}
        </select>
      </label>
      <label>
        Lokasi
        <select value={location} onChange={(event) => setLocation(event.target.value)}>
          {nodes.map((node) => (
            <option key={node.id} value={node.id}>{node.name}</option>
          ))}
        </select>
      </label>
      <label>
        Suhu C
        <input type="number" step="0.1" value={temperature} onChange={(event) => setTemperature(Number(event.target.value))} />
      </label>
      <label>
        Kelembapan %
        <input type="number" step="1" value={humidity} onChange={(event) => setHumidity(Number(event.target.value))} />
      </label>
      <button className="cta" disabled={busy || foods.length === 0} type="submit">
        Buat Blok
      </button>
    </form>
  )
}
