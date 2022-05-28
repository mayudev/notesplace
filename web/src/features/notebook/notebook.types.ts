import { ProtectionLevel } from './notebook.slice'

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
