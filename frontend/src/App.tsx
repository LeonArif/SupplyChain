import { useEffect, useMemo, useState } from 'react'
import { api } from './api/client'
import { AlgoComparison } from './components/AlgorithmLab/AlgoComparison'
import { PerformanceChart } from './components/AlgorithmLab/PerformanceChart'
import { RouteVisualizer } from './components/AlgorithmLab/RouteVisualizer'
import { ChainViewer } from './components/BlockchainExplorer/ChainViewer'
import { TamperSimulator } from './components/BlockchainExplorer/TamperSimulator'
import { CourierCheckIn } from './components/Dashboard/CourierCheckIn'
import { FoodInputForm } from './components/Dashboard/FoodInputForm'
import { FoodTable } from './components/Dashboard/FoodTable'
import { EdgeWeightEditor } from './components/Graph/EdgeWeightEditor'
import { GraphCanvas } from './components/Graph/GraphCanvas'
import { NodeEditor } from './components/Graph/NodeEditor'
import type { BenchmarkPoint, Block, ChainValidation, CompareResult, Edge, Food, GraphState, Node, TransactionData } from './types'
import './App.css'

const emptyGraph: GraphState = { nodes: [], edges: [] }

function App() {
  const [foods, setFoods] = useState<Food[]>([])
  const [graph, setGraph] = useState<GraphState>(emptyGraph)
  const [blocks, setBlocks] = useState<Block[]>([])
  const [validation, setValidation] = useState<ChainValidation>()
  const [compareResult, setCompareResult] = useState<CompareResult>()
  const [benchmark, setBenchmark] = useState<BenchmarkPoint[]>([])
  const [selectedFoodId, setSelectedFoodId] = useState('')
  const [error, setError] = useState('')

  const retailers = useMemo(() => graph.nodes.filter((node) => node.stage === 3), [graph.nodes])

  async function refresh() {
    const [foodData, graphData, chainData, validationData] = await Promise.all([
      api.foods(),
      api.graph(),
      api.chain(),
      api.validateChain(),
    ])
    setFoods(foodData)
    setGraph(graphData)
    setBlocks(chainData)
    setValidation(validationData)
    setSelectedFoodId((current) => current || foodData[0]?.id || '')
  }

  useEffect(() => {
    refresh().catch((err) => setError(err.message))
  }, [])

  async function safe(action: () => Promise<void>) {
    setError('')
    try {
      await action()
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Terjadi kesalahan')
    }
  }

  const selectedFood = foods.find((food) => food.id === selectedFoodId)

  return (
    <main className="app-shell">
      <header className="topbar">
        <div>
          <p>FoodChain</p>
          <h1>Cold-chain dashboard with route optimization and tamper proof audit trail</h1>
        </div>
        <div className="health-pill">
          <span className={validation?.valid ? 'pulse ok' : 'pulse bad'}></span>
          {validation?.valid ? 'Chain valid' : 'Chain corrupted'}
        </div>
      </header>

      {error && <div className="alert">{error}</div>}

      <section className="kpi-grid">
        <Metric label="Produk" value={foods.length.toString()} />
        <Metric label="Node Graf" value={graph.nodes.length.toString()} />
        <Metric label="Blok" value={blocks.length.toString()} />
        <Metric label="Food Aktif" value={selectedFood?.status ?? '-'} />
      </section>

      <section className="layout two">
        <FoodInputForm retailers={retailers} onSubmit={(food) => safe(async () => { await api.addFood(food); await refresh() })} />
        <CourierCheckIn foods={foods} nodes={graph.nodes} onSubmit={(data) => safe(async () => { await api.checkIn(data); await refresh() })} />
      </section>

      <FoodTable
        foods={foods}
        nodes={graph.nodes}
        selectedFoodId={selectedFoodId}
        onSelect={setSelectedFoodId}
        onDelete={(id) => safe(async () => { await api.deleteFood(id); await refresh() })}
      />

      <section className="layout graph-layout">
        <GraphCanvas graph={graph} greedy={compareResult?.greedy} dp={compareResult?.dp} />
        <div className="side-stack">
          <NodeEditor onAdd={(node: Node) => safe(async () => setGraph(await api.addNode(node)))} />
          <EdgeWeightEditor
            nodes={graph.nodes}
            onSave={(edge: Edge) => safe(async () => setGraph(await api.upsertEdge(edge)))}
            onRandomize={() => safe(async () => setGraph(await api.randomizeGraph()))}
            onReset={() => safe(async () => setGraph(await api.resetGraph()))}
          />
        </div>
      </section>

      <section className="layout two">
        <AlgoComparison result={compareResult} onRun={() => safe(async () => setCompareResult(await api.compare(selectedFoodId)))} />
        <RouteVisualizer graph={graph} greedy={compareResult?.greedy} dp={compareResult?.dp} />
      </section>

      <PerformanceChart data={benchmark} onRun={() => safe(async () => setBenchmark(await api.benchmark()))} />

      <section className="layout blockchain-layout">
        <ChainViewer blocks={blocks} validation={validation} />
        <TamperSimulator
          blocks={blocks}
          onTamper={(index, data: TransactionData) => safe(async () => { setBlocks(await api.tamper(index, data)); setValidation(await api.validateChain()) })}
          onRestore={() => safe(async () => { setBlocks(await api.restore()); setValidation(await api.validateChain()) })}
        />
      </section>
    </main>
  )
}

function Metric({ label, value }: { label: string; value: string }) {
  return (
    <div className="metric">
      <span>{label}</span>
      <strong>{value}</strong>
    </div>
  )
}

export default App
