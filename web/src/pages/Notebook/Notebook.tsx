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

  const notebook = useAppSelector(selectNotebookData)

  const status = useAppSelector(state => state.notebook.status)

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
    if (params.id !== notebook.id) {
      fetchData(params.id!)
    }
  }, [dispatch, params.id, notebook.id])

  switch (status) {
    case 'failed':
      return <div>an error occurred</div>
    case 'pending':
      return <div>loading</div>
    case 'idle':
    case 'succeeded':
      return <div>name: {notebook.name}</div>
  }
}
