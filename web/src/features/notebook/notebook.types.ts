import { ProtectionLevel } from './notebook.slice'

export interface Notebook {
  id: string
  name: string | null
  protectionLevel: ProtectionLevel
  createdAt: string | null
  updatedAt: string | null
}

export interface NotebookCreateResponse {
  id: string | null
  status: string
  message: string
}
