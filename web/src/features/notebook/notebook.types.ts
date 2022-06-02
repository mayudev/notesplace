export enum ProtectionLevel {
  None,
  ReadOnly,
  Protected,
}

export type Status = 'idle' | 'pending' | 'succeeded' | 'failed'

export interface NotebookState extends Omit<Notebook, 'notes'> {
  status: Status
  error: string | undefined
}

export interface Notebook {
  id: string
  name: string | null
  protectionLevel: ProtectionLevel
  createdAt: string | null
  updatedAt: string | null
  notes: Note[] | null
}

export interface Note {
  id: string
  notebookId: string
  title: string | null
  order: number
  content: string | null
  createdAt: string | null
  updatedAt: string | null
}

export interface NotebookCreateResponse {
  id: string | null
  status: string
  message: string
}
