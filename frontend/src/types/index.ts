export type FoodStatus = 'NORMAL' | 'WARNING' | 'CRITICAL'

export type Food = {
  id: string
  name: string
  expiryDate: string
  daysLeft: number
  destination: string
  weight: number
  urgency: number
  status: FoodStatus
  fingerprint: string
}

export type Node = {
  id: string
  name: string
  stage: number
}

export type Edge = {
  from: string
  to: string
  cost: number
  time: number
}

export type GraphState = {
  nodes: Node[]
  edges: Edge[]
}

export type RouteResult = {
  algorithm: string
  path: string[]
  totalCost: number
  totalTime: number
  execTimeMs: number
  mode: string
}

export type CompareResult = {
  food: Food
  greedy: RouteResult
  dp: RouteResult
}

export type BenchmarkPoint = {
  nodeCount: number
  greedyMs: number
  dpMs: number
}

export type TransactionData = {
  foodId: string
  foodName: string
  location: string
  temperature: number
  humidity: number
  expiryDate: string
  courierId: string
  eventType: string
}

export type Block = {
  index: number
  timestamp: string
  data: TransactionData
  prevHash: string
  hash: string
  signature: string
  valid: boolean
}

export type ChainValidation = {
  valid: boolean
  invalidIndex: number
  message: string
}
