import type { Block } from '../../types'

type Props = {
  block: Block
}

const shortHash = (hash: string) => hash ? `${hash.slice(0, 10)}...${hash.slice(-6)}` : '-'

export function BlockCard({ block }: Props) {
  return (
    <article className={`block-card ${block.valid ? '' : 'invalid-block'}`}>
      <header>
        <strong>Block #{block.index}</strong>
        <span className={`badge ${block.valid ? 'normal' : 'critical'}`}>{block.valid ? 'VALID' : 'INVALID'}</span>
      </header>
      <p className="muted">{new Date(block.timestamp).toLocaleString('id-ID')}</p>
      <dl>
        <div><dt>Food</dt><dd>{block.data.foodId} - {block.data.foodName}</dd></div>
        <div><dt>Lokasi</dt><dd>{block.data.location}</dd></div>
        <div><dt>Suhu</dt><dd>{block.data.temperature} C</dd></div>
        <div><dt>Humidity</dt><dd>{block.data.humidity}%</dd></div>
        <div><dt>Event</dt><dd>{block.data.eventType}</dd></div>
      </dl>
      <footer>
        <span>Prev {shortHash(block.prevHash)}</span>
        <span>Hash {shortHash(block.hash)}</span>
      </footer>
    </article>
  )
}
