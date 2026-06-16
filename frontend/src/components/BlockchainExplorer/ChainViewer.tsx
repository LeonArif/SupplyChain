import type { Block, ChainValidation } from '../../types'
import { BlockCard } from './BlockCard'

type Props = {
  blocks: Block[]
  validation?: ChainValidation
}

export function ChainViewer({ blocks, validation }: Props) {
  const corrupted = validation && !validation.valid

  return (
    <section className="panel chain-panel">
      <div className="panel-title">
        <span>Blockchain Explorer</span>
        <strong>{blocks.length} block</strong>
      </div>
      {corrupted && (
        <div className="alert">
          BLOCKCHAIN CORRUPTED - {validation.message}
        </div>
      )}
      <div className="chain-list">
        {blocks.map((block) => <BlockCard key={`${block.index}-${block.hash}`} block={block} />)}
      </div>
    </section>
  )
}
