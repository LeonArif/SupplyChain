import type {
  BenchmarkPoint,
  Block,
  ChainValidation,
  CompareResult,
  Edge,
  Food,
  GraphState,
  Node,
  TransactionData,
} from '../types'

const API_BASE = import.meta.env.VITE_API_BASE ?? 'http://localhost:8080/api'

async function request<T>(path: string, init?: RequestInit): Promise<T> {
  const response = await fetch(`${API_BASE}${path}`, {
    headers: { 'Content-Type': 'application/json', ...init?.headers },
    ...init,
  })

  if (!response.ok) {
    const payload = await response.json().catch(() => ({ error: response.statusText }))
    throw new Error(payload.error ?? response.statusText)
  }

  if (response.status === 204) {
    return undefined as T
  }
  return response.json() as Promise<T>
}

export const api = {
  foods: () => request<Food[]>('/foods'),
  addFood: (food: Partial<Food>) => request<Food>('/foods', { method: 'POST', body: JSON.stringify(food) }),
  deleteFood: (id: string) => request<void>(`/foods/${id}`, { method: 'DELETE' }),
  graph: () => request<GraphState>('/graph'),
  addNode: (node: Node) => request<GraphState>('/graph/node', { method: 'POST', body: JSON.stringify(node) }),
  upsertEdge: (edge: Edge) => request<GraphState>('/graph/edge', { method: 'POST', body: JSON.stringify(edge) }),
  randomizeGraph: () => request<GraphState>('/graph/randomize', { method: 'POST' }),
  resetGraph: () => request<GraphState>('/graph/reset', { method: 'POST' }),
  compare: (foodId: string) => request<CompareResult>('/algo/compare', { method: 'POST', body: JSON.stringify({ foodId }) }),
  benchmark: () => request<BenchmarkPoint[]>('/algo/benchmark', { method: 'POST' }),
  chain: () => request<Block[]>('/chain'),
  checkIn: (data: TransactionData) => request<Block>('/chain/checkin', { method: 'POST', body: JSON.stringify(data) }),
  validateChain: () => request<ChainValidation>('/chain/validate'),
  tamper: (index: number, data: TransactionData) =>
    request<Block[]>('/chain/tamper', { method: 'POST', body: JSON.stringify({ index, data }) }),
  restore: () => request<Block[]>('/chain/restore', { method: 'POST' }),
}
