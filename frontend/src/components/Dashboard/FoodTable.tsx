import type { Food, Node } from '../../types'

type Props = {
  foods: Food[]
  nodes: Node[]
  selectedFoodId: string
  onSelect: (id: string) => void
  onDelete: (id: string) => void
}

export function FoodTable({ foods, nodes, selectedFoodId, onSelect, onDelete }: Props) {
  const nodeName = (id: string) => nodes.find((node) => node.id === id)?.name ?? id

  return (
    <section className="panel table-panel">
      <div className="panel-title">
        <span>Daftar Food</span>
        <strong>{foods.length} item</strong>
      </div>
      <div className="table-scroll">
        <table>
          <thead>
            <tr>
              <th>ID</th>
              <th>Nama</th>
              <th>Status</th>
              <th>Tujuan</th>
              <th>Hari</th>
              <th></th>
            </tr>
          </thead>
          <tbody>
            {foods.map((food) => (
              <tr key={food.id} className={selectedFoodId === food.id ? 'selected-row' : ''} onClick={() => onSelect(food.id)}>
                <td>{food.id}</td>
                <td>{food.name}</td>
                <td>
                  <span className={`badge ${food.status.toLowerCase()}`}>{food.status}</span>
                </td>
                <td>{nodeName(food.destination)}</td>
                <td>{food.daysLeft}</td>
                <td>
                  <button className="ghost danger" onClick={(event) => { event.stopPropagation(); onDelete(food.id) }}>
                    Hapus
                  </button>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </section>
  )
}
