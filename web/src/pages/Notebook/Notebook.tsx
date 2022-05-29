import { useEffect } from 'react'
import { useParams } from 'react-router-dom'
import { useAppDispatch, useAppSelector } from '../../app/hooks'
import {
  fetchNotebook,
  selectNotebookData,
} from '../../features/notebook/notebook.slice'

type Params = {
  id: string
}

export default function Notebook() {
  let params = useParams<Params>()
  const dispatch = useAppDispatch()

  const currentId = useAppSelector(state => state.notebook.id)
  const notebook = useAppSelector(selectNotebookData)

  useEffect(() => {
    async function fetchData(id: string) {
      dispatch(
        fetchNotebook({
          id,
          jwt: '', // TODO
        })
      )
    }

    // If current notebook isn't already present in state, fetch it.
    if (params.id !== currentId) {
      fetchData(params.id!)
    }
  }, [dispatch, params.id, currentId])

  return (
    <div>
      <h1>Notebook {params.id!}!</h1>
      <div>{notebook.id}</div>
      <div>{notebook.name}</div>
    </div>
  )
}
